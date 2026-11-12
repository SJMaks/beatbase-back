package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sjmaks/beatbase-back/internal/core/config"
	"go.uber.org/zap"
)

func MustNewPool(ctx context.Context, log *zap.Logger, cfg *config.Config) *pgxpool.Pool {
	pool, err := NewPool(ctx, cfg)
	if err != nil {
		log.Error(fmt.Errorf("failed to create new pool: %w", err).Error())
		panic(fmt.Sprintf("failed to create new pool: %v", err))
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		log.Error(fmt.Sprintf("failed to ping pool: %v", err))
		panic(fmt.Sprintf("failed to ping pool: %v", err))
	}

	return pool
}

func NewPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
