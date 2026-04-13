package repository

import (
	"context"

	"gorm.io/gorm"

	"rental-management-api/internal/entity"
)

type VehicleRepository interface {
	Create(ctx context.Context, data *entity.Vehicle) error
	GetByID(ctx context.Context, id uint) (*entity.Vehicle, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.Vehicle, error)
	List(ctx context.Context) ([]entity.Vehicle, error)
	Update(ctx context.Context, data *entity.Vehicle) error
	Delete(ctx context.Context, id uint) error
}

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) Create(ctx context.Context, data *entity.Vehicle) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *vehicleRepository) GetByID(ctx context.Context, id uint) (*entity.Vehicle, error) {
	var data entity.Vehicle
	if err := r.db.WithContext(ctx).First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *vehicleRepository) GetByColumn(ctx context.Context, column string, value any) (entity.Vehicle, error) {
	var data entity.Vehicle
	if err := r.db.WithContext(ctx).Where(column+" = ?", value).First(&data).Error; err != nil {
		return entity.Vehicle{}, err
	}
	return data, nil
}

func (r *vehicleRepository) List(ctx context.Context) ([]entity.Vehicle, error) {
	var data []entity.Vehicle
	err := r.db.WithContext(ctx).Find(&data).Error
	return data, err
}

func (r *vehicleRepository) Update(ctx context.Context, data *entity.Vehicle) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *vehicleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Vehicle{}, id).Error
}
