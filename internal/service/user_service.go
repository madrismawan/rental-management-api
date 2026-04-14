package service

import (
	"context"
	"fmt"

	"rental-management-api/internal/constant"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/repository"
	"rental-management-api/pkg/errs"

	"gorm.io/gorm"
)

type UserService interface {
	Create(ctx context.Context, data CreateUserInput) (*entity.User, error)
	GetByID(ctx context.Context, id uint) (*entity.User, error)
	GetByColumn(ctx context.Context, column string, value any) (entity.User, error)
	List(ctx context.Context) ([]entity.User, error)
	ListPaginated(ctx context.Context, page int, limit int) (*UserListPaginatedResult, error)
	Update(ctx context.Context, id uint, data UpdateUserInput) (*entity.User, error)
	Delete(ctx context.Context, id uint) error
}

type UserListPaginatedResult struct {
	Items      []entity.User
	Page       int
	Limit      int
	Total      int64
	TotalPages int
}

type CreateUserInput struct {
	Name     string
	Email    string
	Role     constant.UserRole
	Password string
}

type UpdateUserInput struct {
	Name     *string
	Email    *string
	Role     *constant.UserRole
	Password *string
}

type userService struct {
	db   *gorm.DB
	repo repository.UserRepository
}

func NewUserService(db *gorm.DB, repo repository.UserRepository) UserService {
	return &userService{db: db, repo: repo}
}

func (s *userService) Create(ctx context.Context, data CreateUserInput) (*entity.User, error) {
	hashedPassword, err := HashPassword(data.Password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}
	existsUsers, err := s.GetByColumn(ctx, "email", data.Email)
	if err == nil && existsUsers.ID != 0 {
		return nil, errs.ErrEmailDuplicate
	}

	user := entity.User{
		Name:     data.Name,
		Email:    data.Email,
		Role:     data.Role,
		Password: hashedPassword,
	}
	if err := s.repo.Create(ctx, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	return s.repo.GetByID(ctx, id)
}
func (s *userService) GetByColumn(ctx context.Context, column string, value any) (entity.User, error) {
	return s.repo.GetByColumn(ctx, column, value)
}

func (s *userService) List(ctx context.Context) ([]entity.User, error) {
	return s.repo.List(ctx)
}

func (s *userService) ListPaginated(ctx context.Context, page int, limit int) (*UserListPaginatedResult, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	items, total, err := s.repo.ListPaginated(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	totalPages := 0
	if total > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	return &UserListPaginatedResult{
		Items:      items,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (s *userService) Update(ctx context.Context, id uint, data UpdateUserInput) (*entity.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if data.Name != nil {
		user.Name = *data.Name
	}
	if data.Email != nil {
		user.Email = *data.Email
	}
	if data.Role != nil {
		user.Role = *data.Role
	}
	if data.Password != nil && *data.Password != "" {
		hashedPassword, err := HashPassword(*data.Password)
		if err != nil {
			return nil, fmt.Errorf("hash password: %w", err)
		}
		user.Password = hashedPassword
	}
	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
