package destination

import (
	"gorm.io/gorm"
)

type Destination struct {
	gorm.Model
	OriginLat       string `gorm:"type:varchar(100);not null" json:"origin_lat"`
	OriginLong      string `gorm:"type:varchar(100);not null" json:"origin_long"`
	DestinationLat  string `gorm:"type:varchar(100);not null" json:"destination_lat"`
	DestinationLong string `gorm:"type:varchar(100);not null" json:"destination_long"`
}

type RequestDestination struct {
	ShippingID         string `json:"shipping_id" binding:"required" form:"shipping_id"`
	Qty                int    `json:"qty" binding:"required" form:"qty"`
	OriginLongLat      string `json:"origin_long_lat" binding:"required" form:"origin_long_lat"`
	DestinationLongLat string `json:"destination_long_lat" binding:"required" form:"destination_long_lat"`
}
