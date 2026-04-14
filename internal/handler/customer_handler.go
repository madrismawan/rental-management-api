package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
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
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}
	created, err := h.svc.CreateWithUser(ctx, service.CreateCustomerWithUserInput{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		AvatarURL:   req.AvatarURL,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, resource{Message: "customer created", Data: mapper.ToCustomerResource(*created)})
}

func (h *CustomerHandler) List(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToCustomersResource(result.Items), Meta: paginationMeta(result.Page, result.Limit, result.Total, result.TotalPages)})
}

func (h *CustomerHandler) GetByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToCustomerResource(*item)})
}

func (h *CustomerHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	var req dto.UpdateCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}
	item, err := h.svc.Update(ctx, id, service.UpdateCustomerInput{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		AvatarURL:   req.AvatarURL,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "customer updated", Data: mapper.ToCustomerResource(*item)})
}

func (h *CustomerHandler) Delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	if err := h.svc.Delete(ctx, id); err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "customer deleted"})
}
