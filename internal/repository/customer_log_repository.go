package repository

import (
	"context"

	"gorm.io/gorm"

	"rental-management-api/internal/database"
	"rental-management-api/internal/entity"
)

type CustomerLogRepository interface {
	Create(ctx context.Context, data *entity.CustomerLog) error
	GetByColumn(ctx context.Context, column string, value any) (entity.CustomerLog, error)
	List(ctx context.Context) ([]entity.CustomerLog, error)
	ListPaginated(ctx context.Context, page int, limit int, customerID *uint) ([]entity.CustomerLog, int64, error)
}

type customerLogRepository struct {
	db *gorm.DB
}

func NewCustomerLogRepository(db *gorm.DB) CustomerLogRepository {
	return &customerLogRepository{db: db}
}

func (r *customerLogRepository) Create(ctx context.Context, data *entity.CustomerLog) error {
	return database.ExtractDB(ctx, r.db).Create(data).Error
}

func (r *customerLogRepository) GetByColumn(ctx context.Context, column string, value any) (entity.CustomerLog, error) {
	var data entity.CustomerLog
	if err := database.ExtractDB(ctx, r.db).Where(column+" = ?", value).First(&data).Error; err != nil {
		return entity.CustomerLog{}, err
	}
	return data, nil
}

func (r *customerLogRepository) List(ctx context.Context) ([]entity.CustomerLog, error) {
	var data []entity.CustomerLog
	err := database.ExtractDB(ctx, r.db).Order("id DESC").Find(&data).Error
	return data, err
}

func (r *customerLogRepository) ListPaginated(ctx context.Context, page int, limit int, customerID *uint) ([]entity.CustomerLog, int64, error) {
	var data []entity.CustomerLog
	var total int64

	db := database.ExtractDB(ctx, r.db).Model(&entity.CustomerLog{})
	if customerID != nil {
		db = db.Where("customer_id = ?", *customerID)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	query := database.ExtractDB(ctx, r.db)
	if customerID != nil {
		query = query.Where("customer_id = ?", *customerID)
	}
	err := query.
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
