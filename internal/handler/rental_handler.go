package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
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
		ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	created, err := h.svc.Create(ctx, mapper.ToRentalEntity(req))
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, response{Message: "rental created", Data: mapper.ToRentalResponse(*created)})
}

func (h *RentalHandler) List(ctx *gin.Context) {
	items, err := h.svc.List(ctx)
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "ok", Data: mapper.ToRentalsResponse(items)})
}

func (h *RentalHandler) GetByID(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	item, err := h.svc.GetByID(ctx, id)
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "ok", Data: mapper.ToRentalResponse(*item)})
}

func (h *RentalHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	var req dto.UpdateRentalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	item, err := h.svc.Update(ctx, id, func(model *entity.Rental) {
		mapper.ApplyRentalUpdate(model, req)
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "rental updated", Data: mapper.ToRentalResponse(*item)})
}

func (h *RentalHandler) Delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	if err := h.svc.Delete(ctx, id); err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "rental deleted"})
}
