package statusorder

import "github.com/google/uuid"

type StatusOrderServiceInterface interface {
	CreateStatusOrder(input StatusOder) (StatusOder, error)
	GetAllStatusOrders() ([]StatusOder, error)
}

type StatusOrderService struct {
	repository StatusOrderRepository
}

// NewStatusOrderService adalah constructor untuk StatusOrderService
func NewStatusOrderService(repo StatusOrderRepository) StatusOrderServiceInterface {
	return &StatusOrderService{repository: repo}
}

// CreateStatusOrder membuat status order baru
func (s *StatusOrderService) CreateStatusOrder(input StatusOder) (StatusOder, error) {
	statusOrder := StatusOder{
		ID:             uuid.NewString(),
		OrderShipingID: input.OrderShipingID,
		Status:         input.Status,
		Location:       input.Location,
	}
	return s.repository.Create(statusOrder)
}

// GetAllStatusOrders mendapatkan semua status order
func (s *StatusOrderService) GetAllStatusOrders() ([]StatusOder, error) {
	return s.repository.GetAll()
}
