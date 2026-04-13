package repository

import (
	"context"

	"gorm.io/gorm"

	"rental-management-api/internal/entity"
)

type RentalRepository interface {
	Create(ctx context.Context, data *entity.Rental) error
	GetByID(ctx context.Context, id uint) (*entity.Rental, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.Rental, error)
	List(ctx context.Context) ([]entity.Rental, error)
	Update(ctx context.Context, data *entity.Rental) error
	Delete(ctx context.Context, id uint) error
}

type rentalRepository struct {
	db *gorm.DB
}

func NewRentalRepository(db *gorm.DB) RentalRepository {
	return &rentalRepository{db: db}
}

func (r *rentalRepository) Create(ctx context.Context, data *entity.Rental) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *rentalRepository) GetByID(ctx context.Context, id uint) (*entity.Rental, error) {
	var data entity.Rental
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("VehicleIncidents").
		First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *rentalRepository) GetByColumn(ctx context.Context, column string, value any) (entity.Rental, error) {
	var data entity.Rental
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("VehicleIncidents").
		Where(column+" = ?", value).
		First(&data).Error
	if err != nil {
		return entity.Rental{}, err
	}
	return data, nil
}

func (r *rentalRepository) List(ctx context.Context) ([]entity.Rental, error) {
	var data []entity.Rental
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("VehicleIncidents").
		Find(&data).Error
	return data, err
}

func (r *rentalRepository) Update(ctx context.Context, data *entity.Rental) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *rentalRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Rental{}, id).Error
}
