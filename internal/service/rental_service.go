package service

import (
	"context"
	"time"

	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"

	"gorm.io/gorm"
)

type RentalService interface {
	Create(ctx context.Context, data CreateRentalInput) (*entity.Rental, error)
	GetByID(ctx context.Context, id uint) (*entity.Rental, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.Rental, error)
	List(ctx context.Context) ([]entity.Rental, error)
	ListPaginated(ctx context.Context, page int, limit int) (*RentalListPaginatedResult, error)
	Update(ctx context.Context, id uint, data UpdateRentalInput) (*entity.Rental, error)
	Delete(ctx context.Context, id uint) error
}

type RentalListPaginatedResult struct {
	Items      []entity.Rental
	Page       int
	Limit      int
	Total      int64
	TotalPages int
}

type CreateRentalInput struct {
	CustomerID            uint
	VehicleID             uint
	StartDate             time.Time
	EndDate               time.Time
	TotalDay              int
	ReturnDate            *time.Time
	Price                 int64
	PenaltyFee            int64
	Subtotal              int64
	Notes                 string
	Status                entity.RentalStatus
	VehicleConditionStart string
	VehicleConditionEnd   string
	MileageStart          int
	MileageUsed           int
	MileageEnd            int
}

type UpdateRentalInput struct {
	CustomerID            *uint
	VehicleID             *uint
	StartDate             *time.Time
	EndDate               *time.Time
	TotalDay              *int
	ReturnDate            *time.Time
	Price                 *int64
	PenaltyFee            *int64
	Subtotal              *int64
	Notes                 *string
	Status                *entity.RentalStatus
	VehicleConditionStart *string
	VehicleConditionEnd   *string
	MileageStart          *int
	MileageUsed           *int
	MileageEnd            *int
}

type rentalService struct {
	db   *gorm.DB
	repo repository.RentalRepository
}

func NewRentalService(db *gorm.DB, repo repository.RentalRepository) RentalService {
	return &rentalService{db: db, repo: repo}
}

func (s *rentalService) Create(ctx context.Context, data CreateRentalInput) (*entity.Rental, error) {
	rental := entity.Rental{
		CustomerID:            data.CustomerID,
		VehicleID:             data.VehicleID,
		StartDate:             data.StartDate,
		EndDate:               data.EndDate,
		TotalDay:              data.TotalDay,
		ReturnDate:            data.ReturnDate,
		Price:                 data.Price,
		PenaltyFee:            data.PenaltyFee,
		Subtotal:              data.Subtotal,
		Notes:                 data.Notes,
		Status:                data.Status,
		VehicleConditionStart: data.VehicleConditionStart,
		VehicleConditionEnd:   data.VehicleConditionEnd,
		MileageStart:          data.MileageStart,
		MileageUsed:           data.MileageUsed,
		MileageEnd:            data.MileageEnd,
	}
	if err := s.repo.Create(ctx, &rental); err != nil {
		return nil, err
	}
	return &rental, nil
}

func (s *rentalService) GetByID(ctx context.Context, id uint) (*entity.Rental, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *rentalService) GetByColumn(ctx context.Context, column string, value any) (entity.Rental, error) {
	return s.repo.GetByColumn(ctx, column, value)
}

func (s *rentalService) List(ctx context.Context) ([]entity.Rental, error) {
	return s.repo.List(ctx)
}

func (s *rentalService) ListPaginated(ctx context.Context, page int, limit int) (*RentalListPaginatedResult, error) {
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

	return &RentalListPaginatedResult{
		Items:      items,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (s *rentalService) Update(ctx context.Context, id uint, data UpdateRentalInput) (*entity.Rental, error) {
	rental, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if data.CustomerID != nil {
		rental.CustomerID = *data.CustomerID
	}
	if data.VehicleID != nil {
		rental.VehicleID = *data.VehicleID
	}
	if data.StartDate != nil {
		rental.StartDate = *data.StartDate
	}
	if data.EndDate != nil {
		rental.EndDate = *data.EndDate
	}
	if data.TotalDay != nil {
		rental.TotalDay = *data.TotalDay
	}
	if data.ReturnDate != nil {
		rental.ReturnDate = data.ReturnDate
	}
	if data.Price != nil {
		rental.Price = *data.Price
	}
	if data.PenaltyFee != nil {
		rental.PenaltyFee = *data.PenaltyFee
	}
	if data.Subtotal != nil {
		rental.Subtotal = *data.Subtotal
	}
	if data.Notes != nil {
		rental.Notes = *data.Notes
	}
	if data.Status != nil {
		rental.Status = *data.Status
	}
	if data.VehicleConditionStart != nil {
		rental.VehicleConditionStart = *data.VehicleConditionStart
	}
	if data.VehicleConditionEnd != nil {
		rental.VehicleConditionEnd = *data.VehicleConditionEnd
	}
	if data.MileageStart != nil {
		rental.MileageStart = *data.MileageStart
	}
	if data.MileageUsed != nil {
		rental.MileageUsed = *data.MileageUsed
	}
	if data.MileageEnd != nil {
		rental.MileageEnd = *data.MileageEnd
	}

	if err := s.repo.Update(ctx, rental); err != nil {
		return nil, err
	}
	return rental, nil
}

func (s *rentalService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
