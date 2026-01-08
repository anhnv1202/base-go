package setting

import (
	"fmt"
	"strings"
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
	Redis     RedisConfig     `mapstructure:",squash"`
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

// DSNWithoutDB returns DSN connecting to postgres system database (for creating database)
func (d *DatabaseConfig) DSNWithoutDB() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=postgres sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.SSLMode,
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

// RedisConfig holds Redis connection settings
type RedisConfig struct {
	Mode             string        `mapstructure:"REDIS_MODE"`
	Host             string        `mapstructure:"REDIS_HOST"`
	Port             int           `mapstructure:"REDIS_PORT"`
	Password         string        `mapstructure:"REDIS_PASSWORD"`
	DB               int           `mapstructure:"REDIS_DB"`
	MasterName       string        `mapstructure:"REDIS_MASTER_NAME"`
	SentinelAddrs    string        `mapstructure:"REDIS_SENTINEL_ADDRS"`
	SentinelPassword string        `mapstructure:"REDIS_SENTINEL_PASSWORD"`
	PoolSize         int           `mapstructure:"REDIS_POOL_SIZE"`
	MinIdleConns     int           `mapstructure:"REDIS_MIN_IDLE_CONNS"`
	MaxRetries       int           `mapstructure:"REDIS_MAX_RETRIES"`
	DialTimeout      time.Duration `mapstructure:"REDIS_DIAL_TIMEOUT"`
	ReadTimeout      time.Duration `mapstructure:"REDIS_READ_TIMEOUT"`
	WriteTimeout     time.Duration `mapstructure:"REDIS_WRITE_TIMEOUT"`
	PoolTimeout      time.Duration `mapstructure:"REDIS_POOL_TIMEOUT"`
}

func (r *RedisConfig) IsSentinelMode() bool {
	return strings.ToLower(r.Mode) == "sentinel"
}

func (r *RedisConfig) GetSentinelAddrs() []string {
	if r.SentinelAddrs == "" {
		return nil
	}
	var result []string
	for _, addr := range strings.Split(r.SentinelAddrs, ",") {
		if s := strings.TrimSpace(addr); s != "" {
			result = append(result, s)
		}
	}
	return result
}

func (r *RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}