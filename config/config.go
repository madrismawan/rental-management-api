package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBDriver      string
	DBDSN         string
	AdminPassword string
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		Port:          getEnv("HTTP_PORT", getEnv("PORT", "8080")),
		DBDriver:      getEnv("DB_DRIVER", "postgres"),
		DBDSN:         getEnv("DATABASE_URL", getEnv("DB_DSN", "host=localhost user=postgres password=postgres dbname=rental_management port=5432 sslmode=disable TimeZone=Asia/Jakarta")),
		AdminPassword: getEnv("ADMIN_PASSWORD", "admin123"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
