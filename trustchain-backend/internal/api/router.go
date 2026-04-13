package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"

	"github.com/amanp/trustchain-backend/internal/models"
)

// NewRouter builds and returns a configured Gin engine.
func NewRouter(db *pgxpool.Pool, log *zerolog.Logger) *gin.Engine {
	r := gin.New()

	// Use zerolog-compatible request logger
	r.Use(gin.Recovery())
	r.Use(requestLogger(log))

	// ── CORS ─────────────────────────────────────────────────────────────────
	corsConfig := cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://trustchain.xyz",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Authorization", "X-Request-ID",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	// ── Public routes ────────────────────────────────────────────────────────
	api := r.Group("/api")
	{
		api.GET("/health", healthHandler(db))
	}

	return r
}

// healthHandler returns 200 if the service and DB are healthy.
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
