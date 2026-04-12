package mapper

import (
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

func ToCustomerEntity(req dto.CreateCustomerRequest) entity.Customer {
	return entity.Customer{
		UserID:      req.UserID,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		AvatarURL:   req.AvatarURL,
	}
}

func ApplyCustomerUpdate(customer *entity.Customer, req dto.UpdateCustomerRequest) {
	if req.UserID != nil {
		customer.UserID = *req.UserID
	}
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

func ToCustomerResponse(customer entity.Customer) dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          customer.ID,
		UserID:      customer.UserID,
		PhoneNumber: customer.PhoneNumber,
		Address:     customer.Address,
		AvatarURL:   customer.AvatarURL,
		CreatedAt:   customer.CreatedAt,
		UpdatedAt:   customer.UpdatedAt,
	}
}

func ToCustomersResponse(customers []entity.Customer) []dto.CustomerResponse {
	out := make([]dto.CustomerResponse, 0, len(customers))
	for _, item := range customers {
		out = append(out, ToCustomerResponse(item))
	}
	return out
}
