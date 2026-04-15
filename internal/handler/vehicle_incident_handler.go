package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
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
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}
	created, err := h.svc.Create(ctx, service.CreateVehicleIncidentInput{
		VehicleID:    req.VehicleID,
		CustomerID:   req.CustomerID,
		RentalID:     req.RentalID,
		IncidentDate: req.IncidentDate,
		IncidentType: req.IncidentType,
		Description:  req.Description,
		Cost:         req.Cost,
		Status:       req.Status,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, resource{Message: "vehicle incident created", Data: mapper.ToVehicleIncidentResource(*created)})
}

func (h *VehicleIncidentHandler) List(ctx *gin.Context) {
	page, limit, err := parsePaginationQuery(ctx, 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}

	result, err := h.svc.ListPaginated(ctx, page, limit)
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToVehicleIncidentsResource(result.Items), Meta: paginationMeta(result.Page, result.Limit, result.Total, result.TotalPages)})
}

func (h *VehicleIncidentHandler) GetByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToVehicleIncidentResource(*item)})
}

func (h *VehicleIncidentHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	var req dto.UpdateVehicleIncidentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}
	item, err := h.svc.Update(ctx, id, service.UpdateVehicleIncidentInput{
		VehicleID:    req.VehicleID,
		CustomerID:   req.CustomerID,
		RentalID:     req.RentalID,
		IncidentDate: req.IncidentDate,
		IncidentType: req.IncidentType,
		Description:  req.Description,
		Cost:         req.Cost,
		Status:       req.Status,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "vehicle incident updated", Data: mapper.ToVehicleIncidentResource(*item)})
}

func (h *VehicleIncidentHandler) Delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	if err := h.svc.Delete(ctx, id); err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "vehicle incident deleted"})
}
