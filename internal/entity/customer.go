package entity

import "time"

type Customer struct {
	ID          uint      `gorm:"primaryKey;column:id"`
	UserID      uint      `gorm:"column:user_id;not null;index"`
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(20);not null"`
	Address     string    `gorm:"column:address;type:text"`
	AvatarURL   string    `gorm:"column:avatar_url;type:text"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`

	User             User              `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Rentals          []Rental          `gorm:"foreignKey:CustomerID;references:ID"`
	VehicleIncidents []VehicleIncident `gorm:"foreignKey:CustomerID;references:ID"`
}

func (Customer) TableName() string {
	return "customers"
}
