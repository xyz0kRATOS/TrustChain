package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/amanp/trustchain-backend/internal/api"
	"github.com/amanp/trustchain-backend/internal/blockchain"
	"github.com/amanp/trustchain-backend/internal/config"
	"github.com/amanp/trustchain-backend/internal/db"
)

func main() {
	// ── Logger ────────────────────────────────────────────────────────────────
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// ── Config ────────────────────────────────────────────────────────────────
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	lvl, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	log = log.Level(lvl)

	log.Info().Str("port", cfg.Port).Msg("TrustChain backend starting")

	// ── Database ──────────────────────────────────────────────────────────────
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := db.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}
	defer pool.Close()

	log.Info().Msg("database connection established")

	// Ensure DB schema is up to date for local/dev startup.
	migrateCtx, migrateCancel := context.WithTimeout(context.Background(), 30*time.Second)
	if err := db.ApplyMigrations(migrateCtx, pool, "internal/db/migrations"); err != nil {
		migrateCancel()
		log.Fatal().Err(err).Msg("failed to apply database migrations")
	}
	migrateCancel()
	log.Info().Msg("database migrations applied")

	// ── Blockchain watcher (Phase 2 scaffold) ─────────────────────────────────
	blockchain.Start(pool, cfg)

	// ── Router ────────────────────────────────────────────────────────────────
	router := api.NewRouter(pool, cfg, &log)

	// ── HTTP Server ───────────────────────────────────────────────────────────
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start in background
	go func() {
		log.Info().Str("addr", srv.Addr).Msg("HTTP server listening")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("HTTP server error")
		}
	}()

	// ── Graceful shutdown ─────────────────────────────────────────────────────
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("shutting down gracefully...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("server forced to shutdown")
	}
	log.Info().Msg("server stopped")
}
