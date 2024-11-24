package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgreSQL PostgreSQLConfig
}

type PostgreSQLConfig struct {
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     string `envconfig:"DB_PORT"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBName     string `envconfig:"DB_NAME"`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()
	var cfg Config
	err := envconfig.Process("", &cfg)
	return &cfg, err
}
