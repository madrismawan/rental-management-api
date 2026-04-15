package service

import (
	"context"

	"rental-management-api/internal/constant"
	"rental-management-api/internal/database"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"

	"gorm.io/gorm"
)

type CustomerLogService interface {
	Create(ctx context.Context, data CreateCustomerLogInput) (*entity.CustomerLog, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.CustomerLog, error)
	List(ctx context.Context) ([]entity.CustomerLog, error)
	ListPaginated(ctx context.Context, page int, limit int, customerID *uint) (*CustomerLogListPaginatedResult, error)
}

type CustomerLogListPaginatedResult struct {
	Items      []entity.CustomerLog
	Page       int
	Limit      int
	Total      int64
	TotalPages int
}

type CreateCustomerLogInput struct {
	CustomerID   uint
	CustomerName string
	Reason       string
	Status       constant.CustomerLogStatus
}

type UpdateCustomerLogInput struct {
	CustomerID   *uint
	CustomerName *string
	Reason       *string
	Status       *constant.CustomerLogStatus
}

type customerLogService struct {
	db           *gorm.DB
	repo         repository.CustomerLogRepository
	customerRepo repository.CustomerRepository
}

func NewCustomerLogService(db *gorm.DB, repo repository.CustomerLogRepository, customerRepo repository.CustomerRepository) CustomerLogService {
	return &customerLogService{db: db, repo: repo, customerRepo: customerRepo}
}

func (s *customerLogService) Create(ctx context.Context, data CreateCustomerLogInput) (*entity.CustomerLog, error) {
	item := entity.CustomerLog{}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctxTx := database.InjectTx(ctx, tx)

		customer, err := s.customerRepo.GetByID(ctxTx, data.CustomerID)
		if err != nil {
			return err
		}

		customerName := customer.User.Name
		if customerName == "" {
			customerName = data.CustomerName
		}

		customerStatus := constant.CustomerStatus(data.Status)
		customer.Status = customerStatus
		if err := s.customerRepo.Update(ctxTx, customer); err != nil {
			return err
		}

		item = entity.CustomerLog{
			CustomerID:   data.CustomerID,
			CustomerName: customerName,
			Reason:       data.Reason,
			Status:       data.Status,
		}
		if err := s.repo.Create(ctxTx, &item); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *customerLogService) GetByColumn(ctx context.Context, column string, value any) (entity.CustomerLog, error) {
	return s.repo.GetByColumn(ctx, column, value)
}

func (s *customerLogService) List(ctx context.Context) ([]entity.CustomerLog, error) {
	return s.repo.List(ctx)
}

func (s *customerLogService) ListPaginated(ctx context.Context, page int, limit int, customerID *uint) (*CustomerLogListPaginatedResult, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	items, total, err := s.repo.ListPaginated(ctx, page, limit, customerID)
	if err != nil {
		return nil, err
	}

	totalPages := 0
	if total > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	return &CustomerLogListPaginatedResult{
		Items:      items,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}
