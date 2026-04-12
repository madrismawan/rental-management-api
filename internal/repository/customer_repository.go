package repository

import (
	"context"

	"gorm.io/gorm"

	"rental-management-api/internal/entity"
)

type CustomerRepository interface {
	Create(ctx context.Context, data *entity.Customer) error
	GetByID(ctx context.Context, id uint) (*entity.Customer, error)
	List(ctx context.Context) ([]entity.Customer, error)
	Update(ctx context.Context, data *entity.Customer) error
	Delete(ctx context.Context, id uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(ctx context.Context, data *entity.Customer) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *customerRepository) GetByID(ctx context.Context, id uint) (*entity.Customer, error) {
	var data entity.Customer
	if err := r.db.WithContext(ctx).Preload("User").First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *customerRepository) List(ctx context.Context) ([]entity.Customer, error) {
	var data []entity.Customer
	err := r.db.WithContext(ctx).Preload("User").Find(&data).Error
	return data, err
}

func (r *customerRepository) Update(ctx context.Context, data *entity.Customer) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *customerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Customer{}, id).Error
}
