package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/mapper"
	"rental-management-api/internal/service"
)

type CustomerHandler struct {
	svc service.CustomerService
}

func NewCustomerHandler(svc service.CustomerService) *CustomerHandler {
	return &CustomerHandler{svc: svc}
}

func (h *CustomerHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/customers")
	r.POST("", h.Create)
	r.GET("", h.List)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

func (h *CustomerHandler) Create(ctx *gin.Context) {
	var req dto.CreateCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	created, err := h.svc.Create(ctx, mapper.ToCustomerEntity(req))
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, response{Message: "customer created", Data: mapper.ToCustomerResponse(*created)})
}

func (h *CustomerHandler) List(ctx *gin.Context) {
	items, err := h.svc.List(ctx)
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "ok", Data: mapper.ToCustomersResponse(items)})
}

func (h *CustomerHandler) GetByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, response{Message: "ok", Data: mapper.ToCustomerResponse(*item)})
}

func (h *CustomerHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	var req dto.UpdateCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	item, err := h.svc.Update(ctx, id, func(model *entity.Customer) {
		mapper.ApplyCustomerUpdate(model, req)
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "customer updated", Data: mapper.ToCustomerResponse(*item)})
}

func (h *CustomerHandler) Delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	if err := h.svc.Delete(ctx, id); err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "customer deleted"})
}
