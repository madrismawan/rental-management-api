package entity

import (
	"rental-management-api/internal/constant"
	"time"
)

type CustomerLog struct {
	ID           uint                       `gorm:"primaryKey;column:id"`
	CustomerID   uint                       `gorm:"column:customer_id;not null;index"`
	CustomerName string                     `gorm:"column:customer_name;type:varchar(100);not null"`
	Reason       string                     `gorm:"column:reason;type:text"`
	Status       constant.CustomerLogStatus `gorm:"column:status;type:varchar(20);not null"`
	CreatedAt    time.Time                  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time                  `gorm:"column:updated_at;autoUpdateTime"`

	Customer Customer `gorm:"foreignKey:CustomerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (CustomerLog) TableName() string {
	return "customer_logs"
}
