package dto

import (
	"time"

	"rental-management-api/internal/constant"
)

type CreateCustomerLogRequest struct {
	CustomerID   uint                       `json:"customer_id" binding:"required"`
	CustomerName string                     `json:"customer_name" binding:"required"`
	Reason       string                     `json:"reason"`
	Status       constant.CustomerLogStatus `json:"status" binding:"required"`
}

type UpdateCustomerLogRequest struct {
	CustomerID   *uint                       `json:"customer_id"`
	CustomerName *string                     `json:"customer_name"`
	Reason       *string                     `json:"reason"`
	Status       *constant.CustomerLogStatus `json:"status"`
}

type CustomerLogResource struct {
	ID           uint                       `json:"id"`
	CustomerID   uint                       `json:"customer_id"`
	CustomerName string                     `json:"customer_name"`
	Reason       string                     `json:"reason"`
	Status       constant.CustomerLogStatus `json:"status"`
	CreatedAt    time.Time                  `json:"created_at"`
	UpdatedAt    time.Time                  `json:"updated_at"`
}
