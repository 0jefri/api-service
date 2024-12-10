package ordershiping

import (
	"errors"

	"gorm.io/gorm"
)

type OrderShippingRepository interface {
	Create(order OrderShipping) error
	FetchOrders() ([]OrderShipping, error)
}

type orderShippingRepository struct {
	db *gorm.DB
}

func NewOrderShippingRepository(db *gorm.DB) OrderShippingRepository {
	return &orderShippingRepository{db}
}

func (r *orderShippingRepository) Create(order OrderShipping) error {
	if err := r.db.Create(&order).Error; err != nil {
		return errors.New("failed to create order: " + err.Error())
	}
	return nil
}

func (r *orderShippingRepository) FetchOrders() ([]OrderShipping, error) {
	var orders []OrderShipping
	result := r.db.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
