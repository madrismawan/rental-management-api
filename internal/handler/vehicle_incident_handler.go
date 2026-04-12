package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/mapper"
	"rental-management-api/internal/service"
)

type VehicleIncidentHandler struct {
	svc service.VehicleIncidentService
}

func NewVehicleIncidentHandler(svc service.VehicleIncidentService) *VehicleIncidentHandler {
	return &VehicleIncidentHandler{svc: svc}
}

func (h *VehicleIncidentHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/vehicle-incidents")
	r.POST("", h.Create)
	r.GET("", h.List)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

func (h *VehicleIncidentHandler) Create(ctx *gin.Context) {
	var req dto.CreateVehicleIncidentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	created, err := h.svc.Create(ctx, mapper.ToVehicleIncidentEntity(req))
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, response{Message: "vehicle incident created", Data: mapper.ToVehicleIncidentResponse(*created)})
}

func (h *VehicleIncidentHandler) List(ctx *gin.Context) {
	items, err := h.svc.List(ctx)
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "ok", Data: mapper.ToVehicleIncidentsResponse(items)})
}

func (h *VehicleIncidentHandler) GetByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, response{Message: "ok", Data: mapper.ToVehicleIncidentResponse(*item)})
}

func (h *VehicleIncidentHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	var req dto.UpdateVehicleIncidentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	item, err := h.svc.Update(ctx, id, func(model *entity.VehicleIncident) {
		mapper.ApplyVehicleIncidentUpdate(model, req)
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "vehicle incident updated", Data: mapper.ToVehicleIncidentResponse(*item)})
}

func (h *VehicleIncidentHandler) Delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	if err := h.svc.Delete(ctx, id); err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "vehicle incident deleted"})
}
