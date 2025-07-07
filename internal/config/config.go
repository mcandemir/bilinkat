package config

import "time"

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig   `envconfig:"SERVER"`
	App      AppConfig      `envconfig:"APP"`
	Database DatabaseConfig `envconfig:"DATABASE"`
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port         string        `envconfig:"SERVER_PORT" default:"3000"`
	Host         string        `envconfig:"SERVER_HOST" default:"localhost"`
	ReadTimeout  time.Duration `envconfig:"SERVER_READ_TIMEOUT" default:"30s"`
	WriteTimeout time.Duration `envconfig:"SERVER_WRITE_TIMEOUT" default:"30s"`
}

// AppConfig holds application-specific configuration
type AppConfig struct {
	BaseURL     string `envconfig:"APP_BASE_URL" default:"http://localhost:3000"`
	Environment string `envconfig:"APP_ENV" default:"development"`
	LogLevel    string `envconfig:"APP_LOG_LEVEL" default:"info"`
	LogFormat   string `envconfig:"APP_LOG_FORMAT" default:"json"` // json or text
	XAPIKey     string `envconfig:"X_API_KEY" required:"true"`
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	URL string `envconfig:"DATABASE_URL"`
}
