package app

import (
	"context"
	"net/http"

	"geomap/internal/app/logger"
	"geomap/internal/app/router"
	"geomap/internal/app/server"
	"geomap/internal/app/storer"
	"geomap/internal/app/storer/postgres"

	// "github.com/ipfans/fxlogger"
	"go.uber.org/fx"
)

func Create() *fx.App {
	return fx.New(
		options(),
	)
}

func options() fx.Option {
	ctx := func() context.Context {
		return context.Background()
	}
	return fx.Options(
		logger.Module(),
		postgres.Module(),
		router.Module(),
		server.Module(),		
		// v1api.Module(),
		fx.Provide(
			NewConfig,
			fx.Annotate(ctx, fx.As(new(context.Context))),			
			fx.Annotate(
				// func(p *postgres.Postgres) *postgres.Postgres {
				// 	return p
				// },
				postgres.New,
				fx.As(new(storer.CoordinateProvider)),
			),
			fx.Annotate(router.New, fx.As(new(http.Handler))),
			// fx.Annotate(v1api.New,
			// 	fx.As(new(server.API)),
			// ),
		),
		// fx.WithLogger(func(log *logger.Logger) fxevent.Logger {
		// 	return &fxevent.ZapLogger{Logger: log.Logger}
		// }),
		fx.NopLogger,
	)
}
