package config

import (
	"fmt"
	"os"

	"github.com/sjmaks/beatbase-back/internal/core/apperrs"
)

type Config struct {
	DatabaseURL string
}

func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	return cfg
}

func Load() (*Config, error) {
	dbURL, ok := os.LookupEnv("POSTGRES_URL")
	if !ok {
		return nil, fmt.Errorf("failed to look up db url: %w", apperrs.EnvNotFound)
	}

	cfg := &Config{
		DatabaseURL: dbURL,
	}

	return cfg, nil
}
