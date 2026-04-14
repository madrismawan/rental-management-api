package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/constant"
	"rental-management-api/internal/dto"
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
	r.GET("/options", h.GetOptions)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

func (h *VehicleHandler) Create(ctx *gin.Context) {
	var req dto.CreateVehicleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}
	created, err := h.svc.Create(ctx, service.CreateVehicleInput{
		PlateNumber: req.PlateNumber,
		Color:       req.Color,
		Brand:       req.Brand,
		Model:       req.Model,
		CC:          req.CC,
		Year:        req.Year,
		Mileage:     req.Mileage,
		DailyRate:   req.DailyRate,
		Condition:   req.Condition,
		Status:      req.Status,
		Notes:       req.Notes,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, resource{Message: "vehicle created", Data: mapper.ToVehicleResource(*created)})
}

func (h *VehicleHandler) List(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToVehiclesResource(result.Items), Meta: paginationMeta(result.Page, result.Limit, result.Total, result.TotalPages)})
}

func (h *VehicleHandler) GetByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToVehicleResource(*item)})
}

func (h *VehicleHandler) GetOptions(ctx *gin.Context) {
	var status *constant.VehicleStatus
	if rawStatus := ctx.Query("status"); rawStatus != "" {
		parsedStatus := constant.VehicleStatus(rawStatus)
		status = &parsedStatus
	}

	items, err := h.svc.GetOptions(ctx, status)
	if err != nil {
		writeError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToVehicleOptionsResource(items)})
}

func (h *VehicleHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	var req dto.UpdateVehicleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}
	item, err := h.svc.Update(ctx, id, service.UpdateVehicleInput{
		PlateNumber: req.PlateNumber,
		Color:       req.Color,
		Brand:       req.Brand,
		Model:       req.Model,
		CC:          req.CC,
		Year:        req.Year,
		Mileage:     req.Mileage,
		DailyRate:   req.DailyRate,
		Condition:   req.Condition,
		Status:      req.Status,
		Notes:       req.Notes,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "vehicle updated", Data: mapper.ToVehicleResource(*item)})
}

func (h *VehicleHandler) Delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	if err := h.svc.Delete(ctx, id); err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "vehicle deleted"})
}
