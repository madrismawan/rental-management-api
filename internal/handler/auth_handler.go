package handler

import (
	"net/http"
	"rental-management-api/internal/dto"
	"rental-management-api/internal/mapper"
	"rental-management-api/internal/service"
	"rental-management-api/pkg/errs"

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
	group.POST("/login", h.login)
}

func (h *AuthHandler) login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, tokens, err := h.authService.Login(c, req.Email, req.Password)
	if err != nil {
		if err == errs.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data": dto.LoginResponse{
			User:   mapper.ToUserResource(user),
			Tokens: tokens,
		},
	})
}
