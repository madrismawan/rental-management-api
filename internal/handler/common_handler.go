package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type resource struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func parseID(ctx *gin.Context) (uint, error) {
	rawID := ctx.Param("id")
	parsed, err := strconv.ParseUint(rawID, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(parsed), nil
}

func writeError(ctx *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, resource{Message: "resource not found"})
		return
	}
	ctx.JSON(http.StatusInternalServerError, resource{Message: err.Error()})
}
