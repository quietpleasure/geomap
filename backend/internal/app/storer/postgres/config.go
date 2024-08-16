package postgres

import (
	"fmt"
	"net/url"
	"time"

	"geomap/pkg/pgxpool"

	"go.uber.org/config"
)

type Config struct {
	Database  string `yaml:"database"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	User      string `yaml:"user"`
	Pass      string `yaml:"pass"`
	AddParams struct {
		SSLMode               string        `yaml:"ssl-mode"`
		MaxConns              int           `yaml:"max-conns"`
		MaxConnLifeTime       time.Duration `yaml:"max-conn-lifetime"`
		MaxConnIdleTime       time.Duration `yaml:"max-conn-idle-time"`
		MaxConnLifeTimeJitter time.Duration `yaml:"max-conn-lifetime-jitter"`
		HealthCheckPeriod     time.Duration `yaml:"health-check-period"`
	} `yaml:"add-params"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	var cfg *Config
	err := provider.Get("postgres").Populate(&cfg)
	if err != nil {
		return nil, fmt.Errorf("postgres config: %w", err)
	}
	return cfg, nil
}

func (c *Config) String() string {
	url := &url.URL{
		Scheme: pgxpool.SELF_NAME,
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   c.Database,
		User:   url.UserPassword(c.User, c.Pass),
	}
	return url.String()
}
