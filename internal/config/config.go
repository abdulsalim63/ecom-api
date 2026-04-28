package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Addr string // ":8080"
	Env  string // "local", "staging", "production"
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  string
}

// DSN builds the connection string from structured fields.
// No hardcoded passwords anywhere in source code.
func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode,
	)
}

// Load all configuration from env
func Load() (*Config, error) {
	port, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("config: DB_PORT must be an integer: %w", err)
	}

	cfg := &Config{
		App: AppConfig{
			Addr: getEnv("APP_ADDR", ":8080"),
			Env:  getEnv("APP_ENV", "local"),
		},
		Database: DatabaseConfig{
			Host:     mustGetEnv("DB_HOST"),
			Port:     port,
			Name:     mustGetEnv("DB_NAME"),
			User:     mustGetEnv("DB_USER"),
			Password: mustGetEnv("DB_PASSWORD"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		// Fail at startup — better than a cryptic runtime panic later
		panic(fmt.Sprintf("config: required env var %q is not set", key))
	}
	return v
}
