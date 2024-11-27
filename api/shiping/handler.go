package shiping

import (
	"errors"
	"log"
	"net/http"

	"github.com/api-service/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type shipingHandler struct {
	service ShipingService
}

func NewShipingHandler(service ShipingService) *shipingHandler {
	return &shipingHandler{service: service}
}

func (h *shipingHandler) AddShiping(c *gin.Context) {
	var payload Shiping

	payload.ID = uuid.NewString()

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	data, err := h.service.CreateNewShiping(&payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  "Error",
				Message: gorm.ErrRecordNotFound.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: "Create shiping",
		Data:    data,
	})
	log.Println("data: ", data)
}

func (h *shipingHandler) ListShipings(c *gin.Context) {
	data, err := h.service.GetAllShipings()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{
				Code:    http.StatusNotFound,
				Status:  "Error",
				Message: "No shiping data found",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Status:  "Success",
		Message: "Shiping data retrieved",
		Data:    data,
	})
}

func (h *shipingHandler) GetShipingById(c *gin.Context) {
	id := c.Param("id")
	data, err := h.service.GetShipingById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"status":  "Not Found",
				"message": "Shiping not found",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success",
		"message": "Shiping retrieved successfully",
		"data":    data,
	})
}
