package entity

import (
	"rental-management-api/internal/constant"
	"time"
)

type VehicleIncident struct {
	ID           uint                           `gorm:"primaryKey;column:id"`
	VehicleID    uint                           `gorm:"column:vehicle_id;not null;index"`
	CustomerID   uint                           `gorm:"column:customer_id;not null;index"`
	RentalID     uint                           `gorm:"column:rental_id;not null;index"`
	IncidentDate time.Time                      `gorm:"column:incident_date;not null"`
	IncidentType constant.IncidentType          `gorm:"column:incident_type;type:varchar(100);not null"`
	Description  string                         `gorm:"column:description;type:text"`
	PenaltyFee   int64                          `gorm:"column:penalty_fee;not null;default:0"`
	Status       constant.VehicleIncidentStatus `gorm:"column:status;type:varchar(20);not null;default:'open'"`
	CreatedAt    time.Time                      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time                      `gorm:"column:updated_at;autoUpdateTime"`

	Vehicle  Vehicle  `gorm:"foreignKey:VehicleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Customer Customer `gorm:"foreignKey:CustomerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Rental   Rental   `gorm:"foreignKey:RentalID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (VehicleIncident) TableName() string {
	return "vehicle_incidents"
}
