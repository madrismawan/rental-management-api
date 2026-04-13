package mapper

import (
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
)

func ToRentalEntity(req dto.CreateRentalRequest) entity.Rental {
	return entity.Rental{
		CustomerID:            req.CustomerID,
		VehicleID:             req.VehicleID,
		StartDate:             req.StartDate,
		EndDate:               req.EndDate,
		TotalDay:              req.TotalDay,
		ReturnDate:            req.ReturnDate,
		Price:                 req.Price,
		PenaltyFee:            req.PenaltyFee,
		Subtotal:              req.Subtotal,
		Notes:                 req.Notes,
		Status:                req.Status,
		VehicleConditionStart: req.VehicleConditionStart,
		VehicleConditionEnd:   req.VehicleConditionEnd,
		MileageStart:          req.MileageStart,
		MileageUsed:           req.MileageUsed,
		MileageEnd:            req.MileageEnd,
	}
}

func ApplyRentalUpdate(rental *entity.Rental, req dto.UpdateRentalRequest) {
	if req.CustomerID != nil {
		rental.CustomerID = *req.CustomerID
	}
	if req.VehicleID != nil {
		rental.VehicleID = *req.VehicleID
	}
	if req.StartDate != nil {
		rental.StartDate = *req.StartDate
	}
	if req.EndDate != nil {
		rental.EndDate = *req.EndDate
	}
	if req.TotalDay != nil {
		rental.TotalDay = *req.TotalDay
	}
	if req.ReturnDate != nil {
		rental.ReturnDate = req.ReturnDate
	}
	if req.Price != nil {
		rental.Price = *req.Price
	}
	if req.PenaltyFee != nil {
		rental.PenaltyFee = *req.PenaltyFee
	}
	if req.Subtotal != nil {
		rental.Subtotal = *req.Subtotal
	}
	if req.Notes != nil {
		rental.Notes = *req.Notes
	}
	if req.Status != nil {
		rental.Status = *req.Status
	}
	if req.VehicleConditionStart != nil {
		rental.VehicleConditionStart = *req.VehicleConditionStart
	}
	if req.VehicleConditionEnd != nil {
		rental.VehicleConditionEnd = *req.VehicleConditionEnd
	}
	if req.MileageStart != nil {
		rental.MileageStart = *req.MileageStart
	}
	if req.MileageUsed != nil {
		rental.MileageUsed = *req.MileageUsed
	}
	if req.MileageEnd != nil {
		rental.MileageEnd = *req.MileageEnd
	}
}

func ToRentalResource(rental entity.Rental) dto.RentalResource {
	return dto.RentalResource{
		ID:                    rental.ID,
		CustomerID:            rental.CustomerID,
		VehicleID:             rental.VehicleID,
		StartDate:             rental.StartDate,
		EndDate:               rental.EndDate,
		TotalDay:              rental.TotalDay,
		ReturnDate:            rental.ReturnDate,
		Price:                 rental.Price,
		PenaltyFee:            rental.PenaltyFee,
		Subtotal:              rental.Subtotal,
		Notes:                 rental.Notes,
		Status:                rental.Status,
		VehicleConditionStart: rental.VehicleConditionStart,
		VehicleConditionEnd:   rental.VehicleConditionEnd,
		MileageStart:          rental.MileageStart,
		MileageUsed:           rental.MileageUsed,
		MileageEnd:            rental.MileageEnd,
		CreatedAt:             rental.CreatedAt,
		UpdatedAt:             rental.UpdatedAt,
	}
}

func ToRentalsResource(rentals []entity.Rental) []dto.RentalResource {
	out := make([]dto.RentalResource, 0, len(rentals))
	for _, item := range rentals {
		out = append(out, ToRentalResource(item))
	}
	return out
}
