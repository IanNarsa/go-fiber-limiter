package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func LoadConfig() *Config {
	return &Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
}

// ConnectDatabase establishes a connection to the database and returns the connection handle
func ConnectDatabase(cfg *Config) (*sqlx.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
