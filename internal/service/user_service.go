package service

import (
	"context"

	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, data entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id uint) (*entity.User, error)
	List(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, id uint, applyFn func(*entity.User)) (*entity.User, error)
	Delete(ctx context.Context, id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(ctx context.Context, data entity.User) (*entity.User, error) {
	if err := s.repo.Create(ctx, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *userService) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) List(ctx context.Context) ([]entity.User, error) {
	return s.repo.List(ctx)
}

func (s *userService) Update(ctx context.Context, id uint, applyFn func(*entity.User)) (*entity.User, error) {
	data, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	applyFn(data)
	if err := s.repo.Update(ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *userService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
