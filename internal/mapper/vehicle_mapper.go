package mapper

import (
	"fmt"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

func ToVehicleResource(vehicle entity.Vehicle) dto.VehicleResource {
	return dto.VehicleResource{
		ID:          vehicle.ID,
		PlateNumber: vehicle.PlateNumber,
		Color:       vehicle.Color,
		Brand:       vehicle.Brand,
		Model:       vehicle.Model,
		CC:          vehicle.CC,
		Year:        vehicle.Year,
		Mileage:     vehicle.Mileage,
		DailyRate:   vehicle.DailyRate,
		Condition:   vehicle.Condition,
		Status:      vehicle.Status,
		Notes:       vehicle.Notes,
		CreatedAt:   vehicle.CreatedAt,
		UpdatedAt:   vehicle.UpdatedAt,
	}
}

func ToVehiclesResource(vehicles []entity.Vehicle) []dto.VehicleResource {
	out := make([]dto.VehicleResource, 0, len(vehicles))
	for _, item := range vehicles {
		out = append(out, ToVehicleResource(item))
	}
	return out
}

func ToVehicleOptionResource(vehicle entity.Vehicle) dto.VehicleOptionResource {
	return dto.VehicleOptionResource{
		ID:        vehicle.ID,
		Name:      fmt.Sprintf("%s %s (%s)", vehicle.Brand, vehicle.Model, vehicle.PlateNumber),
		DailyRate: vehicle.DailyRate,
		Mileage:   vehicle.Mileage,
		Condition: vehicle.Condition,
	}
}

func ToVehicleOptionsResource(vehicles []entity.Vehicle) []dto.VehicleOptionResource {
	out := make([]dto.VehicleOptionResource, 0, len(vehicles))
	for _, item := range vehicles {
		out = append(out, ToVehicleOptionResource(item))
	}
	return out
}
