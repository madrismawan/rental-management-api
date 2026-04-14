package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"rental-management-api/internal/database"
	"rental-management-api/internal/entity"
	"rental-management-api/pkg/errs"
)

type UserRepository interface {
	Create(ctx context.Context, data *entity.User) error
	GetByID(ctx context.Context, id uint) (*entity.User, error)
	GetByColumn(c context.Context, column string, value any) (entity.User, error)
	List(ctx context.Context) ([]entity.User, error)
	ListPaginated(ctx context.Context, page int, limit int) ([]entity.User, int64, error)
	Update(ctx context.Context, data *entity.User) error
	Delete(ctx context.Context, id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, data *entity.User) error {
	return database.ExtractDB(ctx, r.db).Create(data).Error
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	var data entity.User
	if err := database.ExtractDB(ctx, r.db).First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *userRepository) GetByColumn(ctx context.Context, column string, value any) (entity.User, error) {
	var user entity.User
	if err := database.ExtractDB(ctx, r.db).Where(column+" = ?", value).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, errs.ErrUserNotFound
		}
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) List(ctx context.Context) ([]entity.User, error) {
	var data []entity.User
	err := database.ExtractDB(ctx, r.db).Find(&data).Error
	return data, err
}

func (r *userRepository) ListPaginated(ctx context.Context, page int, limit int) ([]entity.User, int64, error) {
	var data []entity.User
	var total int64

	db := database.ExtractDB(ctx, r.db).Model(&entity.User{})
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

func (r *userRepository) Update(ctx context.Context, data *entity.User) error {
	return database.ExtractDB(ctx, r.db).Save(data).Error
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return database.ExtractDB(ctx, r.db).Delete(&entity.User{}, id).Error
}
