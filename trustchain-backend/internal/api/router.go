package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"

	"github.com/amanp/trustchain-backend/internal/api/middleware"
	"github.com/amanp/trustchain-backend/internal/models"
)

// NewRouter builds and returns a configured Gin engine.
func NewRouter(db *pgxpool.Pool, log *zerolog.Logger) *gin.Engine {
	r := gin.New()

	// ── Middleware — CORS MUST be first ──────────────────────────────────────
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())
	r.Use(requestLogger(log))

	// ── Health check — no auth, no rate limit ────────────────────────────────
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"service":   "trustchain-backend",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
	})

	// ── DB health check ───────────────────────────────────────────────────────
	r.GET("/api/health/db", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
		defer cancel()
		if err := db.Ping(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "error",
				"error":  "database unreachable",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "database": "connected"})
	})

	// ── Public API routes ────────────────────────────────────────────────────
	api := r.Group("/api")
	{
		api.GET("/campaigns", listCampaignsHandler(db))
	}

	return r
}

// listCampaignsHandler is a stub that returns an empty list until the
// campaigns handler package is wired in.
func listCampaignsHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.APIResponse[[]any]{
			Data: []any{},
		})
	}
}

// healthHandler returns 200 if the service and DB are healthy.
// Kept for backward compatibility; the inline routes above are now preferred.
func healthHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := db.Ping(c.Request.Context()); err != nil {
			msg := "database unreachable: " + err.Error()
			c.JSON(http.StatusServiceUnavailable, models.APIResponse[any]{
				Error: &msg,
			})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[map[string]string]{
			Data: map[string]string{
				"status":    "ok",
				"service":   "trustchain-backend",
				"timestamp": time.Now().UTC().Format(time.RFC3339),
			},
		})
	}
}

// requestLogger is a minimal zerolog-based Gin middleware.
func requestLogger(log *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Msg("request")
	}
}
