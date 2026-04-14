package repository

import (
	"context"

	"gorm.io/gorm"

	"rental-management-api/internal/database"
	"rental-management-api/internal/entity"
)

type CustomerRepository interface {
	Create(ctx context.Context, data *entity.Customer) error
	GetByID(ctx context.Context, id uint) (*entity.Customer, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.Customer, error)
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
	return database.ExtractDB(ctx, r.db).Create(data).Error
}

func (r *customerRepository) GetByID(ctx context.Context, id uint) (*entity.Customer, error) {
	var data entity.Customer
	if err := database.ExtractDB(ctx, r.db).Preload("User").First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *customerRepository) GetByColumn(ctx context.Context, column string, value any) (entity.Customer, error) {
	var data entity.Customer
	if err := database.ExtractDB(ctx, r.db).Preload("User").Where(column+" = ?", value).First(&data).Error; err != nil {
		return entity.Customer{}, err
	}
	return data, nil
}

func (r *customerRepository) List(ctx context.Context) ([]entity.Customer, error) {
	var data []entity.Customer
	err := database.ExtractDB(ctx, r.db).Preload("User").Find(&data).Error
	return data, err
}

func (r *customerRepository) Update(ctx context.Context, data *entity.Customer) error {
	return database.ExtractDB(ctx, r.db).Save(data).Error
}

func (r *customerRepository) Delete(ctx context.Context, id uint) error {
	return database.ExtractDB(ctx, r.db).Delete(&entity.Customer{}, id).Error
}
