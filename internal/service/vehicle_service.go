package service

import (
	"context"

	"rental-management-api/internal/constant"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"

	"gorm.io/gorm"
)

type VehicleService interface {
	Create(ctx context.Context, data CreateVehicleInput) (*entity.Vehicle, error)
	GetByID(ctx context.Context, id uint) (*entity.Vehicle, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.Vehicle, error)
	List(ctx context.Context) ([]entity.Vehicle, error)
	Update(ctx context.Context, id uint, data UpdateVehicleInput) (*entity.Vehicle, error)
	Delete(ctx context.Context, id uint) error
}

type CreateVehicleInput struct {
	PlateNumber string
	Color       string
	Brand       string
	Model       string
	CC          int
	Year        int
	Mileage     int
	DailyRate   int64
	Status      constant.VehicleStatus
	Notes       string
}

type UpdateVehicleInput struct {
	PlateNumber *string
	Color       *string
	Brand       *string
	Model       *string
	CC          *int
	Year        *int
	Mileage     *int
	DailyRate   *int64
	Status      *constant.VehicleStatus
	Notes       *string
}

type vehicleService struct {
	db   *gorm.DB
	repo repository.VehicleRepository
}

func NewVehicleService(db *gorm.DB, repo repository.VehicleRepository) VehicleService {
	return &vehicleService{db: db, repo: repo}
}

func (s *vehicleService) Create(ctx context.Context, data CreateVehicleInput) (*entity.Vehicle, error) {
	vehicle := entity.Vehicle{
		PlateNumber: data.PlateNumber,
		Color:       data.Color,
		Brand:       data.Brand,
		Model:       data.Model,
		CC:          data.CC,
		Year:        data.Year,
		Mileage:     data.Mileage,
		DailyRate:   data.DailyRate,
		Status:      data.Status,
		Notes:       data.Notes,
	}
	if err := s.repo.Create(ctx, &vehicle); err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (s *vehicleService) GetByID(ctx context.Context, id uint) (*entity.Vehicle, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *vehicleService) GetByColumn(ctx context.Context, column string, value any) (entity.Vehicle, error) {
	return s.repo.GetByColumn(ctx, column, value)
}

func (s *vehicleService) List(ctx context.Context) ([]entity.Vehicle, error) {
	return s.repo.List(ctx)
}

func (s *vehicleService) Update(ctx context.Context, id uint, data UpdateVehicleInput) (*entity.Vehicle, error) {
	vehicle, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if data.PlateNumber != nil {
		vehicle.PlateNumber = *data.PlateNumber
	}
	if data.Color != nil {
		vehicle.Color = *data.Color
	}
	if data.Brand != nil {
		vehicle.Brand = *data.Brand
	}
	if data.Model != nil {
		vehicle.Model = *data.Model
	}
	if data.CC != nil {
		vehicle.CC = *data.CC
	}
	if data.Year != nil {
		vehicle.Year = *data.Year
	}
	if data.Mileage != nil {
		vehicle.Mileage = *data.Mileage
	}
	if data.DailyRate != nil {
		vehicle.DailyRate = *data.DailyRate
	}
	if data.Status != nil {
		vehicle.Status = *data.Status
	}
	if data.Notes != nil {
		vehicle.Notes = *data.Notes
	}

	if err := s.repo.Update(ctx, vehicle); err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (s *vehicleService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
