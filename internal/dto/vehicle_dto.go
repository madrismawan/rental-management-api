package dto

import (
	"time"

	"rental-management-api/internal/constant"
)

type CreateVehicleRequest struct {
	PlateNumber string                 `json:"plate_number" binding:"required"`
	Color       string                 `json:"color"`
	Brand       string                 `json:"brand" binding:"required"`
	Model       string                 `json:"model" binding:"required"`
	CC          int                    `json:"cc" binding:"required"`
	Year        int                    `json:"year" binding:"required"`
	Mileage     int                    `json:"mileage"`
	DailyRate   int64                  `json:"daily_rate" binding:"required"`
	Status      constant.VehicleStatus `json:"status" binding:"required"`
	Notes       string                 `json:"notes"`
}

type UpdateVehicleRequest struct {
	PlateNumber *string                 `json:"plate_number"`
	Color       *string                 `json:"color"`
	Brand       *string                 `json:"brand"`
	Model       *string                 `json:"model"`
	CC          *int                    `json:"cc"`
	Year        *int                    `json:"year"`
	Mileage     *int                    `json:"mileage"`
	DailyRate   *int64                  `json:"daily_rate"`
	Status      *constant.VehicleStatus `json:"status"`
	Notes       *string                 `json:"notes"`
}

type VehicleResponse struct {
	ID          uint                   `json:"id"`
	PlateNumber string                 `json:"plate_number"`
	Color       string                 `json:"color"`
	Brand       string                 `json:"brand"`
	Model       string                 `json:"model"`
	CC          int                    `json:"cc"`
	Year        int                    `json:"year"`
	Mileage     int                    `json:"mileage"`
	DailyRate   int64                  `json:"daily_rate"`
	Status      constant.VehicleStatus `json:"status"`
	Notes       string                 `json:"notes"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}
