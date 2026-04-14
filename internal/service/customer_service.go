package service

import (
	"context"

	"rental-management-api/internal/database"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"

	"gorm.io/gorm"
)

type CustomerService interface {
	Create(ctx context.Context, data CreateCustomerInput) (*entity.Customer, error)
	GetByID(ctx context.Context, id uint) (*entity.Customer, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.Customer, error)
	List(ctx context.Context) ([]entity.Customer, error)
	Update(ctx context.Context, id uint, data UpdateCustomerInput) (*entity.Customer, error)
	Delete(ctx context.Context, id uint) error
}

type CreateCustomerWithUserInput struct {
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	AvatarURL   string
}

type CreateCustomerInput struct {
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
	db          *gorm.DB
	repo        repository.CustomerRepository
	userService UserService
}

func NewCustomerService(db *gorm.DB, userService UserService, repo repository.CustomerRepository) CustomerService {
	return &customerService{
		db:          db,
		repo:        repo,
		userService: userService,
	}
}

func (s *customerService) CreateWithUser(ctx context.Context, data CreateCustomerWithUserInput) (*entity.Customer, error) {
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctxTx := database.InjectTx(ctx, tx)
		_, err := s.userService.Create(ctxTx, CreateUserInput{
			Name:     data.Name,
			Email:    data.Email,
			Password: data.Password,
		})
		if err != nil {
			return err
		}
		_, err = s.Create(ctxTx, CreateCustomerInput{
			PhoneNumber: data.PhoneNumber,
			Address:     data.Address,
			AvatarURL:   data.AvatarURL,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *customerService) Create(ctx context.Context, data CreateCustomerInput) (*entity.Customer, error) {
	customer := entity.Customer{
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
