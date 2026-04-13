package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/mapper"
	"rental-management-api/internal/service"
)

type RentalHandler struct {
	svc service.RentalService
}

func NewRentalHandler(svc service.RentalService) *RentalHandler {
	return &RentalHandler{svc: svc}
}

func (h *RentalHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/rentals")
	r.POST("", h.Create)
	r.GET("", h.List)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

func (h *RentalHandler) Create(ctx *gin.Context) {
	var req dto.CreateRentalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}
	created, err := h.svc.Create(ctx, service.CreateRentalInput{
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
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, resource{Message: "rental created", Data: mapper.ToRentalResource(*created)})
}

func (h *RentalHandler) List(ctx *gin.Context) {
	items, err := h.svc.List(ctx)
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToRentalsResource(items)})
}

func (h *RentalHandler) GetByID(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	item, err := h.svc.GetByID(ctx, id)
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToRentalResource(*item)})
}

func (h *RentalHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	var req dto.UpdateRentalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}
	item, err := h.svc.Update(ctx, id, service.UpdateRentalInput{
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
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "rental updated", Data: mapper.ToRentalResource(*item)})
}

func (h *RentalHandler) Delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	if err := h.svc.Delete(ctx, id); err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "rental deleted"})
}
