package config

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresPort     int    `env:"POSTGRES_PORT" envDefault:"5433"`
	PostgresUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	PostgresDB       string `env:"POSTGRES_DB" envDefault:"ghostsend"`
	PostgresSSLMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
	MigrationsPath   string `env:"MIGRATION_PATH" envDefault:"db/migrations"`

	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	Port     int    `env:"PORT" envDefault:"8080"`
}

// New returns a new Config struct with values from environment variables and .env file
func New() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil && errors.Is(err, &fs.PathError{}) {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	var info Config
	if err = env.Parse(&info); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &info, nil

}
