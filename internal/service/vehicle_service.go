package service

import (
	"context"

	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"
)

type VehicleService interface {
	Create(ctx context.Context, data entity.Vehicle) (*entity.Vehicle, error)
	GetByID(ctx context.Context, id uint) (*entity.Vehicle, error)
	List(ctx context.Context) ([]entity.Vehicle, error)
	Update(ctx context.Context, id uint, applyFn func(*entity.Vehicle)) (*entity.Vehicle, error)
	Delete(ctx context.Context, id uint) error
}

type vehicleService struct {
	repo repository.VehicleRepository
}

func NewVehicleService(repo repository.VehicleRepository) VehicleService {
	return &vehicleService{repo: repo}
}

func (s *vehicleService) Create(ctx context.Context, data entity.Vehicle) (*entity.Vehicle, error) {
	if err := s.repo.Create(ctx, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *vehicleService) GetByID(ctx context.Context, id uint) (*entity.Vehicle, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *vehicleService) List(ctx context.Context) ([]entity.Vehicle, error) {
	return s.repo.List(ctx)
}

func (s *vehicleService) Update(ctx context.Context, id uint, applyFn func(*entity.Vehicle)) (*entity.Vehicle, error) {
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

func (s *vehicleService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
