package service

import (
	"context"

	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"
)

type CustomerService interface {
	Create(ctx context.Context, data entity.Customer) (*entity.Customer, error)
	GetByID(ctx context.Context, id uint) (*entity.Customer, error)
	List(ctx context.Context) ([]entity.Customer, error)
	Update(ctx context.Context, id uint, applyFn func(*entity.Customer)) (*entity.Customer, error)
	Delete(ctx context.Context, id uint) error
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) Create(ctx context.Context, data entity.Customer) (*entity.Customer, error) {
	if err := s.repo.Create(ctx, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *customerService) GetByID(ctx context.Context, id uint) (*entity.Customer, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *customerService) List(ctx context.Context) ([]entity.Customer, error) {
	return s.repo.List(ctx)
}

func (s *customerService) Update(ctx context.Context, id uint, applyFn func(*entity.Customer)) (*entity.Customer, error) {
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

func (s *customerService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
