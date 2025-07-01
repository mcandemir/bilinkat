package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Load loads configuration from environment variables
func Load() (*Config, error) {
	env := os.Getenv("ENV")

	// In production, don't load .env files, use environment variables directly
	if env == "production" {
		var cfg Config
		if err := envconfig.Process("", &cfg); err != nil {
			return nil, fmt.Errorf("failed to process config: %w", err)
		}
		return &cfg, nil
	}

	// For development/staging, try to load .env file
	envFile := fmt.Sprintf(".env.%s", env)
	if env == "" {
		envFile = ".env"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to process config: %w", err)
	}

	return &cfg, nil
}

// LoadFromFile loads configuration from a specific .env file
func LoadFromFile(filename string) (*Config, error) {
	if err := godotenv.Load(filename); err != nil {
		return nil, fmt.Errorf("failed to load .env file %s: %w", filename, err)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to process config: %w", err)
	}

	return &cfg, nil
}

// MustLoad loads configuration and panics if there's an error
func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}
