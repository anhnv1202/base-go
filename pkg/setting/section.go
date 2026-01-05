package setting

import (
	"fmt"
	"time"
)

// Config holds all application configuration
type Config struct {
	App       AppConfig       `mapstructure:",squash"`
	Database  DatabaseConfig  `mapstructure:",squash"`
	JWT       JWTConfig       `mapstructure:",squash"`
	Log       LogConfig       `mapstructure:",squash"`
	RateLimit RateLimitConfig `mapstructure:",squash"`
	CORS      CORSConfig      `mapstructure:",squash"`
}

// AppConfig holds application-specific settings
type AppConfig struct {
	Env  string `mapstructure:"APP_ENV"`
	Port int    `mapstructure:"APP_PORT"`
	Name string `mapstructure:"APP_NAME"`
}

// DatabaseConfig holds database connection settings
type DatabaseConfig struct {
	Host            string        `mapstructure:"DB_HOST"`
	Port            int           `mapstructure:"DB_PORT"`
	User            string        `mapstructure:"DB_USER"`
	Password        string        `mapstructure:"DB_PASSWORD"`
	Name            string        `mapstructure:"DB_NAME"`
	SSLMode         string        `mapstructure:"DB_SSL_MODE"`
	MaxIdleConns    int           `mapstructure:"DB_MAX_IDLE_CONNS"`
	MaxOpenConns    int           `mapstructure:"DB_MAX_OPEN_CONNS"`
	ConnMaxLifetime time.Duration `mapstructure:"DB_CONN_MAX_LIFETIME"`
}

// DSN returns the PostgreSQL connection string
func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode,
	)
}

// JWTConfig holds JWT authentication settings
type JWTConfig struct {
	Secret        string        `mapstructure:"JWT_SECRET"`
	Expiry        time.Duration `mapstructure:"JWT_EXPIRY"`
	RefreshExpiry time.Duration `mapstructure:"JWT_REFRESH_EXPIRY"`
}

// LogConfig holds logging settings
type LogConfig struct {
	Level      string `mapstructure:"LOG_LEVEL"`
	Format     string `mapstructure:"LOG_FORMAT"`
	FilePath   string `mapstructure:"LOG_FILE_PATH"`
	MaxSize    int    `mapstructure:"LOG_MAX_SIZE"`    // megabytes
	MaxBackups int    `mapstructure:"LOG_MAX_BACKUPS"` // number of old log files
	MaxAge     int    `mapstructure:"LOG_MAX_AGE"`     // days
	Compress   bool   `mapstructure:"LOG_COMPRESS"`    // compress old log files
}

// RateLimitConfig holds rate limiting settings
type RateLimitConfig struct {
	RPS   int `mapstructure:"RATE_LIMIT_RPS"`
	Burst int `mapstructure:"RATE_LIMIT_BURST"`
}

// CORSConfig holds CORS settings
type CORSConfig struct {
	AllowedOrigins string `mapstructure:"CORS_ALLOWED_ORIGINS"`
	AllowedMethods string `mapstructure:"CORS_ALLOWED_METHODS"`
	AllowedHeaders string `mapstructure:"CORS_ALLOWED_HEADERS"`
	MaxAge         int    `mapstructure:"CORS_MAX_AGE"`
}