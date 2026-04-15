package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rental-management-api/internal/dto"
	"rental-management-api/internal/mapper"
	"rental-management-api/internal/service"
)

type CustomerLogHandler struct {
	svc service.CustomerLogService
}

func NewCustomerLogHandler(svc service.CustomerLogService) *CustomerLogHandler {
	return &CustomerLogHandler{svc: svc}
}

func (h *CustomerLogHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/customer-logs")
	r.POST("", h.Create)
	r.GET("", h.List)
}

func (h *CustomerLogHandler) Create(ctx *gin.Context) {
	var req dto.CreateCustomerLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resource{Message: err.Error()})
		return
	}

	created, err := h.svc.Create(ctx, service.CreateCustomerLogInput{
		CustomerID:   req.CustomerID,
		CustomerName: req.CustomerName,
		Reason:       req.Reason,
		Status:       req.Status,
	})
	if err != nil {
		writeError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, resource{Message: "customer log created", Data: mapper.ToCustomerLogResource(*created)})
}

func (h *CustomerLogHandler) List(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK, resource{Message: "ok", Data: mapper.ToCustomerLogsResource(result.Items), Meta: paginationMeta(result.Page, result.Limit, result.Total, result.TotalPages)})
}
