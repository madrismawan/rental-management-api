package dto

import (
	"rental-management-api/internal/constant"
	"time"
)

type CreateRentalRequest struct {
	CustomerID            uint                      `json:"customer_id" binding:"required"`
	VehicleID             uint                      `json:"vehicle_id" binding:"required"`
	StartDate             time.Time                 `json:"start_date" binding:"required"`
	EndDate               time.Time                 `json:"end_date" binding:"required"`
	Notes                 string                    `json:"notes"`
	VehicleConditionStart constant.VehicleCondition `json:"vehicle_condition_start"`
	MileageStart          int                       `json:"mileage_start"`
}

type UpdateRentalRequest struct {
	CustomerID            *uint                      `json:"customer_id"`
	VehicleID             *uint                      `json:"vehicle_id"`
	StartDate             *time.Time                 `json:"start_date"`
	EndDate               *time.Time                 `json:"end_date"`
	TotalDay              *int                       `json:"total_day"`
	ReturnDate            *time.Time                 `json:"return_date"`
	Price                 *int64                     `json:"price"`
	PenaltyFee            *int64                     `json:"penalty_fee"`
	Subtotal              *int64                     `json:"subtotal"`
	Notes                 *string                    `json:"notes"`
	Status                *constant.RentalStatus     `json:"status"`
	VehicleConditionStart *constant.VehicleCondition `json:"vehicle_condition_start"`
	VehicleConditionEnd   *constant.VehicleCondition `json:"vehicle_condition_end"`
	MileageStart          *int                       `json:"mileage_start"`
	MileageUsed           *int                       `json:"mileage_used"`
	MileageEnd            *int                       `json:"mileage_end"`
}

type RentalResource struct {
	ID                    uint                      `json:"id"`
	NoInvoice             string                    `json:"no_invoice"`
	CustomerID            uint                      `json:"customer_id"`
	CustomerName          string                    `json:"customer_name"`
	VehicleID             uint                      `json:"vehicle_id"`
	VehicleName           string                    `json:"vehicle_name"`
	StartDate             time.Time                 `json:"start_date"`
	EndDate               time.Time                 `json:"end_date"`
	TotalDay              int                       `json:"total_day"`
	ReturnDate            *time.Time                `json:"return_date"`
	Price                 int64                     `json:"price"`
	PenaltyFee            int64                     `json:"penalty_fee"`
	Subtotal              int64                     `json:"subtotal"`
	Notes                 string                    `json:"notes"`
	Status                constant.RentalStatus     `json:"status"`
	VehicleConditionStart constant.VehicleCondition `json:"vehicle_condition_start"`
	VehicleConditionEnd   constant.VehicleCondition `json:"vehicle_condition_end"`
	MileageStart          int                       `json:"mileage_start"`
	MileageUsed           int                       `json:"mileage_used"`
	MileageEnd            int                       `json:"mileage_end"`
	CreatedAt             time.Time                 `json:"created_at"`
	UpdatedAt             time.Time                 `json:"updated_at"`
}

type RentalOptionResource struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CompleteRentalRequest struct {
	ReturnDate          time.Time                 `json:"return_date" binding:"required"`
	PenaltyFee          int64                     `json:"penalty_fee"`
	IncidentType        string                    `json:"incident_type"`
	Description         string                    `json:"description"`
	VehicleConditionEnd constant.VehicleCondition `json:"vehicle_condition_end" binding:"required"`
	MileageEnd          int                       `json:"mileage_end" binding:"required"`
}
