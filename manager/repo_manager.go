package manager

import (
	ordershiping "github.com/api-service/api/order_shiping"
	"github.com/api-service/api/shiping"
	statusorder "github.com/api-service/api/status_order"
)

type RepoManager interface {
	ShipingRepo() shiping.ShipingRepository
	OrderShipingRepo() ordershiping.OrderShippingRepository
	StatusOrder() statusorder.StatusOrderRepository
}

type repoManager struct {
	infraManager InfraManager
}

// StatusOrder implements RepoManager.
func (m *repoManager) StatusOrder() statusorder.StatusOrderRepository {
	return statusorder.NewStatusOrderRepository(m.infraManager.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}

func (m *repoManager) ShipingRepo() shiping.ShipingRepository {
	return shiping.NewShipingRepository(m.infraManager.Conn())
}

func (m *repoManager) OrderShipingRepo() ordershiping.OrderShippingRepository {
	return ordershiping.NewOrderShippingRepository(m.infraManager.Conn())
}

// func (m *repoManager) StatusOder() statusorder.StatusRepository {
// 	return statusorder.NewRepository(m.infraManager.Conn())
// }
