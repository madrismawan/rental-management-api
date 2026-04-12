package service

import (
	"context"

	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"
)

type VehicleIncidentService interface {
	Create(ctx context.Context, data entity.VehicleIncident) (*entity.VehicleIncident, error)
	GetByID(ctx context.Context, id uint) (*entity.VehicleIncident, error)
	List(ctx context.Context) ([]entity.VehicleIncident, error)
	Update(ctx context.Context, id uint, applyFn func(*entity.VehicleIncident)) (*entity.VehicleIncident, error)
	Delete(ctx context.Context, id uint) error
}

type vehicleIncidentService struct {
	repo repository.VehicleIncidentRepository
}

func NewVehicleIncidentService(repo repository.VehicleIncidentRepository) VehicleIncidentService {
	return &vehicleIncidentService{repo: repo}
}

func (s *vehicleIncidentService) Create(ctx context.Context, data entity.VehicleIncident) (*entity.VehicleIncident, error) {
	if err := s.repo.Create(ctx, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *vehicleIncidentService) GetByID(ctx context.Context, id uint) (*entity.VehicleIncident, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *vehicleIncidentService) List(ctx context.Context) ([]entity.VehicleIncident, error) {
	return s.repo.List(ctx)
}

func (s *vehicleIncidentService) Update(ctx context.Context, id uint, applyFn func(*entity.VehicleIncident)) (*entity.VehicleIncident, error) {
	data, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	applyFn(data)
	if err := s.repo.Update(ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *vehicleIncidentService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
