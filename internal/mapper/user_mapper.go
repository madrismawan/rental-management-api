package mapper

import (
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

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
