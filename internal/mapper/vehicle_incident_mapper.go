package mapper

import (
	"fmt"
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

func ToVehicleIncidentResource(incident entity.VehicleIncident) dto.VehicleIncidentResource {
	return dto.VehicleIncidentResource{
		ID:           incident.ID,
		VehicleID:    incident.VehicleID,
		VehicleName:  fmt.Sprintf("%s %s (%s)", incident.Vehicle.Brand, incident.Vehicle.Model, incident.Vehicle.PlateNumber),
		CustomerID:   incident.CustomerID,
		RentalID:     incident.RentalID,
		IncidentDate: incident.IncidentDate,
		IncidentType: incident.IncidentType,
		Description:  incident.Description,
		Cost:         incident.Cost,
		Status:       incident.Status,
		CreatedAt:    incident.CreatedAt,
		UpdatedAt:    incident.UpdatedAt,
	}
}

func ToVehicleIncidentsResource(incidents []entity.VehicleIncident) []dto.VehicleIncidentResource {
	out := make([]dto.VehicleIncidentResource, 0, len(incidents))
	for _, item := range incidents {
		out = append(out, ToVehicleIncidentResource(item))
	}
	return out
}
