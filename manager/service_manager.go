package manager

import "github.com/api-service/api/shiping"

type ServiceManager interface {
	ShipingService() shiping.ShipingService
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
