package dto

import (
	"rental-management-api/internal/constant"
	"time"
)

type UserResource struct {
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Role      constant.UserRole `json:"role"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
