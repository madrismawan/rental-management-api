package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/mapper"
	"rental-management-api/internal/service"
)

type VehicleHandler struct {
	svc service.VehicleService
}

func NewVehicleHandler(svc service.VehicleService) *VehicleHandler {
	return &VehicleHandler{svc: svc}
}

func (h *VehicleHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/vehicles")
	r.POST("", h.Create)
	r.GET("", h.List)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

func (h *VehicleHandler) Create(ctx *gin.Context) {
	var req dto.CreateVehicleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	created, err := h.svc.Create(ctx, mapper.ToVehicleEntity(req))
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, response{Message: "vehicle created", Data: mapper.ToVehicleResponse(*created)})
}

func (h *VehicleHandler) List(ctx *gin.Context) {
	items, err := h.svc.List(ctx)
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "ok", Data: mapper.ToVehiclesResponse(items)})
}

func (h *VehicleHandler) GetByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, response{Message: "ok", Data: mapper.ToVehicleResponse(*item)})
}

func (h *VehicleHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	var req dto.UpdateVehicleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	item, err := h.svc.Update(ctx, id, func(model *entity.Vehicle) {
		mapper.ApplyVehicleUpdate(model, req)
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "vehicle updated", Data: mapper.ToVehicleResponse(*item)})
}

func (h *VehicleHandler) Delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	if err := h.svc.Delete(ctx, id); err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "vehicle deleted"})
}
