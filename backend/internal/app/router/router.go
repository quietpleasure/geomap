package router

import (
	"fmt"

	"geomap/internal/app/handlers"
	"geomap/internal/app/logger"
	"geomap/internal/app/middleware"
	"geomap/internal/app/storer"

	"github.com/gin-gonic/gin"
	"go.uber.org/config"
	"go.uber.org/fx"
	//uuid "github.com/matoous/go-nanoid/v2"
)

type Config struct {
	GinMode       string `yaml:"gin-mode"`
	HtmlTemplates string `yaml:"templates"`
	StaticFiles   string `yaml:"static-files"`
	Favicon       string `yaml:"favicon"`
	Sessions      struct {
		Name     string `yaml:"name"`
		MaxAge   int    `yaml:"max-age"`
		HttpOnly bool   `yaml:"http-only"`
		Secure   bool   `yaml:"secure"`
	} `yaml:"sessions"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	var cfg *Config
	err := provider.Get("router").Populate(&cfg)
	if err != nil {
		return nil, fmt.Errorf("router config: %w", err)
	}
	return cfg, nil
}

func Module() fx.Option {
	return fx.Module(
		"router",
		fx.Provide(
			NewConfig,
			New,
		),
	)
}

type router struct {
	middleware *middleware.Middleware
	handlers   *handlers.Handler //info: это можно использовать вместо разных версий api
}

func New(cfg *Config, log *logger.Logger, store storer.CoordinateProvider) *gin.Engine {
	router := &router{
		middleware: middleware.New(log),
		handlers:   handlers.New(log, store),
	}

	return router.gin(cfg)
}
