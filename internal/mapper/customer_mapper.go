package mapper

import (
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

func ToCustomerEntity(req dto.CreateCustomerRequest) entity.Customer {
	return entity.Customer{
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		AvatarURL:   req.AvatarURL,
	}
}

func ApplyCustomerUpdate(customer *entity.Customer, req dto.UpdateCustomerRequest) {
	if req.PhoneNumber != nil {
		customer.PhoneNumber = *req.PhoneNumber
	}
	if req.Address != nil {
		customer.Address = *req.Address
	}
	if req.AvatarURL != nil {
		customer.AvatarURL = *req.AvatarURL
	}
}

func ToCustomerResource(customer entity.Customer) dto.CustomerResource {
	return dto.CustomerResource{
		ID:          customer.ID,
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
