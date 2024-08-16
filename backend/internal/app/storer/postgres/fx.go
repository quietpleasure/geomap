package postgres

import (
	"context"

	"geomap/internal/app/logger"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"postgres",
		fx.Provide(
			NewConfig,
			New,
		),
		fx.Decorate(func(log *logger.Logger) *logger.Logger {
			return &logger.Logger{Logger: log.Named("[POSTGRES]")}
		}),
		fx.Invoke(
			func(lc fx.Lifecycle, log *logger.Logger, psql *Postgres) {
				lc.Append(
					fx.Hook{
						OnStop: func(_ context.Context) error {
							psql.Close()
							log.Warn("pool is closed")
							return nil
						},
					},
				)
			},
		),
	)
}
