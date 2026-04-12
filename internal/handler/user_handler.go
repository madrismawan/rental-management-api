package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
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
		ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	created, err := h.svc.Create(ctx, mapper.ToUserEntity(req))
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, response{Message: "user created", Data: mapper.ToUserResponse(*created)})
}

func (h *UserHandler) List(ctx *gin.Context) {
	items, err := h.svc.List(ctx)
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "ok", Data: mapper.ToUsersResponse(items)})
}

func (h *UserHandler) GetByID(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, response{Message: "ok", Data: mapper.ToUserResponse(*item)})
}

func (h *UserHandler) Update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	var req dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	item, err := h.svc.Update(ctx, id, func(model *entity.User) {
		mapper.ApplyUserUpdate(model, req)
	})
	if err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "user updated", Data: mapper.ToUserResponse(*item)})
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response{Message: "invalid id"})
		return
	}
	if err := h.svc.Delete(ctx, id); err != nil {
		writeError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response{Message: "user deleted"})
}
