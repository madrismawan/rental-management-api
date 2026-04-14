package dto

import (
	"mime/multipart"
	"time"
)

type CreateCustomerRequest struct {
	Name        string                `json:"name" form:"name" binding:"required"`
	Email       string                `json:"email" form:"email" binding:"required"`
	Password    string                `json:"password" form:"password" binding:"required"`
	PhoneNumber string                `json:"phone_number" form:"phone_number" binding:"required"`
	Address     string                `json:"address" form:"address"`
	Avatar      *multipart.FileHeader `form:"avatar"`
}

type UpdateCustomerRequest struct {
	Name        *string               `json:"name" form:"name"`
	Email       *string               `json:"email" form:"email"`
	Password    *string               `json:"password" form:"password"`
	PhoneNumber *string               `json:"phone_number" form:"phone_number"`
	Address     *string               `json:"address" form:"address"`
	AvatarURL   *string               `json:"avatar_url" form:"avatar_url"`
	Avatar      *multipart.FileHeader `form:"avatar"`
}

type CustomerResource struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	AvatarURL   string    `json:"avatar_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CustomerOptionResource struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
