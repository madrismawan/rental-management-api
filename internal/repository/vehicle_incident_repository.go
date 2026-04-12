package repository

import (
	"context"

	"gorm.io/gorm"

	"rental-management-api/internal/entity"
)

type VehicleIncidentRepository interface {
	Create(ctx context.Context, data *entity.VehicleIncident) error
	GetByID(ctx context.Context, id uint) (*entity.VehicleIncident, error)
	List(ctx context.Context) ([]entity.VehicleIncident, error)
	Update(ctx context.Context, data *entity.VehicleIncident) error
	Delete(ctx context.Context, id uint) error
}

type vehicleIncidentRepository struct {
	db *gorm.DB
}

func NewVehicleIncidentRepository(db *gorm.DB) VehicleIncidentRepository {
	return &vehicleIncidentRepository{db: db}
}

func (r *vehicleIncidentRepository) Create(ctx context.Context, data *entity.VehicleIncident) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *vehicleIncidentRepository) GetByID(ctx context.Context, id uint) (*entity.VehicleIncident, error) {
	var data entity.VehicleIncident
	err := r.db.WithContext(ctx).
		Preload("Vehicle").
		Preload("Customer").
		Preload("Rental").
		First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *vehicleIncidentRepository) List(ctx context.Context) ([]entity.VehicleIncident, error) {
	var data []entity.VehicleIncident
	err := r.db.WithContext(ctx).
		Preload("Vehicle").
		Preload("Customer").
		Preload("Rental").
		Find(&data).Error
	return data, err
}

func (r *vehicleIncidentRepository) Update(ctx context.Context, data *entity.VehicleIncident) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *vehicleIncidentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.VehicleIncident{}, id).Error
}
