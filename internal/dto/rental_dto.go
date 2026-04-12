package dto

import (
	"time"

	"rental-management-api/internal/entity"
)

type CreateRentalRequest struct {
	CustomerID            uint                `json:"customer_id" binding:"required"`
	VehicleID             uint                `json:"vehicle_id" binding:"required"`
	StartDate             time.Time           `json:"start_date" binding:"required"`
	EndDate               time.Time           `json:"end_date" binding:"required"`
	TotalDay              int                 `json:"total_day" binding:"required"`
	ReturnDate            *time.Time          `json:"return_date"`
	Price                 int64               `json:"price" binding:"required"`
	PenaltyFee            int64               `json:"penalty_fee"`
	Subtotal              int64               `json:"subtotal" binding:"required"`
	Notes                 string              `json:"notes"`
	Status                entity.RentalStatus `json:"status" binding:"required"`
	VehicleConditionStart string              `json:"vehicle_condition_start"`
	VehicleConditionEnd   string              `json:"vehicle_condition_end"`
	MileageStart          int                 `json:"mileage_start"`
	MileageUsed           int                 `json:"mileage_used"`
	MileageEnd            int                 `json:"mileage_end"`
}

type UpdateRentalRequest struct {
	CustomerID            *uint                `json:"customer_id"`
	VehicleID             *uint                `json:"vehicle_id"`
	StartDate             *time.Time           `json:"start_date"`
	EndDate               *time.Time           `json:"end_date"`
	TotalDay              *int                 `json:"total_day"`
	ReturnDate            *time.Time           `json:"return_date"`
	Price                 *int64               `json:"price"`
	PenaltyFee            *int64               `json:"penalty_fee"`
	Subtotal              *int64               `json:"subtotal"`
	Notes                 *string              `json:"notes"`
	Status                *entity.RentalStatus `json:"status"`
	VehicleConditionStart *string              `json:"vehicle_condition_start"`
	VehicleConditionEnd   *string              `json:"vehicle_condition_end"`
	MileageStart          *int                 `json:"mileage_start"`
	MileageUsed           *int                 `json:"mileage_used"`
	MileageEnd            *int                 `json:"mileage_end"`
}

type RentalResponse struct {
	ID                    uint                `json:"id"`
	CustomerID            uint                `json:"customer_id"`
	VehicleID             uint                `json:"vehicle_id"`
	StartDate             time.Time           `json:"start_date"`
	EndDate               time.Time           `json:"end_date"`
	TotalDay              int                 `json:"total_day"`
	ReturnDate            *time.Time          `json:"return_date"`
	Price                 int64               `json:"price"`
	PenaltyFee            int64               `json:"penalty_fee"`
	Subtotal              int64               `json:"subtotal"`
	Notes                 string              `json:"notes"`
	Status                entity.RentalStatus `json:"status"`
	VehicleConditionStart string              `json:"vehicle_condition_start"`
	VehicleConditionEnd   string              `json:"vehicle_condition_end"`
	MileageStart          int                 `json:"mileage_start"`
	MileageUsed           int                 `json:"mileage_used"`
	MileageEnd            int                 `json:"mileage_end"`
	CreatedAt             time.Time           `json:"created_at"`
	UpdatedAt             time.Time           `json:"updated_at"`
}
