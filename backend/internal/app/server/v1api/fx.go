package v1api

import (
	"geomap/internal/app/logger"

	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Logger *logger.Logger
}

type Result struct {
	fx.Out

	API API
}

func Module() fx.Option {
	return fx.Module(
		"api_v1",
		fx.Provide(
			New,
		),
		fx.Decorate(func(log *logger.Logger) *logger.Logger{
			return &logger.Logger{Logger: log.Named("[APIv1]")}
		} ),
	)
}