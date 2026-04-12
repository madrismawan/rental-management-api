package service

import (
	"context"

	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"
)

type RentalService interface {
	Create(ctx context.Context, data entity.Rental) (*entity.Rental, error)
	GetByID(ctx context.Context, id uint) (*entity.Rental, error)
	List(ctx context.Context) ([]entity.Rental, error)
	Update(ctx context.Context, id uint, applyFn func(*entity.Rental)) (*entity.Rental, error)
	Delete(ctx context.Context, id uint) error
}

type rentalService struct {
	repo repository.RentalRepository
}

func NewRentalService(repo repository.RentalRepository) RentalService {
	return &rentalService{repo: repo}
}

func (s *rentalService) Create(ctx context.Context, data entity.Rental) (*entity.Rental, error) {
	if err := s.repo.Create(ctx, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *rentalService) GetByID(ctx context.Context, id uint) (*entity.Rental, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *rentalService) List(ctx context.Context) ([]entity.Rental, error) {
	return s.repo.List(ctx)
}

func (s *rentalService) Update(ctx context.Context, id uint, applyFn func(*entity.Rental)) (*entity.Rental, error) {
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

func (s *rentalService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
