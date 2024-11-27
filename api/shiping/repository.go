package shiping

import (
	"errors"

	"gorm.io/gorm"
)

type ShipingRepository interface {
	Create(payload *Shiping) (*Shiping, error)
	List() (*[]Shiping, error)
	GetById(id string) (*Shiping, error)
}

type shipingRepository struct {
	db *gorm.DB
}

func NewShipingRepository(db *gorm.DB) ShipingRepository {
	return &shipingRepository{db: db}
}

func (r *shipingRepository) Create(payload *Shiping) (*Shiping, error) {
	shiping := Shiping{
		ID:   payload.ID,
		Name: payload.Name,
	}
	if err := r.db.Create(&shiping).Error; err != nil {
		return nil, errors.New("failed to create data")
	}
	return &shiping, nil
}

func (r *shipingRepository) List() (*[]Shiping, error) {
	var shipings []Shiping
	if err := r.db.Find(&shipings).Error; err != nil {
		return nil, errors.New("failed to retrieve data")
	}
	if len(shipings) == 0 {
		return nil, errors.New("no shiping data available")
	}
	return &shipings, nil
}

func (r *shipingRepository) GetById(id string) (*Shiping, error) {
	var shiping Shiping
	if err := r.db.First(&shiping, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("shiping not found")
		}
		return nil, errors.New("failed to retrieve data")
	}
	return &shiping, nil
}
