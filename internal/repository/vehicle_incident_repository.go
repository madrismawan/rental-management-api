package repository

import (
	"context"

	"gorm.io/gorm"

	"rental-management-api/internal/database"
	"rental-management-api/internal/entity"
)

type VehicleIncidentRepository interface {
	Create(ctx context.Context, data *entity.VehicleIncident) error
	GetByID(ctx context.Context, id uint) (*entity.VehicleIncident, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.VehicleIncident, error)
	List(ctx context.Context) ([]entity.VehicleIncident, error)
	ListPaginated(ctx context.Context, page int, limit int) ([]entity.VehicleIncident, int64, error)
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
	return database.ExtractDB(ctx, r.db).Create(data).Error
}

func (r *vehicleIncidentRepository) GetByID(ctx context.Context, id uint) (*entity.VehicleIncident, error) {
	var data entity.VehicleIncident
	err := database.ExtractDB(ctx, r.db).
		Preload("Vehicle").
		Preload("Customer.User").
		Preload("Rental").
		First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *vehicleIncidentRepository) GetByColumn(ctx context.Context, column string, value any) (entity.VehicleIncident, error) {
	var data entity.VehicleIncident
	err := database.ExtractDB(ctx, r.db).
		Preload("Vehicle").
		Preload("Customer.User").
		Preload("Rental").
		Where(column+" = ?", value).
		First(&data).Error
	if err != nil {
		return entity.VehicleIncident{}, err
	}
	return data, nil
}

func (r *vehicleIncidentRepository) List(ctx context.Context) ([]entity.VehicleIncident, error) {
	var data []entity.VehicleIncident
	err := database.ExtractDB(ctx, r.db).
		Preload("Vehicle").
		Preload("Customer.User").
		Preload("Rental").
		Find(&data).Error
	return data, err
}

func (r *vehicleIncidentRepository) ListPaginated(ctx context.Context, page int, limit int) ([]entity.VehicleIncident, int64, error) {
	var data []entity.VehicleIncident
	var total int64

	db := database.ExtractDB(ctx, r.db).Model(&entity.VehicleIncident{})
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := database.ExtractDB(ctx, r.db).
		Preload("Vehicle").
		Preload("Customer.User").
		Preload("Rental").
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (r *vehicleIncidentRepository) Update(ctx context.Context, data *entity.VehicleIncident) error {
	return database.ExtractDB(ctx, r.db).Save(data).Error
}

func (r *vehicleIncidentRepository) Delete(ctx context.Context, id uint) error {
	return database.ExtractDB(ctx, r.db).Delete(&entity.VehicleIncident{}, id).Error
}
