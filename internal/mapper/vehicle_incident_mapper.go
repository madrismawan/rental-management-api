package mapper

import (
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

func ToVehicleIncidentEntity(req dto.CreateVehicleIncidentRequest) entity.VehicleIncident {
	return entity.VehicleIncident{
		VehicleID:    req.VehicleID,
		CustomerID:   req.CustomerID,
		RentalID:     req.RentalID,
		IncidentDate: req.IncidentDate,
		IncidentType: req.IncidentType,
		Description:  req.Description,
		PenaltyFee:   req.PenaltyFee,
		Status:       req.Status,
	}
}

func ApplyVehicleIncidentUpdate(incident *entity.VehicleIncident, req dto.UpdateVehicleIncidentRequest) {
	if req.VehicleID != nil {
		incident.VehicleID = *req.VehicleID
	}
	if req.CustomerID != nil {
		incident.CustomerID = *req.CustomerID
	}
	if req.RentalID != nil {
		incident.RentalID = *req.RentalID
	}
	if req.IncidentDate != nil {
		incident.IncidentDate = *req.IncidentDate
	}
	if req.IncidentType != nil {
		incident.IncidentType = *req.IncidentType
	}
	if req.Description != nil {
		incident.Description = *req.Description
	}
	if req.PenaltyFee != nil {
		incident.PenaltyFee = *req.PenaltyFee
	}
	if req.Status != nil {
		incident.Status = *req.Status
	}
}

func ToVehicleIncidentResource(incident entity.VehicleIncident) dto.VehicleIncidentResource {
	return dto.VehicleIncidentResource{
		ID:           incident.ID,
		VehicleID:    incident.VehicleID,
		CustomerID:   incident.CustomerID,
		RentalID:     incident.RentalID,
		IncidentDate: incident.IncidentDate,
		IncidentType: incident.IncidentType,
		Description:  incident.Description,
		PenaltyFee:   incident.PenaltyFee,
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
