package service

import (
	"context"
	"time"

	"rental-management-api/internal/constant"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"
)

type VehicleIncidentService interface {
	Create(ctx context.Context, data CreateVehicleIncidentInput) (*entity.VehicleIncident, error)
	GetByID(ctx context.Context, id uint) (*entity.VehicleIncident, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.VehicleIncident, error)
	List(ctx context.Context) ([]entity.VehicleIncident, error)
	Update(ctx context.Context, id uint, data UpdateVehicleIncidentInput) (*entity.VehicleIncident, error)
	Delete(ctx context.Context, id uint) error
}

type CreateVehicleIncidentInput struct {
	VehicleID    uint
	CustomerID   uint
	RentalID     uint
	IncidentDate time.Time
	IncidentType constant.IncidentType
	Description  string
	PenaltyFee   int64
	Status       constant.VehicleIncidentStatus
}

type UpdateVehicleIncidentInput struct {
	VehicleID    *uint
	CustomerID   *uint
	RentalID     *uint
	IncidentDate *time.Time
	IncidentType *constant.IncidentType
	Description  *string
	PenaltyFee   *int64
	Status       *constant.VehicleIncidentStatus
}

type vehicleIncidentService struct {
	repo repository.VehicleIncidentRepository
}

func NewVehicleIncidentService(repo repository.VehicleIncidentRepository) VehicleIncidentService {
	return &vehicleIncidentService{repo: repo}
}

func (s *vehicleIncidentService) Create(ctx context.Context, data CreateVehicleIncidentInput) (*entity.VehicleIncident, error) {
	incident := entity.VehicleIncident{
		VehicleID:    data.VehicleID,
		CustomerID:   data.CustomerID,
		RentalID:     data.RentalID,
		IncidentDate: data.IncidentDate,
		IncidentType: data.IncidentType,
		Description:  data.Description,
		PenaltyFee:   data.PenaltyFee,
		Status:       data.Status,
	}
	if err := s.repo.Create(ctx, &incident); err != nil {
		return nil, err
	}
	return &incident, nil
}

func (s *vehicleIncidentService) GetByID(ctx context.Context, id uint) (*entity.VehicleIncident, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *vehicleIncidentService) GetByColumn(ctx context.Context, column string, value any) (entity.VehicleIncident, error) {
	return s.repo.GetByColumn(ctx, column, value)
}

func (s *vehicleIncidentService) List(ctx context.Context) ([]entity.VehicleIncident, error) {
	return s.repo.List(ctx)
}

func (s *vehicleIncidentService) Update(ctx context.Context, id uint, data UpdateVehicleIncidentInput) (*entity.VehicleIncident, error) {
	incident, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if data.VehicleID != nil {
		incident.VehicleID = *data.VehicleID
	}
	if data.CustomerID != nil {
		incident.CustomerID = *data.CustomerID
	}
	if data.RentalID != nil {
		incident.RentalID = *data.RentalID
	}
	if data.IncidentDate != nil {
		incident.IncidentDate = *data.IncidentDate
	}
	if data.IncidentType != nil {
		incident.IncidentType = *data.IncidentType
	}
	if data.Description != nil {
		incident.Description = *data.Description
	}
	if data.PenaltyFee != nil {
		incident.PenaltyFee = *data.PenaltyFee
	}
	if data.Status != nil {
		incident.Status = *data.Status
	}

	if err := s.repo.Update(ctx, incident); err != nil {
		return nil, err
	}
	return incident, nil
}

func (s *vehicleIncidentService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
