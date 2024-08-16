package server

import (
	"context"
	"net/http"
	"time"

	"geomap/internal/app/logger"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Module() fx.Option {
	return fx.Module(
		"server",
		fx.Provide(
			NewConfig,
			New,
		),
		fx.Decorate(func(log *logger.Logger) *logger.Logger {
			return &logger.Logger{Logger: log.Named("[SERVER]")}
		}),
		fx.Invoke(
			func(lc fx.Lifecycle, shutdowner fx.Shutdowner, web *Server, log *logger.Logger) error {
				lc.Append(
					fx.Hook{
						OnStart: func(_ context.Context) (err error) {
							log.Debug("starting HTTP server")
							go func() {
								err = web.ListenAndServe()
								if err != nil && err != http.ErrServerClosed {
									log.Error("starting HTTP server", zap.Error(err))
									shutdowner.Shutdown()
								}
							}()
							if err == nil {
								log.Debug("started HTTP server", zap.String("address", web.Addr))
							}
							return err
						},
						OnStop: func(ctx context.Context) error {
							log.Warn("stopping HTTP server")
							ctxShutdown, cancel := context.WithTimeout(ctx, 5*time.Second)
							defer cancel()
							web.SetKeepAlivesEnabled(false)
							err := web.Shutdown(ctxShutdown)
							if err != nil {
								log.Error("stopping HTTP server", zap.String("address", web.Addr), zap.Error(err))
								return err
							}
							log.Warn("stopped HTTP server", zap.String("address", web.Addr))
							return nil
						},
					},
				)
				return nil
			},
		),
	)
}
