package manager

import (
	ordershiping "github.com/api-service/api/order_shiping"
	"github.com/api-service/api/shiping"
)

type ServiceManager interface {
	ShipingService() shiping.ShipingService
	OrderShipingService() ordershiping.OrderShippingService
}

type serviceManager struct {
	repoManager RepoManager
}

func NewServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{
		repoManager: repo,
	}
}

func (m *serviceManager) ShipingService() shiping.ShipingService {
	return shiping.NewShipingService(m.repoManager.ShipingRepo())
}

func (m *serviceManager) OrderShipingService() ordershiping.OrderShippingService {
	return ordershiping.NewOrderShippingService(m.repoManager.OrderShipingRepo())
}
