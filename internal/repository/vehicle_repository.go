package repository

import (
	"context"

	"gorm.io/gorm"

	"rental-management-api/internal/database"
	"rental-management-api/internal/entity"
)

type VehicleRepository interface {
	Create(ctx context.Context, data *entity.Vehicle) error
	GetByID(ctx context.Context, id uint) (*entity.Vehicle, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.Vehicle, error)
	List(ctx context.Context) ([]entity.Vehicle, error)
	ListPaginated(ctx context.Context, page int, limit int) ([]entity.Vehicle, int64, error)
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
	return database.ExtractDB(ctx, r.db).Create(data).Error
}

func (r *vehicleRepository) GetByID(ctx context.Context, id uint) (*entity.Vehicle, error) {
	var data entity.Vehicle
	if err := database.ExtractDB(ctx, r.db).First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *vehicleRepository) GetByColumn(ctx context.Context, column string, value any) (entity.Vehicle, error) {
	var data entity.Vehicle
	if err := database.ExtractDB(ctx, r.db).Where(column+" = ?", value).First(&data).Error; err != nil {
		return entity.Vehicle{}, err
	}
	return data, nil
}

func (r *vehicleRepository) List(ctx context.Context) ([]entity.Vehicle, error) {
	var data []entity.Vehicle
	err := database.ExtractDB(ctx, r.db).Find(&data).Error
	return data, err
}

func (r *vehicleRepository) ListPaginated(ctx context.Context, page int, limit int) ([]entity.Vehicle, int64, error) {
	var data []entity.Vehicle
	var total int64

	db := database.ExtractDB(ctx, r.db).Model(&entity.Vehicle{})
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := database.ExtractDB(ctx, r.db).
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (r *vehicleRepository) Update(ctx context.Context, data *entity.Vehicle) error {
	return database.ExtractDB(ctx, r.db).Save(data).Error
}

func (r *vehicleRepository) Delete(ctx context.Context, id uint) error {
	return database.ExtractDB(ctx, r.db).Delete(&entity.Vehicle{}, id).Error
}
