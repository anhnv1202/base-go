// Package config handles application configuration loading using Viper.
// It supports environment-based configuration with .env files.
package initialize

import (
	"fmt"

	"github.com/anhnv1202/base-go/global"
	"github.com/anhnv1202/base-go/pkg/setting"
	"github.com/spf13/viper"
)

// Load reads configuration from environment variables and .env file
func LoadConfig() {
	// Set default values
	setDefaults()

	// Read from .env file if exists
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	// It's okay if .env doesn't exist - we'll use env vars
	_ = viper.ReadInConfig()

	// Read from environment variables
	viper.AutomaticEnv()

	global.Config = &setting.Config{}

	if err := viper.Unmarshal(global.Config); err != nil {
		panic(fmt.Sprintf("failed to unmarshal config: %v", err))
	}

	if err := validate(global.Config); err != nil {
		panic(fmt.Sprintf("config validation failed: %v", err))
	}

}

// setDefaults sets default configuration values
func setDefaults() {
	// App defaults
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("APP_NAME", "ecophone-backend")

	// Database defaults
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_SSL_MODE", "disable")
	viper.SetDefault("DB_MAX_IDLE_CONNS", 10)
	viper.SetDefault("DB_MAX_OPEN_CONNS", 100)
	viper.SetDefault("DB_CONN_MAX_LIFETIME", "1h")

	// JWT defaults
	viper.SetDefault("JWT_EXPIRY", "24h")
	viper.SetDefault("JWT_REFRESH_EXPIRY", "168h")

	// Log defaults
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("LOG_FORMAT", "json")
	viper.SetDefault("LOG_FILE_PATH", "logs/app.log")
	viper.SetDefault("LOG_MAX_SIZE", 100)  // 100MB
	viper.SetDefault("LOG_MAX_BACKUPS", 3) // keep 3 old log files
	viper.SetDefault("LOG_MAX_AGE", 28)    // 28 days
	viper.SetDefault("LOG_COMPRESS", true) // compress old log files

	// Rate limit defaults
	viper.SetDefault("RATE_LIMIT_RPS", 100)
	viper.SetDefault("RATE_LIMIT_BURST", 50)

	// CORS defaults
	viper.SetDefault("CORS_ALLOWED_ORIGINS", "*")
	viper.SetDefault("CORS_ALLOWED_METHODS", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	viper.SetDefault("CORS_ALLOWED_HEADERS", "Origin,Content-Type,Accept,Authorization")
	viper.SetDefault("CORS_MAX_AGE", 86400)

	// Redis defaults
	viper.SetDefault("REDIS_MODE", "standalone")
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", 6379)
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("REDIS_POOL_SIZE", 100)
	viper.SetDefault("REDIS_MIN_IDLE_CONNS", 10)
	viper.SetDefault("REDIS_MAX_RETRIES", 3)
	viper.SetDefault("REDIS_DIAL_TIMEOUT", "5s")
	viper.SetDefault("REDIS_READ_TIMEOUT", "3s")
	viper.SetDefault("REDIS_WRITE_TIMEOUT", "3s")
	viper.SetDefault("REDIS_POOL_TIMEOUT", "4s")
}

// validate checks required configuration fields
func validate(cfg *setting.Config) error {
	if cfg.Database.User == "" {
		return fmt.Errorf("DB_USER is required")
	}
	if cfg.Database.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}
	if cfg.Database.Name == "" {
		return fmt.Errorf("DB_NAME is required")
	}
	if cfg.JWT.Secret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}
	return nil
}

// IsDevelopment returns true if running in development mode
func IsDevelopment() bool {
	return global.Config.App.Env == "development"
}

