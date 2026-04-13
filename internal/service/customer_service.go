package service

import (
	"context"

	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"
)

type CustomerService interface {
	Create(ctx context.Context, data CreateCustomerInput) (*entity.Customer, error)
	GetByID(ctx context.Context, id uint) (*entity.Customer, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.Customer, error)
	List(ctx context.Context) ([]entity.Customer, error)
	Update(ctx context.Context, id uint, data UpdateCustomerInput) (*entity.Customer, error)
	Delete(ctx context.Context, id uint) error
}

type CreateCustomerInput struct {
	UserID      uint
	PhoneNumber string
	Address     string
	AvatarURL   string
}

type UpdateCustomerInput struct {
	UserID      *uint
	PhoneNumber *string
	Address     *string
	AvatarURL   *string
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) Create(ctx context.Context, data CreateCustomerInput) (*entity.Customer, error) {
	customer := entity.Customer{
		UserID:      data.UserID,
		PhoneNumber: data.PhoneNumber,
		Address:     data.Address,
		AvatarURL:   data.AvatarURL,
	}
	if err := s.repo.Create(ctx, &customer); err != nil {
		return nil, err
	}
	return &customer, nil
}

func (s *customerService) GetByID(ctx context.Context, id uint) (*entity.Customer, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *customerService) GetByColumn(ctx context.Context, column string, value any) (entity.Customer, error) {
	return s.repo.GetByColumn(ctx, column, value)
}

func (s *customerService) List(ctx context.Context) ([]entity.Customer, error) {
	return s.repo.List(ctx)
}

func (s *customerService) Update(ctx context.Context, id uint, data UpdateCustomerInput) (*entity.Customer, error) {
	customer, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if data.UserID != nil {
		customer.UserID = *data.UserID
	}
	if data.PhoneNumber != nil {
		customer.PhoneNumber = *data.PhoneNumber
	}
	if data.Address != nil {
		customer.Address = *data.Address
	}
	if data.AvatarURL != nil {
		customer.AvatarURL = *data.AvatarURL
	}

	if err := s.repo.Update(ctx, customer); err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *customerService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
