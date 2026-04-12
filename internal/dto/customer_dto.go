package dto

import "time"

type CreateCustomerRequest struct {
	UserID      uint   `json:"user_id" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address"`
	AvatarURL   string `json:"avatar_url"`
}

type UpdateCustomerRequest struct {
	UserID      *uint   `json:"user_id"`
	PhoneNumber *string `json:"phone_number"`
	Address     *string `json:"address"`
	AvatarURL   *string `json:"avatar_url"`
}

type CustomerResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	AvatarURL   string    `json:"avatar_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
