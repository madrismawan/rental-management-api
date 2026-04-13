package httpx

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Standard response structures
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *Error      `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type Meta struct {
	Timestamp  string      `json:"timestamp"`
	RequestID  string      `json:"requestId,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	Total      int  `json:"total"`
	TotalPages int  `json:"totalPages"`
	HasNext    bool `json:"hasNext"`
	HasPrev    bool `json:"hasPrev"`
}

// Success responses
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta: &Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func SuccessCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta: &Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func SuccessWithPagination(c *gin.Context, message string, data interface{}, pagination Pagination) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta: &Meta{
			Timestamp:  time.Now().UTC().Format(time.RFC3339),
			Pagination: &pagination,
		},
	})
}

// Error responses
func BadRequest(c *gin.Context, code, message string, details interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: message,
		Error: &Error{
			Code:    code,
			Message: message,
			Details: details,
		},
		Meta: &Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func Unauthorized(c *gin.Context, code, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Message: message,
		Error: &Error{
			Code:    code,
			Message: message,
		},
		Meta: &Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func Forbidden(c *gin.Context, code, message string) {
	c.JSON(http.StatusForbidden, Response{
		Success: false,
		Message: message,
		Error: &Error{
			Code:    code,
			Message: message,
		},
		Meta: &Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func NotFound(c *gin.Context, code, message string) {
	c.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: message,
		Error: &Error{
			Code:    code,
			Message: message,
		},
		Meta: &Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func Conflict(c *gin.Context, code, message string) {
	c.JSON(http.StatusConflict, Response{
		Success: false,
		Message: message,
		Error: &Error{
			Code:    code,
			Message: message,
		},
		Meta: &Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func InternalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Message: "Internal server error",
		Error: &Error{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
		},
		Meta: &Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func ValidationError(c *gin.Context, details interface{}) {
	c.JSON(http.StatusUnprocessableEntity, Response{
		Success: false,
		Message: "Validation failed",
		Error: &Error{
			Code:    "VALIDATION_ERROR",
			Message: "The request contains invalid data",
			Details: details,
		},
		Meta: &Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

// Helper for common error codes
var (
	ErrInvalidJSON        = "INVALID_JSON"
	ErrValidation         = "VALIDATION_ERROR"
	ErrNotFound           = "NOT_FOUND"
	ErrUnauthorized       = "UNAUTHORIZED"
	ErrForbidden          = "FORBIDDEN"
	ErrConflict           = "CONFLICT"
	ErrBadRequest         = "BAD_REQUEST"
	ErrInternal           = "INTERNAL_ERROR"
	ErrInvalidCredentials = "INVALID_CREDENTIALS"
	ErrInvalidID          = "INVALID_ID"
)
