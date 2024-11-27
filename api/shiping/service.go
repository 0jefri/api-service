package shiping

import (
	"errors"
	"fmt"

	"github.com/api-service/api/destination"
)

type ShipingService interface {
	CreateNewShiping(payload *Shiping) (*Shiping, error)
	GetAllShipings() (*[]Shiping, error)
	GetShipingById(id string) (*Shiping, error)
	CalculateCost(payload destination.RequestDestination) (*float64, error)
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

func (s *shipingService) CalculateCost(payload destination.RequestDestination) (*float64, error) {
	// Mendapatkan jarak dari fungsi repository
	distance, err := s.repo.GetDestination(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to get destination: %w", err)
	}

	if distance == nil || *distance <= 0 {
		return nil, errors.New("invalid distance value")
	}

	// Kalkulasi ongkos berdasarkan jumlah barang dan jarak
	var costPerKm float64
	if payload.Qty <= 2 {
		costPerKm = 2000
	} else {
		costPerKm = 4000
	}

	cost := *distance * costPerKm
	return &cost, nil
}
