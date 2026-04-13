package appctx

import (
	"errors"
	"rental-management-api/internal/dto"

	"github.com/gin-gonic/gin"
)

type UserContext struct {
	User dto.UserResource
}

const (
	UserContextKey string = "user_context"
)

func SetUser(c *gin.Context, user dto.UserResource) {
	c.Set(UserContextKey, user)
}

func GetUserID(c *gin.Context) (uint, error) {
	userID, exists := c.Get(UserContextKey)
	if !exists {
		return 0, errors.New("user not authenticated")
	}

	user, ok := userID.(dto.UserResource)
	if !ok {
		return 0, errors.New("invalid userID type")
	}
	return user.ID, nil
}
