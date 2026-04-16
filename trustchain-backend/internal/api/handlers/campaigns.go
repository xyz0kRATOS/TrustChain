package handlers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"github.com/amanp/trustchain-backend/internal/models"
)

var walletRe = regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
var emailRe = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
var ethAddressRegex = regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)

type ApplyRequest struct {
	Name              string             `json:"name" binding:"required,min=3,max=100"`
	Category          string             `json:"category" binding:"required"`
	Description       string             `json:"description" binding:"required,min=10,max=2000"`
	GoalAmountUSD     float64            `json:"goalAmountUsd" binding:"required,gt=0"`
	ImageURL          *string            `json:"imageUrl"`
	CreatorWallet     string             `json:"creatorWallet" binding:"required"`
	CreatorName       string             `json:"creatorName" binding:"required"`
	CreatorEmail      string             `json:"creatorEmail" binding:"required,email"`
	CreatorOrg        *string            `json:"creatorOrg"`
	CreatorCountry    string             `json:"creatorCountry" binding:"required"`
	CreatorBio        string             `json:"creatorBio" binding:"required,max=500"`
	Milestones        []MilestoneRequest `json:"milestones" binding:"required,min=1,max=5,dive"`
	DocumentFileNames []string           `json:"documentFileNames"`
}

type MilestoneRequest struct {
	Name             string  `json:"name" binding:"required,max=100"`
	Description      string  `json:"description" binding:"required"`
	AmountUSD        float64 `json:"amountUsd" binding:"required,gt=0"`
	Deadline         string  `json:"deadline" binding:"required"`
	RequiredEvidence string  `json:"requiredEvidence" binding:"required"`
}

type CampaignHandler struct {
	db *pgxpool.Pool
}

func NewCampaignHandler(db *pgxpool.Pool) *CampaignHandler {
	return &CampaignHandler{db: db}
}

func (h *CampaignHandler) ListCampaigns(c *gin.Context) {
	ctx := c.Request.Context()
	status := strings.TrimSpace(c.Query("status"))
	limit := parseQueryInt(c.Query("limit"), 20)
	offset := parseQueryInt(c.Query("offset"), 0)

	where := ""
	args := []any{}
	if strings.EqualFold(status, "ACTIVE") {
		where = "WHERE LOWER(c.status) = 'live'"
	}

	q := fmt.Sprintf(`
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
			COALESCE(COUNT(DISTINCT d.donor_wallet), 0) AS donor_count,
			COALESCE(SUM(d.amount_wei), 0)::text AS total_raised_wei,
			c.created_at
		FROM campaigns c
		LEFT JOIN donations d ON d.campaign_id = c.id
		%s
		GROUP BY c.id
		ORDER BY c.created_at DESC
		LIMIT $1 OFFSET $2
	`, where)
	args = append(args, limit, offset)

	rows, err := h.db.Query(ctx, q, args...)
	if err != nil {
		msg := "failed to load campaigns"
		c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.CampaignResponse]{Data: nil, Error: &msg})
		return
	}
	defer rows.Close()

	results := make([]models.CampaignResponse, 0)
	for rows.Next() {
		var r models.CampaignResponse
		var goalWei string
		var createdAt time.Time
		if err := rows.Scan(
			&r.ID,
			&r.ContractAddress,
			&r.CreatorWallet,
			&r.Category,
			&r.Name,
			&r.Description,
			&r.GoalAmountUSD,
			&goalWei,
			&r.Status,
			&r.DocumentHash,
			&r.ImageURL,
			&r.CreatorName,
			&r.CreatorEmail,
			&r.CreatorOrg,
			&r.CreatorCountry,
			&r.CreatorBio,
			&r.DocumentFileNames,
			&r.DonorCount,
			&r.TotalRaisedWei,
			&createdAt,
		); err != nil {
			msg := "failed to parse campaigns"
			c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.CampaignResponse]{Data: nil, Error: &msg})
			return
		}
		r.CreatedAt = createdAt.UTC().Format(time.RFC3339)
		if goalWei == "0" {
			r.GoalAmountWei = nil
		} else {
			r.GoalAmountWei = &goalWei
		}
		milestones, err := h.listMilestones(ctx, r.ID)
		if err != nil {
			msg := "failed to load milestones"
			c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.CampaignResponse]{Data: nil, Error: &msg})
			return
		}
		r.Milestones = milestones
		results = append(results, r)
	}

	if err := rows.Err(); err != nil {
		msg := "failed to iterate campaigns"
		c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.CampaignResponse]{Data: nil, Error: &msg})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse[[]models.CampaignResponse]{Data: results, Error: nil})
}

func (h *CampaignHandler) GetCampaign(c *gin.Context) {
	ctx := c.Request.Context()
	id := strings.TrimSpace(c.Param("id"))

	q := `
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
			COALESCE(COUNT(DISTINCT d.donor_wallet), 0) AS donor_count,
			COALESCE(SUM(d.amount_wei), 0)::text AS total_raised_wei,
			c.created_at
		FROM campaigns c
		LEFT JOIN donations d ON d.campaign_id = c.id
		WHERE c.id = $1
		GROUP BY c.id
	`

	var r models.CampaignResponse
	var goalWei string
	var createdAt time.Time
	err := h.db.QueryRow(ctx, q, id).Scan(
		&r.ID,
		&r.ContractAddress,
		&r.CreatorWallet,
		&r.Category,
		&r.Name,
		&r.Description,
		&r.GoalAmountUSD,
		&goalWei,
		&r.Status,
		&r.DocumentHash,
		&r.ImageURL,
		&r.CreatorName,
		&r.CreatorEmail,
		&r.CreatorOrg,
		&r.CreatorCountry,
		&r.CreatorBio,
		&r.DocumentFileNames,
		&r.DonorCount,
		&r.TotalRaisedWei,
		&createdAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			msg := "Campaign not found"
			c.JSON(http.StatusNotFound, models.APIResponse[*models.CampaignResponse]{Data: nil, Error: &msg})
			return
		}
		msg := "failed to load campaign"
		c.JSON(http.StatusInternalServerError, models.APIResponse[*models.CampaignResponse]{Data: nil, Error: &msg})
		return
	}

	if goalWei == "0" {
		r.GoalAmountWei = nil
	} else {
		r.GoalAmountWei = &goalWei
	}
	r.CreatedAt = createdAt.UTC().Format(time.RFC3339)

	milestones, err := h.listMilestones(ctx, id)
	if err != nil {
		msg := "failed to load milestones"
		c.JSON(http.StatusInternalServerError, models.APIResponse[*models.CampaignResponse]{Data: nil, Error: &msg})
		return
	}
	r.Milestones = milestones

	c.JSON(http.StatusOK, models.APIResponse[models.CampaignResponse]{Data: r, Error: nil})
}

func (h *CampaignHandler) ApplyCampaign(c *gin.Context) {
	ctx := c.Request.Context()
	var req ApplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		return
	}

	if !ethAddressRegex.MatchString(req.CreatorWallet) {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Invalid wallet address format",
		})
		return
	}
	req.CreatorWallet = strings.ToLower(req.CreatorWallet)

	var totalMilestoneUSD float64
	deadlineByIndex := make([]time.Time, len(req.Milestones))
	for i, m := range req.Milestones {
		totalMilestoneUSD += m.AmountUSD

		deadline, err := time.Parse(time.RFC3339, m.Deadline)
		if err != nil {
			deadline, err = time.Parse("2006-01-02", m.Deadline)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"data":  nil,
					"error": fmt.Sprintf("Milestone %d has invalid deadline format", i+1),
				})
				return
			}
		}

		if deadline.Before(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{
				"data":  nil,
				"error": fmt.Sprintf("Milestone %d deadline must be in the future", i+1),
			})
			return
		}
		deadlineByIndex[i] = deadline
	}

	diff := totalMilestoneUSD - req.GoalAmountUSD
	if diff > 1.0 || diff < -1.0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"error": fmt.Sprintf(
				"Milestone amounts ($%.2f) must sum to campaign goal ($%.2f)",
				totalMilestoneUSD, req.GoalAmountUSD,
			),
		})
		return
	}

	tx, err := h.db.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Str("handler", "ApplyCampaign").Msg("failed to begin transaction")
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil, "error": "Internal server error"})
		return
	}
	defer tx.Rollback(ctx)

	campaignQ := `
		INSERT INTO campaigns (
			creator_wallet, category, name, description, goal_amount_usd,
			goal_amount_wei, status, image_url, creator_name, creator_email,
			creator_org, creator_country, creator_bio, document_file_names
		)
		VALUES ($1,$2,$3,$4,$5,0,'pending',$6,$7,$8,$9,$10,$11,$12)
		RETURNING id::text
	`

	var campaignID string
	err = tx.QueryRow(ctx, campaignQ,
		req.CreatorWallet,
		nullIfBlank(req.Category),
		strings.TrimSpace(req.Name),
		strings.TrimSpace(req.Description),
		req.GoalAmountUSD,
		nullIfBlankPtr(req.ImageURL),
		nullIfBlank(req.CreatorName),
		nullIfBlank(req.CreatorEmail),
		nullIfBlankPtr(req.CreatorOrg),
		nullIfBlank(req.CreatorCountry),
		nullIfBlank(req.CreatorBio),
		req.DocumentFileNames,
	).Scan(&campaignID)
	if err != nil {
		log.Error().Err(err).Str("handler", "ApplyCampaign").Msg("failed to insert campaign")
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil, "error": "Internal server error"})
		return
	}

	milestoneQ := `
		INSERT INTO milestones (
			campaign_id, sequence_index, name, description, amount_usd,
			amount_wei, deadline, required_evidence, status
		)
		VALUES ($1,$2,$3,$4,$5,0,$6,$7,'pending')
	`
	for i, m := range req.Milestones {
		_, err := tx.Exec(ctx, milestoneQ,
			campaignID,
			i,
			strings.TrimSpace(m.Name),
			strings.TrimSpace(m.Description),
			m.AmountUSD,
			deadlineByIndex[i].UTC(),
			strings.TrimSpace(m.RequiredEvidence),
		)
		if err != nil {
			log.Error().Err(err).Int("index", i).Str("handler", "ApplyCampaign").Msg("failed to insert milestone")
			c.JSON(http.StatusInternalServerError, gin.H{"data": nil, "error": "Internal server error"})
			return
		}
	}

	if _, err := tx.Exec(ctx, `
		INSERT INTO campaign_activity (type, campaign_id, campaign_name, wallet)
		VALUES ('campaign_applied', $1::uuid, $2, $3)
	`, campaignID, strings.TrimSpace(req.Name), req.CreatorWallet); err != nil {
		log.Warn().Err(err).Str("handler", "ApplyCampaign").Msg("failed to insert activity record")
	}

	if err := tx.Commit(ctx); err != nil {
		log.Error().Err(err).Str("handler", "ApplyCampaign").Msg("failed to commit transaction")
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil, "error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":      campaignID,
			"message": "Application received successfully",
		},
		"error": nil,
	})
}

func (h *CampaignHandler) listMilestones(ctx context.Context, campaignID string) ([]models.MilestoneResponse, error) {
	rows, err := h.db.Query(ctx, `
		SELECT
			id::text,
			sequence_index,
			COALESCE(name, 'Milestone ' || (sequence_index + 1)::text),
			description,
			COALESCE(amount_usd::text, '0'),
			amount_wei::text,
			deadline,
			required_evidence,
			LOWER(status)
		FROM milestones
		WHERE campaign_id = $1
		ORDER BY sequence_index ASC
	`, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]models.MilestoneResponse, 0)
	for rows.Next() {
		var m models.MilestoneResponse
		var amountWei string
		var deadline time.Time
		if err := rows.Scan(
			&m.ID,
			&m.SequenceIndex,
			&m.Name,
			&m.Description,
			&m.AmountUSD,
			&amountWei,
			&deadline,
			&m.RequiredEvidence,
			&m.Status,
		); err != nil {
			return nil, err
		}
		if amountWei == "0" {
			m.AmountWei = nil
		} else {
			m.AmountWei = &amountWei
		}
		m.Deadline = deadline.UTC().Format(time.RFC3339)
		out = append(out, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func validateApplication(req models.CampaignApplyRequest) string {
	if l := len(strings.TrimSpace(req.Name)); l < 3 || l > 100 {
		return "name must be 3-100 characters"
	}
	if l := len(strings.TrimSpace(req.Description)); l < 10 || l > 2000 {
		return "description must be 10-2000 characters"
	}
	if req.GoalAmountUSD <= 0 {
		return "goalAmountUsd must be greater than 0"
	}
	if !walletRe.MatchString(strings.TrimSpace(req.CreatorWallet)) {
		return "creatorWallet must be a valid wallet address"
	}
	if !emailRe.MatchString(strings.TrimSpace(req.CreatorEmail)) {
		return "creatorEmail must be a valid email"
	}
	if len(req.Milestones) < 1 || len(req.Milestones) > 5 {
		return "milestones must contain 1-5 items"
	}

	now := time.Now().UTC()
	sum := 0.0
	for i, m := range req.Milestones {
		if strings.TrimSpace(m.Name) == "" {
			return fmt.Sprintf("milestone %d name is required", i+1)
		}
		if strings.TrimSpace(m.Description) == "" {
			return fmt.Sprintf("milestone %d description is required", i+1)
		}
		if m.AmountUSD <= 0 {
			return fmt.Sprintf("milestone %d amountUsd must be greater than 0", i+1)
		}
		t, err := time.Parse(time.RFC3339, m.Deadline)
		if err != nil {
			return fmt.Sprintf("milestone %d deadline must be an ISO date", i+1)
		}
		if !t.After(now) {
			return "all milestone deadlines must be in the future"
		}
		sum += m.AmountUSD
	}
	if math.Abs(sum-req.GoalAmountUSD) > 1 {
		return "milestone amounts must sum to goalAmountUsd"
	}

	return ""
}

func parseQueryInt(raw string, fallback int) int {
	if raw == "" {
		return fallback
	}
	n, err := strconv.Atoi(raw)
	if err != nil || n < 0 {
		return fallback
	}
	return n
}

func nullIfBlank(v string) *string {
	trimmed := strings.TrimSpace(v)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}

func nullIfBlankPtr(v *string) *string {
	if v == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*v)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}
