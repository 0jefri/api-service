package ordershiping

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderShippingHandler struct {
	service OrderShippingService
}

func NewOrderShippingHandler(service OrderShippingService) *OrderShippingHandler {
	return &OrderShippingHandler{service}
}

func (h *OrderShippingHandler) CreateOrder(c *gin.Context) {
	var order OrderShipping

	if err := c.ShouldBindJSON(&order); err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	fmt.Println(order)

	err := h.service.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
