package config

import (
	"fmt"
	"hr/internal/persistence"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DBConfig persistence.DBConfig `envPrefix:"DB_"`
}

func Load() (*Config, error) {
	godotenv.Load()

	config := Config{}
	err := env.Parse(&config)
	if err != nil {
		return nil, fmt.Errorf("parsing environment variables: %w", err)
	}

	return &config, nil
}
