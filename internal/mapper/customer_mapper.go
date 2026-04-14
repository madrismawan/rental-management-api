package mapper

import (
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

func ToCustomerResource(customer entity.Customer) dto.CustomerResource {
	return dto.CustomerResource{
		ID:          customer.ID,
		UserID:      customer.UserID,
		Name:        customer.User.Name,
		Email:       customer.User.Email,
		Password:    customer.User.Password,
		PhoneNumber: customer.PhoneNumber,
		Address:     customer.Address,
		AvatarURL:   customer.AvatarURL,
		CreatedAt:   customer.CreatedAt,
		UpdatedAt:   customer.UpdatedAt,
	}
}

func ToCustomersResource(customers []entity.Customer) []dto.CustomerResource {
	out := make([]dto.CustomerResource, 0, len(customers))
	for _, item := range customers {
		out = append(out, ToCustomerResource(item))
	}
	return out
}
