package handler

import (
	"rental-management-api/internal/dto"
	"rental-management-api/internal/mapper"
	"rental-management-api/internal/service"
	"rental-management-api/pkg/errs"
	"rental-management-api/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) RegisterRoutes(router gin.IRouter) {
	group := router.Group("/auth")
	group.POST("/register", h.register)
	group.POST("/login", h.login)
	group.POST("/refresh-token", h.login)
}

func (h *AuthHandler) register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.BadRequest(c, httpx.ErrInvalidJSON, "", err.Error())
		return
	}

	res, err := h.authService.Register(c, req)
	if err != nil {
		switch err {
		case errs.ErrEmailDuplicate:
			httpx.Conflict(c, httpx.ErrConflict, "email already registered")
			return
		default:
			httpx.InternalError(c, err)
			return
		}
	}

	httpx.Success(c, "Register successful", gin.H{
		"user": mapper.ToUserResource(res),
	})
}

func (h *AuthHandler) login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.BadRequest(c, httpx.ErrInvalidJSON, "", err.Error())
		return
	}

	user, tokens, err := h.authService.Login(c, req.Email, req.Password)
	if err != nil {
		if err == errs.ErrInvalidCredentials {
			httpx.Unauthorized(c, httpx.ErrInvalidCredentials, "invalid email or password")
			return
		}
		httpx.InternalError(c, err)
		return
	}

	httpx.Success(c, "Login successful", dto.LoginResponse{
		User:   mapper.ToUserResource(user),
		Tokens: tokens,
	})
}
