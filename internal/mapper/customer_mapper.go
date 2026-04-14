package mapper

import (
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

type AvatarURLResolver interface {
	ResolveURL(objectRef string) (string, error)
}

func ToCustomerResource(customer entity.Customer, resolver AvatarURLResolver) dto.CustomerResource {
	res := dto.CustomerResource{
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

	if resolver != nil && res.AvatarURL != "" {
		resolvedURL, err := resolver.ResolveURL(res.AvatarURL)
		if err == nil {
			res.AvatarURL = resolvedURL
		}
	}

	return res
}

func ToCustomersResource(customers []entity.Customer, resolver AvatarURLResolver) []dto.CustomerResource {
	out := make([]dto.CustomerResource, 0, len(customers))
	for _, item := range customers {
		out = append(out, ToCustomerResource(item, resolver))
	}
	return out
}
