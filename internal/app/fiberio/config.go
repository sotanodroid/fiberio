package server

import (
	"github.com/joeshaw/envdecode"
)

// Config ...
type Config struct {
	DBURL string `env:"DB_URL,required"`
}

// NewConfig ...
func NewConfig() (*Config, error) {
	cfg := new(Config)
	if err := envdecode.Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
