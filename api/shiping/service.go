package shiping

import "errors"

type ShipingService interface {
	CreateNewShiping(payload *Shiping) (*Shiping, error)
	GetAllShipings() (*[]Shiping, error)
	GetShipingById(id string) (*Shiping, error)
}

type shipingService struct {
	repo ShipingRepository
}

func NewShipingService(repo ShipingRepository) ShipingService {
	return &shipingService{repo: repo}
}

func (s *shipingService) CreateNewShiping(payload *Shiping) (*Shiping, error) {
	if payload.Name == "" {
		return nil, errors.New("name is required")
	}
	return s.repo.Create(payload)
}

func (s *shipingService) GetAllShipings() (*[]Shiping, error) {
	return s.repo.List()
}

func (s *shipingService) GetShipingById(id string) (*Shiping, error) {
	data, err := s.repo.GetById(id)
	if err != nil {
		return nil, errors.New("failed to retrieve shiping by id")
	}
	return data, nil
}
