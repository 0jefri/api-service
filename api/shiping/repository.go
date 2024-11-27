package shiping

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/api-service/api/destination"
	"github.com/api-service/helper"
	"gorm.io/gorm"
)

type ShipingRepository interface {
	Create(payload *Shiping) (*Shiping, error)
	List() (*[]Shiping, error)
	GetById(id string) (*Shiping, error)
	GetDestination(destination destination.RequestDestination) (*float64, error)
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

func (r *shipingRepository) GetDestination(destination destination.RequestDestination) (*float64, error) {

	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/driving/%s;%s?overview=false",
		destination.OriginLongLat, destination.DestinationLongLat)
	var header http.Header
	data, err := helper.HTTPRequest("GET", header, url, nil)
	if err != nil {
		return nil, errors.New("error http request direction")
	}

	var dataMap map[string]interface{}

	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		return nil, errors.New("error decode data")
	}

	routes, ok := dataMap["routes"].([]interface{})
	if !ok {
		return nil, errors.New("error decode routes")
	}

	// Periksa jika routes ada dan tidak kosong
	if len(routes) == 0 {
		return nil, errors.New("routes array is empty")
	}

	// Ambil elemen pertama dari slice routes
	route, ok := routes[0].(map[string]interface{})
	if !ok {
		return nil, errors.New("error decoding route")
	}

	distance, ok := route["distance"].(float64)
	if !ok {
		return nil, errors.New("error decode distance")
	}

	return &distance, nil
}
