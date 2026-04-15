package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"rental-management-api/internal/constant"
	"rental-management-api/internal/database"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"

	"gorm.io/gorm"
)

type RentalService interface {
	Create(ctx context.Context, data CreateRentalInput) (*entity.Rental, error)
	GetByID(ctx context.Context, id uint) (*entity.Rental, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.Rental, error)
	GetOptions(ctx context.Context) ([]entity.Rental, error)
	List(ctx context.Context) ([]entity.Rental, error)
	ListPaginated(ctx context.Context, page int, limit int) (*RentalListPaginatedResult, error)
	Update(ctx context.Context, id uint, data UpdateRentalInput) (*entity.Rental, error)
	Active(ctx context.Context, id uint) (*entity.Rental, error)
	Cancel(ctx context.Context, id uint) (*entity.Rental, error)
	Complete(ctx context.Context, id uint, data CompleteRentalInput) (*entity.Rental, error)
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
	Notes                 string
	VehicleConditionStart constant.VehicleCondition
	MileageStart          int
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
	Status                *constant.RentalStatus
	VehicleConditionStart *constant.VehicleCondition
	VehicleConditionEnd   *constant.VehicleCondition
	MileageStart          *int
	MileageUsed           *int
	MileageEnd            *int
}

type CompleteRentalInput struct {
	ReturnDate          time.Time
	PenaltyFee          int64
	IncidentType        string
	Description         string
	VehicleConditionEnd constant.VehicleCondition
	MileageEnd          int
}

type rentalService struct {
	db                     *gorm.DB
	vehicleService         VehicleService
	vehicleIncidentService VehicleIncidentService
	repo                   repository.RentalRepository
}

func NewRentalService(db *gorm.DB, repo repository.RentalRepository, vehicleService VehicleService, vehicleIncidentService VehicleIncidentService) RentalService {
	return &rentalService{db: db, repo: repo, vehicleService: vehicleService, vehicleIncidentService: vehicleIncidentService}
}

func (s *rentalService) Create(ctx context.Context, data CreateRentalInput) (*entity.Rental, error) {
	if data.EndDate.Before(data.StartDate) {
		return nil, fmt.Errorf("end_date must be greater than or equal to start_date")
	}

	vehicle, err := s.vehicleService.GetByID(ctx, data.VehicleID)
	if err != nil {
		return nil, err
	}

	totalDay := int(data.EndDate.Sub(data.StartDate).Hours() / 24)
	if totalDay <= 0 {
		totalDay = 1
	}

	price := vehicle.DailyRate
	subtotal := price * int64(totalDay)

	var rental entity.Rental
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctxTx := database.InjectTx(ctx, tx)
		noInvoice, err := GenerateInvoiceNumber()
		if err != nil {
			return err
		}

		rental = entity.Rental{
			NoInvoice:             noInvoice,
			CustomerID:            data.CustomerID,
			VehicleID:             data.VehicleID,
			StartDate:             data.StartDate,
			EndDate:               data.EndDate,
			TotalDay:              totalDay,
			Price:                 price,
			PenaltyFee:            0,
			Subtotal:              subtotal,
			Notes:                 data.Notes,
			Status:                constant.RentalStatusActive,
			VehicleConditionStart: data.VehicleConditionStart,
			MileageStart:          data.MileageStart,
			MileageUsed:           0,
			MileageEnd:            data.MileageStart,
		}

		err = s.repo.Create(ctxTx, &rental)
		if err != nil {
			return err
		}
		vehicle.Status = constant.VehicleStatusRented
		if _, err := s.vehicleService.Update(ctxTx, vehicle.ID, UpdateVehicleInput{
			Status: &vehicle.Status,
		}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &rental, nil
}

func GenerateInvoiceNumber() (string, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	suffix := rand.Intn(9000) + 1000

	return fmt.Sprintf("INV%s-%d", time.Now().Format("20060102"), suffix), nil
}

func (s *rentalService) GetByID(ctx context.Context, id uint) (*entity.Rental, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *rentalService) GetByColumn(ctx context.Context, column string, value any) (entity.Rental, error) {
	return s.repo.GetByColumn(ctx, column, value)
}

func (s *rentalService) GetOptions(ctx context.Context) ([]entity.Rental, error) {
	return s.repo.GetOptions(ctx)
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

func (s *rentalService) Active(ctx context.Context, id uint) (*entity.Rental, error) {
	rental, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctxTx := database.InjectTx(ctx, tx)

		statusRental := constant.RentalStatusActive
		rental.Status = statusRental
		if err := s.repo.Update(ctxTx, rental); err != nil {
			return err
		}

		statusVehicle := constant.VehicleStatusRented
		if _, err := s.vehicleService.Update(ctxTx, rental.VehicleID, UpdateVehicleInput{
			Status: &statusVehicle,
		}); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, id)
}

func (s *rentalService) Cancel(ctx context.Context, id uint) (*entity.Rental, error) {
	rental, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	rental.Status = constant.RentalStatusCancelled
	if err := s.repo.Update(ctx, rental); err != nil {
		return nil, err
	}

	return rental, nil
}

func (s *rentalService) Complete(ctx context.Context, id uint, data CompleteRentalInput) (*entity.Rental, error) {
	rental, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctxTx := database.InjectTx(ctx, tx)
		penaltyFee := rental.PenaltyFee + data.PenaltyFee
		subtotal := rental.Subtotal + data.PenaltyFee
		status := constant.RentalStatusCompleted
		mileageUsed := data.MileageEnd - rental.MileageStart
		rental, err = s.Update(ctxTx, id, UpdateRentalInput{
			ReturnDate:          &data.ReturnDate,
			PenaltyFee:          &penaltyFee,
			Subtotal:            &subtotal,
			Status:              &status,
			VehicleConditionEnd: &data.VehicleConditionEnd,
			MileageUsed:         &mileageUsed,
			MileageEnd:          &data.MileageEnd,
		})

		statusVehicle := constant.VehicleStatusAvailable
		if data.IncidentType != "" {
			_, err := s.vehicleIncidentService.Create(ctxTx, CreateVehicleIncidentInput{
				VehicleID:    rental.VehicleID,
				CustomerID:   &rental.CustomerID,
				RentalID:     &rental.ID,
				IncidentDate: data.ReturnDate,
				IncidentType: constant.IncidentType(data.IncidentType),
				Description:  data.Description,
				Cost:         data.PenaltyFee,
				Status:       constant.VehicleIncidentStatusOpen,
			})
			if err != nil {
				return err
			}
			statusVehicle = constant.VehicleStatusUnavailable
		}

		s.vehicleService.Update(ctxTx, rental.VehicleID, UpdateVehicleInput{
			Mileage:   &data.MileageEnd,
			Condition: &data.VehicleConditionEnd,
			Status:    &statusVehicle,
		})

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return rental, nil
}

func (s *rentalService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
