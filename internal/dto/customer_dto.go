package dto

import "time"

type CreateCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address"`
	AvatarURL   string `json:"avatar_url"`
}

type UpdateCustomerRequest struct {
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	Password    *string `json:"password"`
	PhoneNumber *string `json:"phone_number"`
	Address     *string `json:"address"`
	AvatarURL   *string `json:"avatar_url"`
}

type CustomerResource struct {
	ID           uint      `json:"id"`
	UserName     string    `json:"name"`
	UserEmail    string    `json:"email"`
	UserPassword string    `json:"password"`
	PhoneNumber  string    `json:"phone_number"`
	Address      string    `json:"address"`
	AvatarURL    string    `json:"avatar_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
