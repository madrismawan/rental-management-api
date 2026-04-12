package database

import (
	"fmt"
	"strings"

	"rental-management-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg config.Config) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch strings.ToLower(cfg.DBDriver) {
	case "postgres", "postgresql":
		dialector = postgres.Open(cfg.DBDSN)
	case "sqlite", "sqlite3", "":
		dialector = sqlite.Open(cfg.DBDSN)
	default:
		return nil, fmt.Errorf("unsupported DB_DRIVER %q", cfg.DBDriver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := Migrate(db); err != nil {
		return nil, err
	}

	if err := SeedAdminUser(db, cfg); err != nil {
		return nil, err
	}

	return db, nil
}
