package database

import (
	"rental-management-api/internal/entity"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Customer{},
		&entity.CustomerLog{},
		&entity.Vehicle{},
		&entity.Rental{},
		&entity.VehicleIncident{},
	)
}
