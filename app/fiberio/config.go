package server

import (
	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	DBURL   string `envconfig:"DB_URL"`
	AppPort string `envconfig:"APP_PORT"`
}

// NewConfig ...
func NewConfig() (*Config, error) {
	cfg := new(Config)
	if err := envconfig.Process("FIBERIO", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
