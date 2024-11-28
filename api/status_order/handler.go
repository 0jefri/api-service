package statusorder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StatusHandler struct {
	service StatusOrderServiceInterface
}

func NewHandler(service StatusOrderServiceInterface) *StatusHandler {
	return &StatusHandler{service: service}
}

func (h *StatusHandler) CreateStatusOrder(c *gin.Context) {
	var input StatusOder
	input.ID = uuid.NewString()
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusOrder, err := h.service.CreateStatusOrder(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Status order created successfully",
		"status_order": statusOrder,
	})
}

func (h *StatusHandler) GetStatusOrders(c *gin.Context) {
	statusOrders, err := h.service.GetAllStatusOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_orders": statusOrders,
	})
}
