package manager

import (
	ordershiping "github.com/api-service/api/order_shiping"
	"github.com/api-service/api/shiping"
	statusorder "github.com/api-service/api/status_order"
)

type ServiceManager interface {
	ShipingService() shiping.ShipingService
	OrderShipingService() ordershiping.OrderShippingService
	StatusService() statusorder.StatusOrderServiceInterface
}

type serviceManager struct {
	repoManager RepoManager
}

// StatusService implements ServiceManager.
func (m *serviceManager) StatusService() statusorder.StatusOrderServiceInterface {
	return statusorder.NewStatusOrderService(m.repoManager.StatusOrder())
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
