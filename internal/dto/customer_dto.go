package dto

import (
	"mime/multipart"
	"time"

	"rental-management-api/internal/constant"
)

type CreateCustomerRequest struct {
	Name        string                `json:"name" form:"name" binding:"required"`
	Email       string                `json:"email" form:"email" binding:"required"`
	Password    string                `json:"password" form:"password" binding:"required"`
	PhoneNumber string                `json:"phone_number" form:"phone_number" binding:"required"`
	Address     string                `json:"address" form:"address"`
	Status      constant.CustomerStatus `json:"status" form:"status" binding:"required"`
	Avatar      *multipart.FileHeader `form:"avatar"`
}

type UpdateCustomerRequest struct {
	Name        *string               `json:"name" form:"name"`
	Email       *string               `json:"email" form:"email"`
	Password    *string               `json:"password" form:"password"`
	PhoneNumber *string               `json:"phone_number" form:"phone_number"`
	Address     *string               `json:"address" form:"address"`
	Status      *constant.CustomerStatus `json:"status" form:"status"`
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
	Status      constant.CustomerStatus `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CustomerOptionResource struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
