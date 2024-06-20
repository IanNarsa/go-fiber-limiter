package config

import (
	"os"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}

// ConnectDatabase establishes a connection to the database and returns the connection handle
func ConnectDatabase(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
