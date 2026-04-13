package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"rental-management-api/internal/entity"
	"rental-management-api/pkg/errs"
)

type UserRepository interface {
	Create(ctx context.Context, data *entity.User) error
	GetByID(ctx context.Context, id uint) (*entity.User, error)
	GetByColumn(c context.Context, column string, value any) (entity.User, error)
	List(ctx context.Context) ([]entity.User, error)
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
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	var data entity.User
	if err := r.db.WithContext(ctx).First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *userRepository) GetByColumn(ctx context.Context, column string, value any) (entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where(column+" = ?", value).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, errs.ErrUserNotFound
		}
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) List(ctx context.Context) ([]entity.User, error) {
	var data []entity.User
	err := r.db.WithContext(ctx).Find(&data).Error
	return data, err
}

func (r *userRepository) Update(ctx context.Context, data *entity.User) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.User{}, id).Error
}
