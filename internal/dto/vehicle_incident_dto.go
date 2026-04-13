package dto

import (
	"time"

	"rental-management-api/internal/constant"
)

type CreateVehicleIncidentRequest struct {
	VehicleID    uint                           `json:"vehicle_id" binding:"required"`
	CustomerID   *uint                          `json:"customer_id"`
	RentalID     *uint                          `json:"rental_id"`
	IncidentDate time.Time                      `json:"incident_date" binding:"required"`
	IncidentType constant.IncidentType          `json:"incident_type" binding:"required"`
	Description  string                         `json:"description"`
	PenaltyFee   int64                          `json:"penalty_fee"`
	Status       constant.VehicleIncidentStatus `json:"status" binding:"required"`
}

type UpdateVehicleIncidentRequest struct {
	VehicleID    *uint                           `json:"vehicle_id"`
	CustomerID   *uint                           `json:"customer_id"`
	RentalID     *uint                           `json:"rental_id"`
	IncidentDate *time.Time                      `json:"incident_date"`
	IncidentType *constant.IncidentType          `json:"incident_type"`
	Description  *string                         `json:"description"`
	PenaltyFee   *int64                          `json:"penalty_fee"`
	Status       *constant.VehicleIncidentStatus `json:"status"`
}

type VehicleIncidentResource struct {
	ID           uint                           `json:"id"`
	VehicleID    uint                           `json:"vehicle_id"`
	CustomerID   *uint                          `json:"customer_id"`
	RentalID     *uint                          `json:"rental_id"`
	IncidentDate time.Time                      `json:"incident_date"`
	IncidentType constant.IncidentType          `json:"incident_type"`
	Description  string                         `json:"description"`
	PenaltyFee   int64                          `json:"penalty_fee"`
	Status       constant.VehicleIncidentStatus `json:"status"`
	CreatedAt    time.Time                      `json:"created_at"`
	UpdatedAt    time.Time                      `json:"updated_at"`
}
