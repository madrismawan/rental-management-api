package service

import (
	"context"
	"time"

	"rental-management-api/internal/constant"
	"rental-management-api/internal/database"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"

	"gorm.io/gorm"
)

type VehicleIncidentService interface {
	Create(ctx context.Context, data CreateVehicleIncidentInput) (*entity.VehicleIncident, error)
	GetByID(ctx context.Context, id uint) (*entity.VehicleIncident, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.VehicleIncident, error)
	List(ctx context.Context) ([]entity.VehicleIncident, error)
	ListPaginated(ctx context.Context, page int, limit int) (*VehicleIncidentListPaginatedResult, error)
	Update(ctx context.Context, id uint, data UpdateVehicleIncidentInput) (*entity.VehicleIncident, error)
	Progress(ctx context.Context, id uint) (*entity.VehicleIncident, error)
	Closed(ctx context.Context, id uint) (*entity.VehicleIncident, error)
	Resolved(ctx context.Context, id uint) (*entity.VehicleIncident, error)
	Delete(ctx context.Context, id uint) error
}

type VehicleIncidentListPaginatedResult struct {
	Items      []entity.VehicleIncident
	Page       int
	Limit      int
	Total      int64
	TotalPages int
}

type CreateVehicleIncidentInput struct {
	VehicleID    uint
	CustomerID   *uint
	RentalID     *uint
	IncidentDate time.Time
	IncidentType constant.IncidentType
	Description  string
	Cost         int64
	Status       constant.VehicleIncidentStatus
}

type UpdateVehicleIncidentInput struct {
	VehicleID    *uint
	CustomerID   *uint
	RentalID     *uint
	IncidentDate *time.Time
	IncidentType *constant.IncidentType
	Description  *string
	Cost         *int64
	Status       *constant.VehicleIncidentStatus
}

type vehicleIncidentService struct {
	db             *gorm.DB
	vehicleService VehicleService
	rentalRepo     repository.RentalRepository
	repo           repository.VehicleIncidentRepository
}

func NewVehicleIncidentService(db *gorm.DB, repo repository.VehicleIncidentRepository, vehicleService VehicleService, rentalRepo repository.RentalRepository) VehicleIncidentService {
	return &vehicleIncidentService{db: db, repo: repo, vehicleService: vehicleService, rentalRepo: rentalRepo}
}

func (s *vehicleIncidentService) Create(ctx context.Context, data CreateVehicleIncidentInput) (*entity.VehicleIncident, error) {
	customerID := data.CustomerID
	if data.RentalID != nil {
		rental, err := s.rentalRepo.GetByID(ctx, *data.RentalID)
		if err != nil {
			return nil, err
		}
		customerID = &rental.CustomerID
	}

	incident := entity.VehicleIncident{
		VehicleID:    data.VehicleID,
		CustomerID:   customerID,
		RentalID:     data.RentalID,
		IncidentDate: data.IncidentDate,
		IncidentType: data.IncidentType,
		Description:  data.Description,
		Cost:         data.Cost,
		Status:       constant.VehicleIncidentStatusOpen,
	}
	if err := s.repo.Create(ctx, &incident); err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, incident.ID)
}

func (s *vehicleIncidentService) GetByID(ctx context.Context, id uint) (*entity.VehicleIncident, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *vehicleIncidentService) GetByColumn(ctx context.Context, column string, value any) (entity.VehicleIncident, error) {
	return s.repo.GetByColumn(ctx, column, value)
}

func (s *vehicleIncidentService) List(ctx context.Context) ([]entity.VehicleIncident, error) {
	return s.repo.List(ctx)
}

func (s *vehicleIncidentService) ListPaginated(ctx context.Context, page int, limit int) (*VehicleIncidentListPaginatedResult, error) {
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

	return &VehicleIncidentListPaginatedResult{
		Items:      items,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (s *vehicleIncidentService) Update(ctx context.Context, id uint, data UpdateVehicleIncidentInput) (*entity.VehicleIncident, error) {
	incident, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if data.VehicleID != nil {
		incident.VehicleID = *data.VehicleID
	}
	if data.CustomerID != nil {
		incident.CustomerID = data.CustomerID
	}
	if data.RentalID != nil {
		incident.RentalID = data.RentalID
	}
	if data.IncidentDate != nil {
		incident.IncidentDate = *data.IncidentDate
	}
	if data.IncidentType != nil {
		incident.IncidentType = *data.IncidentType
	}
	if data.Description != nil {
		incident.Description = *data.Description
	}
	if data.Cost != nil {
		incident.Cost = *data.Cost
	}
	if data.Status != nil {
		incident.Status = *data.Status
	}

	if err := s.repo.Update(ctx, incident); err != nil {
		return nil, err
	}
	return incident, nil
}

func (s *vehicleIncidentService) Progress(ctx context.Context, id uint) (*entity.VehicleIncident, error) {
	incident, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctxTx := database.InjectTx(ctx, tx)

		statusIncident := constant.VehicleIncidentStatusInProgress
		if _, err := s.Update(ctxTx, id, UpdateVehicleIncidentInput{
			Status: &statusIncident,
		}); err != nil {
			return err
		}

		statusVehicle := constant.VehicleStatusMaintenance
		if _, err := s.vehicleService.Update(ctxTx, incident.VehicleID, UpdateVehicleInput{
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

func (s *vehicleIncidentService) Closed(ctx context.Context, id uint) (*entity.VehicleIncident, error) {
	incident, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	statusIncident := constant.VehicleIncidentStatusClosed
	incident.Status = statusIncident
	if err := s.repo.Update(ctx, incident); err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, id)
}

func (s *vehicleIncidentService) Resolved(ctx context.Context, id uint) (*entity.VehicleIncident, error) {
	incident, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctxTx := database.InjectTx(ctx, tx)

		statusIncident := constant.VehicleIncidentStatusResolved
		if _, err := s.Update(ctxTx, id, UpdateVehicleIncidentInput{
			Status: &statusIncident,
		}); err != nil {
			return err
		}

		statusVehicle := constant.VehicleStatusAvailable
		if _, err := s.vehicleService.Update(ctxTx, incident.VehicleID, UpdateVehicleInput{
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

func (s *vehicleIncidentService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
