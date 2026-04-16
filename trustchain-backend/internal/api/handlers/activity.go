package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/amanp/trustchain-backend/internal/models"
)

type ActivityHandler struct {
	db *pgxpool.Pool
}

func NewActivityHandler(db *pgxpool.Pool) *ActivityHandler {
	return &ActivityHandler{db: db}
}

func (h *ActivityHandler) Recent(c *gin.Context) {
	ctx := c.Request.Context()
	rows, err := h.db.Query(ctx, `
		SELECT
			id::text,
			type,
			COALESCE(campaign_name, ''),
			COALESCE(campaign_id::text, ''),
			CASE WHEN amount_wei IS NULL THEN NULL ELSE amount_wei::text END,
			wallet,
			created_at,
			tx_hash
		FROM campaign_activity
		ORDER BY created_at DESC
		LIMIT 20
	`)
	if err != nil {
		msg := "failed to load activity"
		c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.ActivityEvent]{Data: nil, Error: &msg})
		return
	}
	defer rows.Close()

	out := make([]models.ActivityEvent, 0)
	for rows.Next() {
		var e models.ActivityEvent
		var createdAt time.Time
		if err := rows.Scan(&e.ID, &e.Type, &e.CampaignName, &e.CampaignID, &e.Amount, &e.Wallet, &createdAt, &e.TxHash); err != nil {
			msg := "failed to parse activity"
			c.JSON(http.StatusInternalServerError, models.APIResponse[[]models.ActivityEvent]{Data: nil, Error: &msg})
			return
		}
		e.Timestamp = createdAt.UTC().Format(time.RFC3339)
		out = append(out, e)
	}

	c.JSON(http.StatusOK, models.APIResponse[[]models.ActivityEvent]{Data: out, Error: nil})
}
