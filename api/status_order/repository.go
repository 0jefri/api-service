package statusorder

import "gorm.io/gorm"

// StatusOrderRepository adalah interface untuk repository StatusOrder
type StatusOrderRepository interface {
	Create(status StatusOder) (StatusOder, error)
	GetAll() ([]StatusOder, error)
}

type StatusOrderRepositoryImpl struct {
	db *gorm.DB
}

func NewStatusOrderRepository(db *gorm.DB) StatusOrderRepository {
	return &StatusOrderRepositoryImpl{db: db}
}

func (r *StatusOrderRepositoryImpl) Create(status StatusOder) (StatusOder, error) {
	if err := r.db.Create(&status).Error; err != nil {
		return StatusOder{}, err
	}
	return status, nil
}

func (r *StatusOrderRepositoryImpl) GetAll() ([]StatusOder, error) {
	var statusOrders []StatusOder
	if err := r.db.Find(&statusOrders).Error; err != nil {
		return nil, err
	}
	return statusOrders, nil
}
