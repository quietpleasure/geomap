package logger

import (
	"fmt"

	"geomap/pkg/zaplog"

	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

type Config struct {
	Level       string `yaml:"level"`
	TimeFormat  string `yaml:"time-format"`
	FilePath    string `yaml:"filepath"`
	WritePretty bool   `yaml:"write-pretty"`
	WithCaller  bool   `yaml:"with-caller"`
	FullCaller  bool   `yaml:"full-caller"`
	Rotation    struct {
		MaxSize         int  `yaml:"max-size"`
		MaxBackups      int  `yaml:"max-backups"`
		MaxAge          int  `yaml:"max-age"`
		Localtime       bool `yaml:"localtime"`
		Compress        bool `yaml:"compress"`
		RotateAtStartup bool `yaml:"rotate-startup"`
	} `yaml:"rotation"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	err := provider.Get("logger").Populate(&cfg)
	if err != nil {
		return nil, fmt.Errorf("logger config: %w", err)
	}
	return &cfg, nil
}

func New(cfg *Config) (*Logger, error) {
	log, err := zaplog.New(options(cfg)...)
	if err != nil {
		return nil, err
	}
	return &Logger{log}, nil
}

func Module() fx.Option {
	return fx.Module(
		"logger",
		fx.Provide(
			NewConfig,
			New,
		),
	)
}

func options(cfg *Config) []zaplog.Option {
	return append(
		make([]zaplog.Option, 0),
		zaplog.WithLevel(cfg.Level),
		zaplog.WithFile(cfg.FilePath),
		zaplog.WithCustomTimestamp(cfg.TimeFormat),
		zaplog.WithPretty(cfg.WritePretty),
		zaplog.WithCaller(cfg.WithCaller),
		zaplog.WithRotateAtStartup(cfg.Rotation.RotateAtStartup),
		zaplog.WithMaxSize(cfg.Rotation.MaxSize),
		zaplog.WithMaxBackups(cfg.Rotation.MaxBackups),
		zaplog.WithCompress(cfg.Rotation.Compress),
		zaplog.WithMaxAge(cfg.Rotation.MaxAge),
		zaplog.WithLocalTime(cfg.Rotation.Localtime),
	)
}
