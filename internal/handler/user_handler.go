package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/mapper"
	"rental-management-api/internal/service"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/users")
	r.POST("", h.Create)
	r.GET("", h.List)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

func (h *UserHandler) Create(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}
	created, err := h.svc.Create(ctx, service.CreateUserInput{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Password: req.Password,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, resource{Message: "user created", Data: mapper.ToUserResource(*created)})
}

func (h *UserHandler) List(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToUsersResource(result.Items), Meta: paginationMeta(result.Page, result.Limit, result.Total, result.TotalPages)})
}

func (h *UserHandler) GetByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToUserResource(*item)})
}

func (h *UserHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	var req dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}
	item, err := h.svc.Update(ctx, id, service.UpdateUserInput{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Password: req.Password,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "user updated", Data: mapper.ToUserResource(*item)})
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: "invalid id"})
		return
	}
	if err := h.svc.Delete(ctx, id); err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resource{Message: "user deleted"})
}
