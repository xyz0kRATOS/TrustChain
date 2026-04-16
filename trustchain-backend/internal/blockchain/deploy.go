package blockchain

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"github.com/amanp/trustchain-backend/internal/blockchain/bindings"
	"github.com/amanp/trustchain-backend/internal/config"
)

type DeployResult struct {
	TxHash          string `json:"txHash"`
	ContractAddress string `json:"contractAddress"`
	CampaignID      uint64 `json:"campaignId"`
}

type campaignDeployData struct {
	CreatorWallet string
	GoalUSD       float64
	Name          string
}

type milestoneDeployData struct {
	Name             string
	Description      string
	AmountUSD        float64
	Deadline         time.Time
	RequiredEvidence string
}

const ETHPriceUSD = 3200.0

func DeployCampaign(ctx context.Context, cfg *config.Config, db *pgxpool.Pool, campaignID string) (*DeployResult, error) {
	if strings.TrimSpace(cfg.CampaignFactoryAddress) == "" {
		return nil, fmt.Errorf("CAMPAIGN_FACTORY_ADDRESS is required")
	}

	client, err := NewClient(cfg.AlchemyBaseSepolia, cfg.AdminPrivateKey, cfg.ChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create blockchain client: %w", err)
	}

	return DeployCampaignWithClient(ctx, client, db, cfg.CampaignFactoryAddress, campaignID)
}

func DeployCampaignWithClient(
	ctx context.Context,
	client *Client,
	pool *pgxpool.Pool,
	factoryAddress string,
	dbCampaignID string,
) (*DeployResult, error) {
	campaign, milestones, err := loadCampaignForDeploy(ctx, pool, dbCampaignID)
	if err != nil {
		return nil, fmt.Errorf("failed to load campaign: %w", err)
	}

	descArr := make([]string, len(milestones))
	amountArr := make([]*big.Int, len(milestones))
	deadlineArr := make([]*big.Int, len(milestones))
	evidenceArr := make([]string, len(milestones))

	for i, m := range milestones {
		descArr[i] = strings.TrimSpace(m.Name) + ": " + strings.TrimSpace(m.Description)
		evidenceArr[i] = strings.TrimSpace(m.RequiredEvidence)
		amountArr[i] = usdToWei(m.AmountUSD)
		deadlineArr[i] = big.NewInt(m.Deadline.Unix())
	}

	goalWei := usdToWei(campaign.GoalUSD)

	hashInput := dbCampaignID + campaign.Name
	hashBytes := sha256.Sum256([]byte(hashInput))
	var docHash [32]byte
	copy(docHash[:], hashBytes[:])

	opts, err := client.NewTransactOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create tx opts: %w", err)
	}

	factory, err := bindings.NewCampaignFactory(common.HexToAddress(factoryAddress), client.Eth())
	if err != nil {
		return nil, fmt.Errorf("failed to bind factory: %w", err)
	}

	log.Info().
		Str("campaign", campaign.Name).
		Str("goalWei", goalWei.String()).
		Int("milestones", len(milestones)).
		Msg("deploying campaign to blockchain")

	tx, err := factory.CreateCampaign(
		opts,
		common.HexToAddress(campaign.CreatorWallet),
		goalWei,
		docHash,
		descArr,
		amountArr,
		deadlineArr,
		evidenceArr,
	)
	if err != nil {
		return nil, fmt.Errorf("createCampaign transaction failed: %w", err)
	}

	receiptCtx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()

	receipt, err := waitForReceipt(receiptCtx, client.Eth(), tx)
	if err != nil {
		return nil, fmt.Errorf("waiting for receipt failed: %w", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("transaction reverted: https://sepolia.basescan.org/tx/%s", tx.Hash().Hex())
	}

	campaignAddress, campaignID, err := parseCampaignCreatedEvent(factory, receipt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CampaignCreated event: %w", err)
	}

	_, err = pool.Exec(ctx, `
		UPDATE campaigns
		SET contract_address = $1, updated_at = NOW()
		WHERE id = $2
	`, campaignAddress, dbCampaignID)
	if err != nil {
		return nil, fmt.Errorf("failed to persist deployed contract address: %w", err)
	}

	_, _ = pool.Exec(ctx, `
		INSERT INTO campaign_activity (type, campaign_id, campaign_name, wallet, tx_hash, created_at)
		VALUES ('campaign_live', $1::uuid, $2, $3, $4, NOW())
	`, dbCampaignID, campaign.Name, strings.ToLower(campaign.CreatorWallet), tx.Hash().Hex())

	return &DeployResult{
		TxHash:          tx.Hash().Hex(),
		ContractAddress: campaignAddress,
		CampaignID:      campaignID,
	}, nil
}

func waitForReceipt(ctx context.Context, eth *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(ctx, eth, tx)
}

func parseCampaignCreatedEvent(factory *bindings.CampaignFactory, receipt *types.Receipt) (string, uint64, error) {
	for _, lg := range receipt.Logs {
		evt, err := factory.ParseCampaignCreated(*lg)
		if err != nil {
			continue
		}
		if evt == nil {
			continue
		}
		return evt.CampaignAddress.Hex(), evt.CampaignId.Uint64(), nil
	}
	return "", 0, fmt.Errorf("CampaignCreated event not found")
}

func loadCampaignForDeploy(ctx context.Context, db *pgxpool.Pool, campaignID string) (*campaignDeployData, []milestoneDeployData, error) {
	var c campaignDeployData
	if err := db.QueryRow(ctx, `
		SELECT LOWER(creator_wallet), COALESCE(goal_amount_usd::float8, 0), name
		FROM campaigns
		WHERE id = $1
	`, campaignID).Scan(&c.CreatorWallet, &c.GoalUSD, &c.Name); err != nil {
		return nil, nil, err
	}

	rows, err := db.Query(ctx, `
		SELECT
			COALESCE(name, 'Milestone'),
			description,
			COALESCE(amount_usd::float8, 0),
			deadline,
			required_evidence
		FROM milestones
		WHERE campaign_id = $1
		ORDER BY sequence_index ASC
	`, campaignID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	milestones := make([]milestoneDeployData, 0)
	for rows.Next() {
		var m milestoneDeployData
		if err := rows.Scan(&m.Name, &m.Description, &m.AmountUSD, &m.Deadline, &m.RequiredEvidence); err != nil {
			return nil, nil, err
		}
		milestones = append(milestones, m)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	if len(milestones) == 0 {
		return nil, nil, fmt.Errorf("campaign has no milestones")
	}

	return &c, milestones, nil
}

func usdToWei(usd float64) *big.Int {
	ethPriceWei := new(big.Float).SetFloat64(1e18)
	usdPerEth := new(big.Float).SetFloat64(ETHPriceUSD)
	usdF := new(big.Float).SetFloat64(usd)
	weiF := new(big.Float).Mul(new(big.Float).Quo(usdF, usdPerEth), ethPriceWei)
	weiInt, _ := weiF.Int(nil)
	if weiInt == nil {
		return big.NewInt(0)
	}
	return weiInt
}
