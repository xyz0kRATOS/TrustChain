package blockchain

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"github.com/amanp/trustchain-backend/internal/blockchain/bindings"
	"github.com/amanp/trustchain-backend/internal/config"
)

var (
	topicDonationReceived  = common.HexToHash("0x3168ca6b1e292883c0668008556fb887f7cce92bcd3c1c57c15a34dcdccd5892")
	topicFundsReleased     = common.HexToHash("0x807158a396c8ce26fb6a4a44930e1b49876133ba7806dbaac491ea4713e4515e")
	topicMilestoneApproved = common.HexToHash("0xfe731b8534f38a55c98725a977efe67da793f35fb32ca4d1d947c01d80259bc2")
	topicCampaignFrozen    = common.HexToHash("0xc60a19caf4129aba2174b1b7a37d436ac19d86d92abc679d195808d7cae974b5")
	topicCampaignCreated   = common.HexToHash("0x3134cbb780ced4ba25bc7e71a70fee127918e70aa46c0e79eb2f9a8d0b56c4cf")
)

type Watcher struct {
	client    *Client
	pool      *pgxpool.Pool
	cfg       *config.Config
	lastBlock uint64
}

func NewWatcher(client *Client, pool *pgxpool.Pool, cfg *config.Config) *Watcher {
	return &Watcher{client: client, pool: pool, cfg: cfg}
}

func (w *Watcher) Start(ctx context.Context) {
	log.Info().Msg("blockchain watcher starting")
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	latest, err := w.client.Eth().BlockNumber(ctx)
	if err != nil {
		log.Error().Err(err).Msg("watcher: failed to get latest block")
	} else if latest > 100 {
		w.lastBlock = latest - 100
	}

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("blockchain watcher stopped")
			return
		case <-ticker.C:
			if err := w.poll(ctx); err != nil {
				log.Error().Err(err).Msg("watcher poll error")
			}
		}
	}
}

func (w *Watcher) poll(ctx context.Context) error {
	latest, err := w.client.Eth().BlockNumber(ctx)
	if err != nil {
		return err
	}
	if latest <= w.lastBlock {
		return nil
	}

	addresses, err := w.loadCampaignAddresses(ctx)
	if err != nil {
		return err
	}

	if common.IsHexAddress(w.cfg.CampaignFactoryAddress) {
		addresses = append(addresses, common.HexToAddress(w.cfg.CampaignFactoryAddress))
	}
	if len(addresses) == 0 {
		w.lastBlock = latest
		return nil
	}

	fromBlock := w.lastBlock + 1
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(latest)),
		Addresses: addresses,
	}

	logsFound, err := w.client.Eth().FilterLogs(ctx, query)
	if err != nil {
		return err
	}

	for _, vlog := range logsFound {
		if err := w.processLog(ctx, vlog); err != nil {
			log.Error().Err(err).Str("txHash", vlog.TxHash.Hex()).Msg("failed to process log")
		}
	}

	w.lastBlock = latest
	log.Debug().Uint64("from", fromBlock).Uint64("to", latest).Int("logs", len(logsFound)).Msg("watcher poll complete")
	return nil
}

func (w *Watcher) processLog(ctx context.Context, vlog types.Log) error {
	if len(vlog.Topics) == 0 {
		return nil
	}
	topic := vlog.Topics[0]

	if strings.EqualFold(vlog.Address.Hex(), w.cfg.CampaignFactoryAddress) && topic == topicCampaignCreated {
		factory, err := bindings.NewCampaignFactoryFilterer(vlog.Address, w.client.Eth())
		if err != nil {
			return err
		}
		evt, err := factory.ParseCampaignCreated(vlog)
		if err != nil {
			return err
		}
		_, _ = w.pool.Exec(ctx, `
			INSERT INTO campaign_activity (type, campaign_name, wallet, tx_hash, block_number, created_at)
			SELECT 'campaign_live', $1, $2, $3, $4, NOW()
			WHERE NOT EXISTS (
				SELECT 1 FROM campaign_activity WHERE type = 'campaign_live' AND tx_hash = $3
			)
		`, "On-chain campaign", strings.ToLower(evt.CreatorWallet.Hex()), vlog.TxHash.Hex(), int64(vlog.BlockNumber))
		return nil
	}

	campaignFilterer, err := bindings.NewCampaignFilterer(vlog.Address, w.client.Eth())
	if err != nil {
		return err
	}

	switch topic {
	case topicDonationReceived:
		evt, err := campaignFilterer.ParseDonationReceived(vlog)
		if err != nil {
			return err
		}
		return w.handleDonationReceived(ctx, vlog, evt, vlog.Address.Hex())
	case topicFundsReleased:
		evt, err := campaignFilterer.ParseFundsReleased(vlog)
		if err != nil {
			return err
		}
		return w.handleFundsReleased(ctx, vlog, evt, vlog.Address.Hex())
	case topicMilestoneApproved:
		evt, err := campaignFilterer.ParseMilestoneApproved(vlog)
		if err != nil {
			return err
		}
		return w.handleMilestoneApproved(ctx, vlog, evt, vlog.Address.Hex())
	case topicCampaignFrozen:
		evt, err := campaignFilterer.ParseCampaignFrozen(vlog)
		if err != nil {
			return err
		}
		return w.handleCampaignFrozen(ctx, vlog, evt, vlog.Address.Hex())
	default:
		return nil
	}
}

func (w *Watcher) handleDonationReceived(
	ctx context.Context,
	vlog types.Log,
	event *bindings.CampaignDonationReceived,
	campaignAddress string,
) error {
	campaignID, campaignName, err := w.campaignByAddress(ctx, campaignAddress)
	if err != nil {
		return err
	}
	if campaignID == "" {
		return nil
	}

	_, err = w.pool.Exec(ctx, `
		INSERT INTO donations (campaign_id, donor_wallet, amount_wei, tx_hash, block_number, donated_at)
		VALUES ($1::uuid, $2, $3, $4, $5, NOW())
		ON CONFLICT (tx_hash) DO NOTHING
	`, campaignID, strings.ToLower(event.Donor.Hex()), event.Amount.String(), vlog.TxHash.Hex(), int64(vlog.BlockNumber))
	if err != nil {
		return err
	}

	_, err = w.pool.Exec(ctx, `
		INSERT INTO access_grants (campaign_id, donor_wallet, granted_at)
		VALUES ($1::uuid, $2, NOW())
		ON CONFLICT (campaign_id, donor_wallet) DO NOTHING
	`, campaignID, strings.ToLower(event.Donor.Hex()))
	if err != nil {
		return err
	}

	_, err = w.pool.Exec(ctx, `
		INSERT INTO campaign_activity (type, campaign_id, campaign_name, wallet, amount_wei, tx_hash, block_number, created_at)
		SELECT 'donation', $1::uuid, $2, $3, $4, $5, $6, NOW()
		WHERE NOT EXISTS (
			SELECT 1 FROM campaign_activity WHERE type = 'donation' AND tx_hash = $5
		)
	`, campaignID, campaignName, strings.ToLower(event.Donor.Hex()), event.Amount.String(), vlog.TxHash.Hex(), int64(vlog.BlockNumber))
	return err
}

func (w *Watcher) handleFundsReleased(
	ctx context.Context,
	vlog types.Log,
	event *bindings.CampaignFundsReleased,
	campaignAddress string,
) error {
	campaignID, campaignName, err := w.campaignByAddress(ctx, campaignAddress)
	if err != nil {
		return err
	}
	if campaignID == "" {
		return nil
	}

	_, err = w.pool.Exec(ctx, `
		UPDATE milestones
		SET status = 'completed', released_at = NOW()
		WHERE campaign_id = $1::uuid AND sequence_index = $2
	`, campaignID, int(event.MilestoneIndex.Int64()))
	if err != nil {
		return err
	}

	_, err = w.pool.Exec(ctx, `
		INSERT INTO campaign_activity (type, campaign_id, campaign_name, amount_wei, tx_hash, block_number, created_at)
		SELECT 'funds_released', $1::uuid, $2, $3, $4, $5, NOW()
		WHERE NOT EXISTS (
			SELECT 1 FROM campaign_activity WHERE type = 'funds_released' AND tx_hash = $4
		)
	`, campaignID, campaignName, event.Amount.String(), vlog.TxHash.Hex(), int64(vlog.BlockNumber))
	return err
}

func (w *Watcher) handleMilestoneApproved(
	ctx context.Context,
	vlog types.Log,
	event *bindings.CampaignMilestoneApproved,
	campaignAddress string,
) error {
	campaignID, campaignName, err := w.campaignByAddress(ctx, campaignAddress)
	if err != nil {
		return err
	}
	if campaignID == "" {
		return nil
	}

	_, err = w.pool.Exec(ctx, `
		UPDATE milestones
		SET status = 'approved', approved_at = NOW()
		WHERE campaign_id = $1::uuid AND sequence_index = $2
	`, campaignID, int(event.MilestoneIndex.Int64()))
	if err != nil {
		return err
	}

	_, err = w.pool.Exec(ctx, `
		INSERT INTO campaign_activity (type, campaign_id, campaign_name, wallet, tx_hash, block_number, created_at)
		SELECT 'milestone_approved', $1::uuid, $2, $3, $4, $5, NOW()
		WHERE NOT EXISTS (
			SELECT 1 FROM campaign_activity WHERE type = 'milestone_approved' AND tx_hash = $4
		)
	`, campaignID, campaignName, strings.ToLower(event.Approver.Hex()), vlog.TxHash.Hex(), int64(vlog.BlockNumber))
	return err
}

func (w *Watcher) handleCampaignFrozen(
	ctx context.Context,
	vlog types.Log,
	event *bindings.CampaignCampaignFrozen,
	campaignAddress string,
) error {
	campaignID, campaignName, err := w.campaignByAddress(ctx, campaignAddress)
	if err != nil {
		return err
	}
	if campaignID == "" {
		return nil
	}

	_, err = w.pool.Exec(ctx, `
		UPDATE campaigns
		SET status = 'frozen', updated_at = NOW()
		WHERE id = $1::uuid
	`, campaignID)
	if err != nil {
		return err
	}

	_, err = w.pool.Exec(ctx, `
		INSERT INTO campaign_activity (type, campaign_id, campaign_name, wallet, tx_hash, block_number, created_at)
		SELECT 'campaign_frozen', $1::uuid, $2, $3, $4, $5, NOW()
		WHERE NOT EXISTS (
			SELECT 1 FROM campaign_activity WHERE type = 'campaign_frozen' AND tx_hash = $4
		)
	`, campaignID, campaignName, strings.ToLower(event.FrozenBy.Hex()), vlog.TxHash.Hex(), int64(vlog.BlockNumber))
	return err
}

func (w *Watcher) campaignByAddress(ctx context.Context, campaignAddress string) (string, string, error) {
	var campaignID string
	var campaignName string
	err := w.pool.QueryRow(ctx, `
		SELECT id::text, name
		FROM campaigns
		WHERE LOWER(contract_address) = LOWER($1)
	`, campaignAddress).Scan(&campaignID, &campaignName)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", "", nil
		}
		return "", "", err
	}
	return campaignID, campaignName, nil
}

func (w *Watcher) loadCampaignAddresses(ctx context.Context) ([]common.Address, error) {
	rows, err := w.pool.Query(ctx, `
		SELECT contract_address
		FROM campaigns
		WHERE contract_address IS NOT NULL
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	addresses := make([]common.Address, 0)
	for rows.Next() {
		var addrStr string
		if err := rows.Scan(&addrStr); err != nil {
			continue
		}
		if common.IsHexAddress(addrStr) {
			addresses = append(addresses, common.HexToAddress(addrStr))
		}
	}
	return addresses, rows.Err()
}

func Start(pool *pgxpool.Pool, cfg *config.Config) {
	if strings.TrimSpace(cfg.AlchemyBaseSepolia) == "" || strings.TrimSpace(cfg.AdminPrivateKey) == "" {
		log.Warn().Msg("blockchain watcher disabled: missing RPC URL or ADMIN_PRIVATE_KEY")
		return
	}

	client, err := NewClient(cfg.AlchemyBaseSepolia, cfg.AdminPrivateKey, cfg.ChainID)
	if err != nil {
		log.Error().Err(err).Msg("failed to initialize blockchain client for watcher")
		return
	}

	watcher := NewWatcher(client, pool, cfg)
	go watcher.Start(context.Background())

	startLiveCampaignDonationWatchers(context.Background(), client, pool)
}

func WatchCampaignDonations(ctx context.Context, client *Client, db *pgxpool.Pool, contractAddress string, campaignID string) {
	if !common.IsHexAddress(contractAddress) {
		log.Warn().Str("contractAddress", contractAddress).Str("campaignID", campaignID).Msg("invalid campaign contract address for donation watcher")
		return
	}

	addr := common.HexToAddress(contractAddress)
	filterer, err := bindings.NewCampaignFilterer(addr, client.Eth())
	if err != nil {
		log.Error().Err(err).Str("contractAddress", contractAddress).Str("campaignID", campaignID).Msg("failed to bind campaign filterer for donation watcher")
		return
	}

	var campaignName string
	if err := db.QueryRow(ctx, `
		SELECT name
		FROM campaigns
		WHERE id = $1::uuid
	`, campaignID).Scan(&campaignName); err != nil {
		log.Error().Err(err).Str("campaignID", campaignID).Msg("failed to load campaign name for donation watcher")
		return
	}

	eventCh := make(chan *bindings.CampaignDonationReceived)
	sub, err := filterer.WatchDonationReceived(&bind.WatchOpts{Context: ctx}, eventCh, nil, nil)
	if err != nil {
		log.Error().Err(err).Str("campaignID", campaignID).Str("contractAddress", contractAddress).Msg("failed to subscribe to DonationReceived")
		return
	}
	defer sub.Unsubscribe()

	log.Info().Str("campaignID", campaignID).Str("contractAddress", contractAddress).Msg("campaign donation watcher subscribed")

	for {
		select {
		case <-ctx.Done():
			log.Info().Str("campaignID", campaignID).Str("contractAddress", contractAddress).Msg("campaign donation watcher stopped")
			return
		case err := <-sub.Err():
			if err != nil {
				log.Error().Err(err).Str("campaignID", campaignID).Str("contractAddress", contractAddress).Msg("campaign donation watcher subscription error")
			}
			return
		case evt := <-eventCh:
			if evt == nil {
				continue
			}

			donorWallet := strings.ToLower(evt.Donor.Hex())
			txHash := evt.Raw.TxHash.Hex()
			amountWei := evt.Amount.String()

			_, err := db.Exec(ctx, `
				INSERT INTO donations (campaign_id, donor_wallet, amount_wei, tx_hash, block_number, donated_at)
				VALUES ($1::uuid, $2, $3, $4, $5, NOW())
				ON CONFLICT (tx_hash) DO NOTHING
			`, campaignID, donorWallet, amountWei, txHash, int64(evt.Raw.BlockNumber))
			if err != nil {
				log.Error().Err(err).Str("campaignID", campaignID).Str("txHash", txHash).Msg("failed to insert donation from subscription")
				continue
			}

			_, err = db.Exec(ctx, `
				INSERT INTO campaign_activity (type, campaign_id, campaign_name, wallet, tx_hash, amount_wei, created_at)
				VALUES ('donation_received', $1::uuid, $2, $3, $4, $5, NOW())
			`, campaignID, campaignName, donorWallet, txHash, amountWei)
			if err != nil {
				log.Error().Err(err).Str("campaignID", campaignID).Str("txHash", txHash).Msg("failed to insert campaign_activity from subscription")
				continue
			}

			_, err = db.Exec(ctx, `
				INSERT INTO access_grants (campaign_id, donor_wallet, granted_at)
				VALUES ($1::uuid, $2, NOW())
				ON CONFLICT (campaign_id, donor_wallet) DO NOTHING
			`, campaignID, donorWallet)
			if err != nil {
				log.Error().Err(err).Str("campaignID", campaignID).Str("txHash", txHash).Msg("failed to insert access grant from subscription")
				continue
			}
		}
	}
}

func startLiveCampaignDonationWatchers(ctx context.Context, client *Client, db *pgxpool.Pool) {
	rows, err := db.Query(ctx, `
		SELECT id::text, contract_address
		FROM campaigns
		WHERE LOWER(status) = 'live' AND contract_address IS NOT NULL
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to load live campaigns for donation watchers")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var campaignID string
		var contractAddress string
		if err := rows.Scan(&campaignID, &contractAddress); err != nil {
			log.Error().Err(err).Msg("failed to scan live campaign row for donation watcher")
			continue
		}

		go WatchCampaignDonations(ctx, client, db, contractAddress, campaignID)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("failed iterating live campaigns for donation watchers")
	}
}
