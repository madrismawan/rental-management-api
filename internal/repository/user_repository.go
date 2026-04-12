package repository

import (
	"context"

	"gorm.io/gorm"

	"rental-management-api/internal/entity"
)

type UserRepository interface {
	Create(ctx context.Context, data *entity.User) error
	GetByID(ctx context.Context, id uint) (*entity.User, error)
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
