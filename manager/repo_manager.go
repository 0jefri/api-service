package manager

import "github.com/api-service/api/shiping"

type RepoManager interface {
	// UserRepo() repository.UserRepository
	ShipingRepo() shiping.ShipingRepository
}

type repoManager struct {
	infraManager InfraManager
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}

// func (m *repoManager) UserRepo() repository.UserRepository {
// 	return repository.NewUserRepository(m.infraManager.Conn())
// }

func (m *repoManager) ShipingRepo() shiping.ShipingRepository {
	return shiping.NewShipingRepository(m.infraManager.Conn())
}
