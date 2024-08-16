package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"geomap/internal/app/logger"
	serverhttp "geomap/pkg/server-http"

	"go.uber.org/config"
	"go.uber.org/zap"
)

type Server struct {
	*http.Server
	// api API
}

type Config struct {
	Host           string        `yaml:"host"`
	Port           int           `yaml:"port"`
	MaxHeaderBytes int           `yaml:"max-header-bytes"`
	WriteTimeout   time.Duration `yaml:"write-timeout"`
	ReadTimeout    time.Duration `yaml:"read-timeout"`
	IdleTimeout    time.Duration `yaml:"idle-timeout"`
}


func NewConfig(provider config.Provider) (*Config, error) {
	var cfg *Config
	err := provider.Get("server").Populate(&cfg)
	if err != nil {
		return nil, fmt.Errorf("server config: %w", err)
	}
	return cfg, nil
}

func New(ctx context.Context, handler http.Handler, cfg *Config, log *logger.Logger) (*Server, error) {
	l := log.With(zap.String("address", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)))
	server, err := serverhttp.New(ctx, handler, options(cfg)...)
	if err != nil {
		l.Error("HTTP server", zap.Error(err))
		return nil, err
	}
	l.Debug("initialized HTTP server")
	return &Server{server}, nil
}

func options(cfg *Config) []serverhttp.Option {
	return append(
		make([]serverhttp.Option, 0),
		serverhttp.WithHost(cfg.Host),
		serverhttp.WithPort(cfg.Port),
		serverhttp.WithWriteTimeout(cfg.WriteTimeout),
		serverhttp.WithReadTimeout(cfg.ReadTimeout),
		serverhttp.WithIdleTimeout(cfg.IdleTimeout),
		serverhttp.WithMaxHeaderBytes(cfg.MaxHeaderBytes),
	)
}
