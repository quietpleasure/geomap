package app

import (
	"fmt"
	"path/filepath"

	"go.uber.org/config"
	"go.uber.org/fx"
	// cfg "github.com/tinrab/kit/cfg"
)

type ResultConfig struct {
	fx.Out

	Provider config.Provider
	Config   Config
}

type Config struct {
	Name string `yaml:"name"`
}

func NewConfig() (ResultConfig, error) {
	loader, err := config.NewYAML(config.File(filepath.Join("config", ".yaml")))
	if err != nil {
		return ResultConfig{}, fmt.Errorf("load config file: %w", err)
	}

	config := new(Config)
	if err := loader.Get("app").Populate(config); err != nil {
		return ResultConfig{}, fmt.Errorf("failed to populate config: %w", err)
	}

	return ResultConfig{
		Provider: loader,
		Config:   *config,
	}, nil
}
