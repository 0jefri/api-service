package manager

import (
	ordershiping "github.com/api-service/api/order_shiping"
	"github.com/api-service/api/shiping"
)

type RepoManager interface {
	ShipingRepo() shiping.ShipingRepository
	OrderShipingRepo() ordershiping.OrderShippingRepository
}

type repoManager struct {
	infraManager InfraManager
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
