package postgres

import (
	"context"

	"geomap/internal/app/logger"
	pool "geomap/pkg/pgxpool"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Postgres struct {
	*pgxpool.Pool
}

func New(ctx context.Context, cfg *Config, log *logger.Logger) (*Postgres, error) {
	l := log.With(zap.Any("address", cfg))
	pool, err := pool.New(ctx, options(cfg)...)
	if err != nil {
		l.Error("opening a new pool", zap.Error(err))
		return nil, err
	}
	l.Debug("new pool opened")
	return &Postgres{pool}, nil
}

func options(cfg *Config) []pool.Option {
	return append(
		make([]pool.Option, 0),
		pool.WithHost(cfg.Host),
		pool.WithPort(cfg.Port),
		pool.WithDatabase(cfg.Database),
		pool.WithUser(cfg.User),
		pool.WithPass(cfg.Pass),
		pool.WithSSLMode(cfg.AddParams.SSLMode),
		pool.WithMaxConns(cfg.AddParams.MaxConns),
		pool.WithMaxConnLifeTime(cfg.AddParams.MaxConnLifeTime),
		pool.WithMaxConnLifeTimeJitter(cfg.AddParams.MaxConnLifeTimeJitter),
		pool.WithMaxConnIdleTime(cfg.AddParams.MaxConnIdleTime),
		pool.WithHealthCheckPeriod(cfg.AddParams.HealthCheckPeriod),
	)
}
