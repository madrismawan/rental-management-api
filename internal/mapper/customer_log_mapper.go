package mapper

import (
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

func ToCustomerLogResource(item entity.CustomerLog) dto.CustomerLogResource {
	return dto.CustomerLogResource{
		ID:           item.ID,
		CustomerID:   item.CustomerID,
		CustomerName: item.CustomerName,
		Reason:       item.Reason,
		Status:       item.Status,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}
}

func ToCustomerLogsResource(items []entity.CustomerLog) []dto.CustomerLogResource {
	out := make([]dto.CustomerLogResource, 0, len(items))
	for _, item := range items {
		out = append(out, ToCustomerLogResource(item))
	}
	return out
}
