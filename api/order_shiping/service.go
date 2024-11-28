package ordershiping

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

type OrderShippingService interface {
	CreateOrder(order OrderShipping) error
}

type orderShippingService struct {
	repo OrderShippingRepository
}

func NewOrderShippingService(repo OrderShippingRepository) OrderShippingService {
	return &orderShippingService{repo}
}

func (s *orderShippingService) CreateOrder(order OrderShipping) error {
	if order.EcommerceID == "" {
		return errors.New("ecommerceID is required")
	}

	if order.OrderID == "" || order.Shipping.ID == "" {
		log.Println("eror :", order.OrderID, order.ShippingID)
		return errors.New("orderID and shippingID are required")
	}

	order.EcommerceID = uuid.NewString()

	return s.repo.Create(order)
}
