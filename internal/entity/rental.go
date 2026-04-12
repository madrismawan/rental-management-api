package entity

import "time"

type RentalStatus string

const (
	RentalStatusPending   RentalStatus = "pending"
	RentalStatusActive    RentalStatus = "active"
	RentalStatusCompleted RentalStatus = "completed"
	RentalStatusCanceled  RentalStatus = "canceled"
)

type Rental struct {
	ID                    uint         `gorm:"primaryKey;column:id"`
	CustomerID            uint         `gorm:"column:customer_id;not null;index"`
	VehicleID             uint         `gorm:"column:vehicle_id;not null;index"`
	StartDate             time.Time    `gorm:"column:start_date;not null"`
	EndDate               time.Time    `gorm:"column:end_date;not null"`
	TotalDay              int          `gorm:"column:total_day;not null"`
	ReturnDate            *time.Time   `gorm:"column:return_date"`
	Price                 int64        `gorm:"column:price;not null"`
	PenaltyFee            int64        `gorm:"column:penalty_fee;not null;default:0"`
	Subtotal              int64        `gorm:"column:subtotal;not null"`
	Notes                 string       `gorm:"column:notes;type:text"`
	Status                RentalStatus `gorm:"column:status;type:varchar(20);not null;default:'pending'"`
	VehicleConditionStart string       `gorm:"column:vehicle_condition_start;type:text"`
	VehicleConditionEnd   string       `gorm:"column:vehicle_condition_end;type:text"`
	MileageStart          int          `gorm:"column:mileage_start;not null;default:0"`
	MileageUsed           int          `gorm:"column:mileage_used;not null;default:0"`
	MileageEnd            int          `gorm:"column:mileage_end;not null;default:0"`
	CreatedAt             time.Time    `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt             time.Time    `gorm:"column:updated_at;autoUpdateTime"`

	Customer         Customer          `gorm:"foreignKey:CustomerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Vehicle          Vehicle           `gorm:"foreignKey:VehicleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	VehicleIncidents []VehicleIncident `gorm:"foreignKey:RentalID;references:ID"`
}

func (Rental) TableName() string {
	return "rental"
}
