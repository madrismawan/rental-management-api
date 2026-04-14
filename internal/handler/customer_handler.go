package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/mapper"
	"rental-management-api/internal/service"
)

type CustomerHandler struct {
	svc        service.CustomerService
	storageSvc service.StorageService
}

func NewCustomerHandler(svc service.CustomerService, storageSvc service.StorageService) *CustomerHandler {
	return &CustomerHandler{svc: svc, storageSvc: storageSvc}
}

func (h *CustomerHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/customers")
	r.POST("", h.Create)
	r.GET("", h.List)
	r.GET("/options", h.GetOptions)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

func (h *CustomerHandler) Create(ctx *gin.Context) {
	var req dto.CreateCustomerRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}

	avatarURL := ""
	if req.Avatar != nil {
		uploadedURL, err := h.storageSvc.Upload(ctx, req.Avatar, "customers/avatars")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resource{Message: err.Error()})
			return
		}
		avatarURL = uploadedURL
	}

	created, err := h.svc.CreateWithUser(ctx, service.CreateCustomerWithUserInput{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		AvatarURL:   avatarURL,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, resource{Message: "customer created", Data: mapper.ToCustomerResource(*created, h.storageSvc)})
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
	resources := mapper.ToCustomersResource(result.Items, h.storageSvc)
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: resources, Meta: paginationMeta(result.Page, result.Limit, result.Total, result.TotalPages)})
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
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToCustomerResource(*item, h.storageSvc)})
}

func (h *CustomerHandler) GetOptions(ctx *gin.Context) {
	items, err := h.svc.GetOptions(ctx)
	if err != nil {
		writeError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToCustomerOptionsResource(items)})
}

func (h *CustomerHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	var req dto.UpdateCustomerRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}

	avatarURL := req.AvatarURL
	if req.Avatar != nil {
		uploadedURL, err := h.storageSvc.Upload(ctx, req.Avatar, "customers/avatars")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resource{Message: err.Error()})
			return
		}
		avatarURL = &uploadedURL
	}

	item, err := h.svc.Update(ctx, id, service.UpdateCustomerInput{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		AvatarURL:   avatarURL,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "customer updated", Data: mapper.ToCustomerResource(*item, h.storageSvc)})
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
