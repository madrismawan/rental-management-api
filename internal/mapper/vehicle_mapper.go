package mapper

import (
	"fmt"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

func ToVehicleEntity(req dto.CreateVehicleRequest) entity.Vehicle {
	return entity.Vehicle{
		PlateNumber: req.PlateNumber,
		Color:       req.Color,
		Brand:       req.Brand,
		Model:       req.Model,
		CC:          req.CC,
		Year:        req.Year,
		Mileage:     req.Mileage,
		DailyRate:   req.DailyRate,
		Condition:   req.Condition,
		Status:      req.Status,
		Notes:       req.Notes,
	}
}

func ApplyVehicleUpdate(vehicle *entity.Vehicle, req dto.UpdateVehicleRequest) {
	if req.PlateNumber != nil {
		vehicle.PlateNumber = *req.PlateNumber
	}
	if req.Color != nil {
		vehicle.Color = *req.Color
	}
	if req.Brand != nil {
		vehicle.Brand = *req.Brand
	}
	if req.Model != nil {
		vehicle.Model = *req.Model
	}
	if req.CC != nil {
		vehicle.CC = *req.CC
	}
	if req.Year != nil {
		vehicle.Year = *req.Year
	}
	if req.Mileage != nil {
		vehicle.Mileage = *req.Mileage
	}
	if req.DailyRate != nil {
		vehicle.DailyRate = *req.DailyRate
	}
	if req.Condition != nil {
		vehicle.Condition = *req.Condition
	}
	if req.Status != nil {
		vehicle.Status = *req.Status
	}
	if req.Notes != nil {
		vehicle.Notes = *req.Notes
	}
}

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
		Conditon:  vehicle.Condition,
	}
}

func ToVehicleOptionsResource(vehicles []entity.Vehicle) []dto.VehicleOptionResource {
	out := make([]dto.VehicleOptionResource, 0, len(vehicles))
	for _, item := range vehicles {
		out = append(out, ToVehicleOptionResource(item))
	}
	return out
}
