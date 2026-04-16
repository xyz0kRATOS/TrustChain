package db

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

const migrationTable = "trustchain_schema_migrations"

type migrationFile struct {
	Version string
	Name    string
	Path    string
}

// ApplyMigrations executes pending *.up.sql migrations from migrationsDir.
// Applied versions are tracked in schema_migrations.
func ApplyMigrations(ctx context.Context, pool *pgxpool.Pool, migrationsDir string) error {
	if _, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS `+migrationTable+` (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`); err != nil {
		return fmt.Errorf("db: create migration tracker: %w", err)
	}

	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("db: read migrations dir %q: %w", migrationsDir, err)
	}

	files := make([]migrationFile, 0)
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".up.sql") {
			continue
		}
		base := strings.TrimSuffix(name, ".up.sql")
		parts := strings.SplitN(base, "_", 2)
		if len(parts) == 0 || strings.TrimSpace(parts[0]) == "" {
			continue
		}

		files = append(files, migrationFile{
			Version: parts[0],
			Name:    name,
			Path:    filepath.Join(migrationsDir, name),
		})
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	for _, f := range files {
		var alreadyApplied bool
		if err := pool.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM `+migrationTable+` WHERE version = $1)`, f.Version).Scan(&alreadyApplied); err != nil {
			return fmt.Errorf("db: check migration version %s: %w", f.Version, err)
		}
		if alreadyApplied {
			continue
		}

		sqlBytes, err := os.ReadFile(f.Path)
		if err != nil {
			return fmt.Errorf("db: read migration %s: %w", f.Name, err)
		}

		tx, err := pool.Begin(ctx)
		if err != nil {
			return fmt.Errorf("db: begin migration %s: %w", f.Name, err)
		}

		if _, err := tx.Exec(ctx, string(sqlBytes)); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("db: apply migration %s: %w", f.Name, err)
		}

		if _, err := tx.Exec(ctx, `INSERT INTO `+migrationTable+` (version) VALUES ($1)`, f.Version); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("db: record migration %s: %w", f.Name, err)
		}

		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("db: commit migration %s: %w", f.Name, err)
		}
	}

	return nil
}
