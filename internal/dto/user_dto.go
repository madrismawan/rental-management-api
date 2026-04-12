package dto

import (
	"rental-management-api/internal/constant"
	"time"
)

type CreateUserRequest struct {
	Name     string            `json:"name" binding:"required"`
	Email    string            `json:"email" binding:"required,email"`
	Role     constant.UserRole `json:"role" binding:"required"`
	Password string            `json:"password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
	Name     *string            `json:"name"`
	Email    *string            `json:"email"`
	Role     *constant.UserRole `json:"role"`
	Password *string            `json:"password"`
}

type UserResponse struct {
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Role      constant.UserRole `json:"role"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
