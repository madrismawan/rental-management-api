package database

import (
	"fmt"

	"gorm.io/gorm"

	"rental-management-api/internal/entity"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(entity.Models()...); err != nil {
		return fmt.Errorf("auto migrate: %w", err)
	}
	return nil
}
