package service

import (
	"context"

	"rental-management-api/internal/constant"
	"rental-management-api/internal/database"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"

	"gorm.io/gorm"
)

type CustomerService interface {
	Create(ctx context.Context, data CreateCustomerInput) (*entity.Customer, error)
	CreateWithUser(ctx context.Context, data CreateCustomerWithUserInput) (*entity.Customer, error)
	GetByID(ctx context.Context, id uint) (*entity.Customer, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.Customer, error)
	GetOptions(ctx context.Context) ([]entity.Customer, error)
	List(ctx context.Context) ([]entity.Customer, error)
	ListPaginated(ctx context.Context, page int, limit int) (*CustomerListPaginatedResult, error)
	Update(ctx context.Context, id uint, data UpdateCustomerInput) (*entity.Customer, error)
	Delete(ctx context.Context, id uint) error
}

type CustomerListPaginatedResult struct {
	Items      []entity.Customer
	Page       int
	Limit      int
	Total      int64
	TotalPages int
}

type CreateCustomerWithUserInput struct {
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	Status      constant.CustomerStatus
	AvatarURL   string
}

type CreateCustomerInput struct {
	UserID      uint
	PhoneNumber string
	Address     string
	Status      constant.CustomerStatus
	AvatarURL   string
}

type UpdateCustomerInput struct {
	UserID      *uint
	Name        *string
	Email       *string
	Password    *string
	PhoneNumber *string
	Address     *string
	Status      *constant.CustomerStatus
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
	var customer *entity.Customer
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctxTx := database.InjectTx(ctx, tx)
		user, err := s.userService.Create(ctxTx, CreateUserInput{
			Name:     data.Name,
			Email:    data.Email,
			Password: data.Password,
			Role:     constant.UserRoleCustomer,
		})
		if err != nil {
			return err
		}
		customer, err = s.Create(ctxTx, CreateCustomerInput{
			UserID:      user.ID,
			PhoneNumber: data.PhoneNumber,
			Address:     data.Address,
			Status:      data.Status,
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
	return customer, nil
}

func (s *customerService) Create(ctx context.Context, data CreateCustomerInput) (*entity.Customer, error) {
	status := data.Status
	if status == "" {
		status = constant.CustomerStatusActive
	}

	customer := entity.Customer{
		UserID:      data.UserID,
		PhoneNumber: data.PhoneNumber,
		Address:     data.Address,
		Status:      status,
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

func (s *customerService) GetOptions(ctx context.Context) ([]entity.Customer, error) {
	return s.repo.GetOptions(ctx)
}

func (s *customerService) List(ctx context.Context) ([]entity.Customer, error) {
	return s.repo.List(ctx)
}

func (s *customerService) ListPaginated(ctx context.Context, page int, limit int) (*CustomerListPaginatedResult, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	items, total, err := s.repo.ListPaginated(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	totalPages := 0
	if total > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	return &CustomerListPaginatedResult{
		Items:      items,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (s *customerService) Update(ctx context.Context, id uint, data UpdateCustomerInput) (*entity.Customer, error) {
	var customer *entity.Customer
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctxTx := database.InjectTx(ctx, tx)

		var err error
		customer, err = s.repo.GetByID(ctxTx, id)
		if err != nil {
			return err
		}

		if data.UserID != nil {
			customer.UserID = *data.UserID
		}

		if data.Name != nil || data.Email != nil || data.Password != nil {
			_, err = s.userService.Update(ctxTx, customer.UserID, UpdateUserInput{
				Name:     data.Name,
				Email:    data.Email,
				Password: data.Password,
			})
			if err != nil {
				return err
			}
		}

		if data.PhoneNumber != nil {
			customer.PhoneNumber = *data.PhoneNumber
		}
		if data.Address != nil {
			customer.Address = *data.Address
		}
		if data.Status != nil {
			customer.Status = *data.Status
		}
		if data.AvatarURL != nil {
			customer.AvatarURL = *data.AvatarURL
		}

		if err := s.repo.Update(ctxTx, customer); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *customerService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
