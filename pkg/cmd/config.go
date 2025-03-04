package cmd

import (
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

// Config is the configuration for the server
type Config struct {
	Port        int           `env:"PORT" envDefault:"8080"`
	DBUser      string        `env:"DB_USER" envDefault:"postgres"`
	DBPassword  string        `env:"DB_PASSWORD" envDefault:"password"`
	DBHost      string        `env:"DB_HOST" envDefault:"localhost"`
	DBPort      int           `env:"DB_PORT" envDefault:"5432"`
	DBName      string        `env:"DB_NAME" envDefault:"postgres"`
	TokenExpiry time.Duration `env:"TOKEN_EXPIRY" envDefault:"24h"`
	JWTSecret   string        `env:"JWT_SECRET" envDefault:"secret"`
}

func newConfig() (*Config, error) {
	// `.env` をロード（エラーがあっても続行）
	_ = godotenv.Load()

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
