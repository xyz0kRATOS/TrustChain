package handlers

import (
	"context"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
 
	"github.com/amanp/trustchain-backend/internal/blockchain/bindings"

	"github.com/amanp/trustchain-backend/internal/blockchain"
	"github.com/amanp/trustchain-backend/internal/config"
	"github.com/amanp/trustchain-backend/internal/models"
)

type AdminHandler struct {
	db       *pgxpool.Pool
	campaign *CampaignHandler
	cfg      *config.Config
}

func NewAdminHandler(db *pgxpool.Pool, campaign *CampaignHandler, cfg *config.Config) *AdminHandler {
	return &AdminHandler{db: db, campaign: campaign, cfg: cfg}
}

type adminApproveRequest struct {
	AdminWallet string  `json:"adminWallet"`
	Notes       *string `json:"notes"`
}

type adminRejectRequest struct {
	AdminWallet string `json:"adminWallet"`
	Reason      string `json:"reason"`
}

func (h *AdminHandler) ListApplications(c *gin.Context) {
	ctx := c.Request.Context()
	rows, err := h.db.Query(ctx, `SELECT id::text FROM campaigns WHERE LOWER(status) = 'pending' ORDER BY created_at ASC`)
	if err != nil {
		msg := "failed to load applications"
		c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.CampaignResponse]{Data: nil, Error: &msg})
		return
	}
	defer rows.Close()

	out := make([]models.CampaignResponse, 0)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			msg := "failed to parse applications"
			c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.CampaignResponse]{Data: nil, Error: &msg})
			return
		}
		campaign, err := h.fetchCampaign(ctx, id)
		if err != nil {
			msg := "failed to load campaign"
			c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.CampaignResponse]{Data: nil, Error: &msg})
			return
		}
		out = append(out, *campaign)
	}

	c.JSON(http.StatusOK, models.APIResponse[[]models.CampaignResponse]{Data: out, Error: nil})
}

func (h *AdminHandler) ListCampaigns(c *gin.Context) {
	ctx := c.Request.Context()
	status := strings.ToLower(strings.TrimSpace(c.Query("status")))
	q := `SELECT id::text FROM campaigns`
	args := []any{}
	if status != "" {
		q += ` WHERE LOWER(status) = $1`
		args = append(args, status)
	}
	q += ` ORDER BY created_at DESC`

	rows, err := h.db.Query(ctx, q, args...)
	if err != nil {
		msg := "failed to load campaigns"
		c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.CampaignResponse]{Data: nil, Error: &msg})
		return
	}
	defer rows.Close()

	out := make([]models.CampaignResponse, 0)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			msg := "failed to parse campaigns"
			c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.CampaignResponse]{Data: nil, Error: &msg})
			return
		}
		campaign, err := h.fetchCampaign(ctx, id)
		if err != nil {
			msg := "failed to load campaign"
			c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.CampaignResponse]{Data: nil, Error: &msg})
			return
		}
		out = append(out, *campaign)
	}

	c.JSON(http.StatusOK, models.APIResponse[[]models.CampaignResponse]{Data: out, Error: nil})
}

func (h *AdminHandler) ApproveCampaign(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	var req adminApproveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := "Invalid request payload"
		c.JSON(http.StatusBadRequest, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}
	adminWallet := strings.ToLower(strings.TrimSpace(req.AdminWallet))

	tx, err := h.db.Begin(ctx)
	if err != nil {
		msg := "failed to start transaction"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}
	defer tx.Rollback(ctx)

	cmd, err := tx.Exec(ctx, `UPDATE campaigns SET status = 'live', updated_at = NOW() WHERE id = $1`, id)
	if err != nil || cmd.RowsAffected() == 0 {
		msg := "Campaign not found"
		c.JSON(http.StatusNotFound, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO admin_actions (action_type, campaign_id, admin_wallet, notes)
		VALUES ('campaign_approved', $1::uuid, $2, $3)
	`, id, adminWallet, req.Notes)
	if err != nil {
		msg := "failed to record admin action"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO campaign_activity (type, campaign_id, campaign_name, wallet)
		SELECT 'campaign_approved', id, name, $2 FROM campaigns WHERE id = $1
	`, id, adminWallet)
	if err != nil {
		msg := "failed to create activity"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		msg := "failed to approve campaign"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	campaign, err := h.fetchCampaign(ctx, id)
	if err != nil {
		msg := "failed to load campaign"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}
	c.JSON(http.StatusOK, models.APIResponse[models.CampaignResponse]{Data: *campaign, Error: nil})
}

func (h *AdminHandler) RejectCampaign(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	var req adminRejectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := "Invalid request payload"
		c.JSON(http.StatusBadRequest, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}
	if strings.TrimSpace(req.Reason) == "" {
		msg := "reason is required"
		c.JSON(http.StatusBadRequest, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}
	adminWallet := strings.ToLower(strings.TrimSpace(req.AdminWallet))

	tx, err := h.db.Begin(ctx)
	if err != nil {
		msg := "failed to start transaction"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}
	defer tx.Rollback(ctx)

	cmd, err := tx.Exec(ctx, `UPDATE campaigns SET status = 'rejected', updated_at = NOW() WHERE id = $1`, id)
	if err != nil || cmd.RowsAffected() == 0 {
		msg := "Campaign not found"
		c.JSON(http.StatusNotFound, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO admin_actions (action_type, campaign_id, admin_wallet, notes)
		VALUES ('campaign_rejected', $1::uuid, $2, $3)
	`, id, adminWallet, req.Reason)
	if err != nil {
		msg := "failed to record admin action"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		msg := "failed to reject campaign"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse[map[string]string]{Data: map[string]string{"message": "Campaign rejected"}, Error: nil})
}

func (h *AdminHandler) Stats(c *gin.Context) {
	ctx := c.Request.Context()
	type statsResponse struct {
		PendingCount   int    `json:"pendingCount"`
		LiveCount      int    `json:"liveCount"`
		CompletedCount int    `json:"completedCount"`
		FrozenCount    int    `json:"frozenCount"`
		TotalDonations int    `json:"totalDonations"`
		TotalRaisedUSD string `json:"totalRaisedUsd"`
	}

	var s statsResponse
	_ = h.db.QueryRow(ctx, `SELECT COUNT(*) FROM campaigns WHERE LOWER(status) = 'pending'`).Scan(&s.PendingCount)
	_ = h.db.QueryRow(ctx, `SELECT COUNT(*) FROM campaigns WHERE LOWER(status) = 'live'`).Scan(&s.LiveCount)
	_ = h.db.QueryRow(ctx, `SELECT COUNT(*) FROM campaigns WHERE LOWER(status) = 'completed'`).Scan(&s.CompletedCount)
	_ = h.db.QueryRow(ctx, `SELECT COUNT(*) FROM campaigns WHERE LOWER(status) = 'frozen'`).Scan(&s.FrozenCount)
	_ = h.db.QueryRow(ctx, `SELECT COUNT(*) FROM donations`).Scan(&s.TotalDonations)
	_ = h.db.QueryRow(ctx, `SELECT COALESCE(SUM(goal_amount_usd), 0)::text FROM campaigns WHERE LOWER(status) = 'live'`).Scan(&s.TotalRaisedUSD)

	c.JSON(http.StatusOK, models.APIResponse[statsResponse]{Data: s, Error: nil})
}

func (h *AdminHandler) DeployCampaign(c *gin.Context) {
	ctx := c.Request.Context()
	id := strings.TrimSpace(c.Param("id"))
	adminWallet := strings.ToLower(strings.TrimSpace(c.GetHeader("X-Admin-Wallet")))

	var status string
	var contractAddress string
	err := h.db.QueryRow(ctx, `
		SELECT LOWER(status), COALESCE(contract_address, '')
		FROM campaigns
		WHERE id = $1
	`, id).Scan(&status, &contractAddress)
	if err != nil {
		if err == pgx.ErrNoRows {
			msg := "Campaign not found"
			c.JSON(http.StatusNotFound, models.APIResponse[any]{Data: nil, Error: &msg})
			return
		}
		msg := "failed to load campaign"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	if contractAddress != "" {
		msg := "campaign is already deployed"
		c.JSON(http.StatusConflict, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}
	if status != "live" {
		msg := "campaign must be live before deployment"
		c.JSON(http.StatusBadRequest, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	deployResult, err := blockchain.DeployCampaign(ctx, h.cfg, h.db, id)
	if err != nil {
		msg := "deployment failed: " + err.Error()
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	_, _ = h.db.Exec(ctx, `
		INSERT INTO admin_actions (action_type, campaign_id, admin_wallet, notes, safe_tx_hash)
		VALUES ('campaign_deployed', $1::uuid, $2, $3, $4)
	`, id, adminWallet, "deployed to base sepolia", deployResult.TxHash)

	c.JSON(http.StatusOK, models.APIResponse[map[string]string]{
		Data: map[string]string{
			"campaignId":        id,
			"txHash":            deployResult.TxHash,
			"contractAddress":   deployResult.ContractAddress,
			"network":           "baseSepolia",
			"deploymentMessage": "Campaign deployed successfully",
		},
		Error: nil,
	})
}

func (h *AdminHandler) ApproveMilestone(c *gin.Context) {
	ctx := c.Request.Context()
	campaignID := strings.TrimSpace(c.Param("id"))
	idxRaw := strings.TrimSpace(c.Param("idx"))
	idx, err := strconv.ParseInt(idxRaw, 10, 64)
	if err != nil || idx < 0 {
		msg := "Invalid milestone index"
		c.JSON(http.StatusBadRequest, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	var contractAddress string
	var campaignName string
	err = h.db.QueryRow(ctx, `
		SELECT COALESCE(contract_address, ''), name
		FROM campaigns
		WHERE id = $1
	`, campaignID).Scan(&contractAddress, &campaignName)
	if err != nil {
		if err == pgx.ErrNoRows {
			msg := "Campaign not found"
			c.JSON(http.StatusNotFound, models.APIResponse[any]{Data: nil, Error: &msg})
			return
		}
		msg := "failed to load campaign"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	if contractAddress == "" {
		msg := "Campaign has not been deployed to blockchain yet"
		c.JSON(http.StatusBadRequest, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	client, err := blockchain.NewClient(h.cfg.AlchemyBaseSepolia, h.cfg.AdminPrivateKey, h.cfg.ChainID)
	if err != nil {
		msg := "failed to initialize blockchain client"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	opts, err := client.NewTransactOpts(ctx)
	if err != nil {
		msg := "failed to prepare transaction"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	campaignContract, err := bindings.NewCampaign(common.HexToAddress(contractAddress), client.Eth())
	if err != nil {
		msg := "failed to bind campaign contract"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	tx, err := campaignContract.ApproveMilestone(opts, big.NewInt(idx))
	if err != nil {
		msg := "failed to approve milestone on-chain"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	receiptCtx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()
	receipt, err := bind.WaitMined(receiptCtx, client.Eth(), tx)
	if err != nil || receipt.Status == types.ReceiptStatusFailed {
		msg := "on-chain milestone approval failed"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	_, _ = h.db.Exec(ctx, `
		UPDATE milestones
		SET status = 'approved', approved_at = NOW()
		WHERE campaign_id = $1::uuid AND sequence_index = $2
	`, campaignID, idx)

	adminWallet := strings.ToLower(strings.TrimSpace(c.GetHeader("X-Admin-Wallet")))
	_, _ = h.db.Exec(ctx, `
		INSERT INTO campaign_activity (type, campaign_id, campaign_name, wallet, tx_hash, created_at)
		VALUES ('milestone_approved', $1::uuid, $2, $3, $4, NOW())
	`, campaignID, campaignName, adminWallet, tx.Hash().Hex())

	releaseTime := time.Now().Add(5 * time.Minute).UTC().Format(time.RFC3339)
	c.JSON(http.StatusOK, models.APIResponse[map[string]string]{
		Data: map[string]string{
			"txHash":          tx.Hash().Hex(),
			"releaseTime":     releaseTime,
			"contractAddress": contractAddress,
		},
		Error: nil,
	})
}

func (h *AdminHandler) ExecuteMilestoneRelease(c *gin.Context) {
	ctx := c.Request.Context()
	campaignID := strings.TrimSpace(c.Param("id"))
	idxRaw := strings.TrimSpace(c.Param("idx"))
	idx, err := strconv.ParseInt(idxRaw, 10, 64)
	if err != nil || idx < 0 {
		msg := "Invalid milestone index"
		c.JSON(http.StatusBadRequest, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	var contractAddress string
	var campaignName string
	err = h.db.QueryRow(ctx, `
		SELECT COALESCE(contract_address, ''), name
		FROM campaigns
		WHERE id = $1
	`, campaignID).Scan(&contractAddress, &campaignName)
	if err != nil {
		if err == pgx.ErrNoRows {
			msg := "Campaign not found"
			c.JSON(http.StatusNotFound, models.APIResponse[any]{Data: nil, Error: &msg})
			return
		}
		msg := "failed to load campaign"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	if contractAddress == "" {
		msg := "Campaign has not been deployed to blockchain yet"
		c.JSON(http.StatusBadRequest, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}
	if strings.TrimSpace(h.cfg.TimelockAddress) == "" {
		msg := "TIMELOCK_ADDRESS is not configured"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	client, err := blockchain.NewClient(h.cfg.AlchemyBaseSepolia, h.cfg.AdminPrivateKey, h.cfg.ChainID)
	if err != nil {
		msg := "failed to initialize blockchain client"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	opts, err := client.NewTransactOpts(ctx)
	if err != nil {
		msg := "failed to prepare transaction"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	campaignParsed, err := bindings.CampaignMetaData.GetAbi()
	if err != nil {
		msg := "failed to load campaign ABI"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	data, err := campaignParsed.Pack("executeMilestoneRelease", big.NewInt(idx))
	if err != nil {
		msg := "failed to encode timelock call"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	timelock, err := bindings.NewTimelockController(common.HexToAddress(h.cfg.TimelockAddress), client.Eth())
	if err != nil {
		msg := "failed to bind timelock contract"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	var predecessor [32]byte
	var salt [32]byte
	big.NewInt(idx).FillBytes(salt[:])

	tx, err := timelock.Execute(
		opts,
		common.HexToAddress(contractAddress),
		big.NewInt(0),
		data,
		predecessor,
		salt,
	)
	if err != nil {
		msg := "failed to execute release through timelock"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	receiptCtx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()
	receipt, err := bind.WaitMined(receiptCtx, client.Eth(), tx)
	if err != nil || receipt.Status == types.ReceiptStatusFailed {
		msg := "on-chain release execution failed"
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Data: nil, Error: &msg})
		return
	}

	_, _ = h.db.Exec(ctx, `
		UPDATE milestones
		SET status = 'completed', released_at = NOW()
		WHERE campaign_id = $1::uuid AND sequence_index = $2
	`, campaignID, idx)

	adminWallet := strings.ToLower(strings.TrimSpace(c.GetHeader("X-Admin-Wallet")))
	_, _ = h.db.Exec(ctx, `
		INSERT INTO campaign_activity (type, campaign_id, campaign_name, wallet, tx_hash, created_at)
		VALUES ('funds_released', $1::uuid, $2, $3, $4, NOW())
	`, campaignID, campaignName, adminWallet, tx.Hash().Hex())

	c.JSON(http.StatusOK, models.APIResponse[map[string]string]{
		Data: map[string]string{
			"txHash":          tx.Hash().Hex(),
			"contractAddress": contractAddress,
		},
		Error: nil,
	})
}

func (h *AdminHandler) fetchCampaign(ctx context.Context, id string) (*models.CampaignResponse, error) {
	campaign := &models.CampaignResponse{}
	var goalWei string
	var createdAt time.Time
	err := h.db.QueryRow(ctx, `
		SELECT
			c.id::text,
			c.contract_address,
			LOWER(c.creator_wallet),
			c.category,
			c.name,
			c.description,
			COALESCE(c.goal_amount_usd::text, '0'),
			c.goal_amount_wei::text,
			LOWER(c.status),
			c.document_hash,
			c.image_url,
			c.creator_name,
			c.creator_email,
			c.creator_org,
			c.creator_country,
			c.creator_bio,
			COALESCE(c.document_file_names, '{}'::text[]),
			COALESCE(COUNT(DISTINCT d.donor_wallet), 0),
			COALESCE(SUM(d.amount_wei), 0)::text,
			c.created_at
		FROM campaigns c
		LEFT JOIN donations d ON d.campaign_id = c.id
		WHERE c.id = $1
		GROUP BY c.id
	`, id).Scan(
		&campaign.ID,
		&campaign.ContractAddress,
		&campaign.CreatorWallet,
		&campaign.Category,
		&campaign.Name,
		&campaign.Description,
		&campaign.GoalAmountUSD,
		&goalWei,
		&campaign.Status,
		&campaign.DocumentHash,
		&campaign.ImageURL,
		&campaign.CreatorName,
		&campaign.CreatorEmail,
		&campaign.CreatorOrg,
		&campaign.CreatorCountry,
		&campaign.CreatorBio,
		&campaign.DocumentFileNames,
		&campaign.DonorCount,
		&campaign.TotalRaisedWei,
		&createdAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	if goalWei != "0" {
		campaign.GoalAmountWei = &goalWei
	}
	campaign.CreatedAt = createdAt.UTC().Format(time.RFC3339)

	rows, err := h.db.Query(ctx, `
		SELECT id::text, sequence_index, COALESCE(name, 'Milestone ' || (sequence_index + 1)::text), description,
		COALESCE(amount_usd::text, '0'), amount_wei::text, deadline, required_evidence, LOWER(status)
		FROM milestones
		WHERE campaign_id = $1
		ORDER BY sequence_index ASC
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	campaign.Milestones = make([]models.MilestoneResponse, 0)
	for rows.Next() {
		var m models.MilestoneResponse
		var amountWei string
		var deadline time.Time
		if err := rows.Scan(&m.ID, &m.SequenceIndex, &m.Name, &m.Description, &m.AmountUSD, &amountWei, &deadline, &m.RequiredEvidence, &m.Status); err != nil {
			return nil, err
		}
		if amountWei != "0" {
			m.AmountWei = &amountWei
		}
		m.Deadline = deadline.UTC().Format(time.RFC3339)
		campaign.Milestones = append(campaign.Milestones, m)
	}

	return campaign, nil
}
