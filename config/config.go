package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	defaultAppName         = "basudeva-api"
	defaultEnvironment     = "development"
	defaultHost            = "0.0.0.0"
	defaultPort            = 8080
	defaultLogLevel        = "info"
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 10 * time.Second
	defaultShutdownTimeout = 10 * time.Second
	defaultDatabaseURL     = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	defaultAuthSecret      = "changeme"
	defaultTokenTTL        = time.Hour
)

type Config struct {
	AppName     string
	Environment string
	Logger      LoggerConfig
	HTTP        HTTPConfig
	Shutdown    ShutdownConfig
	Database    DatabaseConfig
	Auth        AuthConfig
}

type LoggerConfig struct {
	Level string
}

type HTTPConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type ShutdownConfig struct {
	Timeout time.Duration
}

type DatabaseConfig struct {
	URL string
}

type AuthConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	TokenTTL           time.Duration
}

func MustLoad() Config {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}
	return cfg
}

func Load() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		AppName:     getString("APP_NAME", defaultAppName),
		Environment: getString("APP_ENV", defaultEnvironment),
		Logger: LoggerConfig{
			Level: getString("LOG_LEVEL", defaultLogLevel),
		},
		HTTP: HTTPConfig{
			Host:         getString("HTTP_HOST", defaultHost),
			Port:         getInt("HTTP_PORT", defaultPort),
			ReadTimeout:  getDuration("HTTP_READ_TIMEOUT", defaultReadTimeout),
			WriteTimeout: getDuration("HTTP_WRITE_TIMEOUT", defaultWriteTimeout),
		},
		Shutdown: ShutdownConfig{
			Timeout: getDuration("SHUTDOWN_TIMEOUT", defaultShutdownTimeout),
		},
		Database: DatabaseConfig{
			URL: getString("DATABASE_URL", defaultDatabaseURL),
		},
		Auth: AuthConfig{
			AccessTokenSecret:  getString("AUTH_ACCESS_TOKEN_SECRET", defaultAuthSecret),
			RefreshTokenSecret: getString("AUTH_REFRESH_TOKEN_SECRET", defaultAuthSecret),
			TokenTTL:           getDuration("AUTH_TOKEN_TTL", defaultTokenTTL),
		},
	}

	return cfg, nil
}

func (c HTTPConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func getString(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func getDuration(key string, fallback time.Duration) time.Duration {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return duration
}
