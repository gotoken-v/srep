package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SomeField string `envconfig:"SOME_FIELD"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
