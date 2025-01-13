package store

import (
	"context"
	"fmt"
	"time"

	"github.com/asztemborski/syncro/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDb(config config.DatabaseConfig) (*sqlx.DB, error) {
	if config.Dsn == "" {
		return nil, fmt.Errorf("database DSN cannot be empty")
	}

	db, err := sqlx.Open("pgx", config.Dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxIdleTime(time.Duration(config.MaxIdleConns) * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return db, nil
}
