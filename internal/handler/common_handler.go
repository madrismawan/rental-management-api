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
	Meta    any    `json:"meta,omitempty"`
}

var (
	errInvalidPage  = errors.New("invalid page")
	errInvalidLimit = errors.New("invalid limit")
)

func parseID(ctx *gin.Context) (uint, error) {
	rawID := ctx.Param("id")
	parsed, err := strconv.ParseUint(rawID, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(parsed), nil
}

func parsePaginationQuery(ctx *gin.Context, defaultLimit int) (int, int, error) {
	page := 1
	limit := defaultLimit
	if limit <= 0 {
		limit = 10
	}

	if rawPage := ctx.Query("page"); rawPage != "" {
		parsedPage, err := strconv.Atoi(rawPage)
		if err != nil || parsedPage <= 0 {
			return 0, 0, errInvalidPage
		}
		page = parsedPage
	}

	if rawLimit := ctx.Query("limit"); rawLimit != "" {
		parsedLimit, err := strconv.Atoi(rawLimit)
		if err != nil || parsedLimit <= 0 {
			return 0, 0, errInvalidLimit
		}
		limit = parsedLimit
	}

	return page, limit, nil
}

func paginationMeta(page int, limit int, total int64, totalPages int) gin.H {
	return gin.H{
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	}
}

func writeError(ctx *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, resource{Message: "resource not found"})
		return
	}
	ctx.JSON(http.StatusInternalServerError, resource{Message: err.Error()})
}
