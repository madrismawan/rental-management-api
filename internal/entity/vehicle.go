package entity

import (
	"rental-management-api/internal/constant"
	"time"
)

type Vehicle struct {
	ID          uint                   `gorm:"primaryKey;column:id"`
	PlateNumber string                 `gorm:"column:plate_number;type:varchar(20);not null;uniqueIndex"`
	Color       string                 `gorm:"column:color;type:varchar(50)"`
	Brand       string                 `gorm:"column:brand;type:varchar(100);not null"`
	Model       string                 `gorm:"column:model;type:varchar(100);not null"`
	CC          int                    `gorm:"column:cc;not null"`
	Year        int                    `gorm:"column:year;not null"`
	Mileage     int                    `gorm:"column:mileage;not null;default:0"`
	DailyRate   int64                  `gorm:"column:daily_rate;not null"`
	Status      constant.VehicleStatus `gorm:"column:status;type:varchar(20);not null;default:'available'"`
	Notes       string                 `gorm:"column:notes;type:text"`
	CreatedAt   time.Time              `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time              `gorm:"column:updated_at;autoUpdateTime"`

	Rentals          []Rental          `gorm:"foreignKey:VehicleID;references:ID"`
	VehicleIncidents []VehicleIncident `gorm:"foreignKey:VehicleID;references:ID"`
}

func (Vehicle) TableName() string {
	return "vehicle"
}
