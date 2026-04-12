package entity

import (
	"rental-management-api/internal/constant"
	"time"
)

type User struct {
	ID        uint              `gorm:"primaryKey;column:id"`
	Name      string            `gorm:"column:name;type:varchar(100);not null"`
	Email     string            `gorm:"column:email;type:varchar(150);not null;uniqueIndex"`
	Role      constant.UserRole `gorm:"column:role;type:varchar(20);not null;default:'customer'"`
	Password  string            `gorm:"column:password;type:varchar(255);not null"`
	CreatedAt time.Time         `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time         `gorm:"column:updated_at;autoUpdateTime"`

	Customers []Customer `gorm:"foreignKey:UserID;references:ID"`
}

func (User) TableName() string {
	return "users"
}
