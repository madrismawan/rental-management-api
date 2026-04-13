package mapper

import (
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

func ToUserEntity(req dto.CreateUserRequest) entity.User {
	return entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Password: req.Password,
	}
}

func ApplyUserUpdate(user *entity.User, req dto.UpdateUserRequest) {
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Role != nil {
		user.Role = *req.Role
	}
	if req.Password != nil {
		user.Password = *req.Password
	}
}

func ToUserResource(user entity.User) dto.UserResource {
	return dto.UserResource{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUsersResource(users []entity.User) []dto.UserResource {
	out := make([]dto.UserResource, 0, len(users))
	for _, item := range users {
		out = append(out, ToUserResource(item))
	}
	return out
}
