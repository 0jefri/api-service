package ordershiping

import "github.com/api-service/api/shiping"

type OrderShipping struct {
	OrderID              string          `gorm:"type:varchar(255);primaryKey;not null;unique" json:"order_id" binding:"required"`
	EcommerceID          string          `gorm:"type:uuid;primaryKey;not null;unique" json:"ecommerce_ID" binding:"required"`
	ShippingID           string          `gorm:"type:uuid;not null"`
	Shipping             shiping.Shiping `gorm:"foreignKey:ShippingID;references:ID" json:"shipping"`
	OriginLongitude      string          `gorm:"type:varchar(100);not null" json:"origin_longitude"`
	DestinationLatitude  string          `gorm:"type:varchar(100);not null" json:"destination_latitude"`
	DestinationLongitude string          `gorm:"type:varchar(100);not null" json:"destination_longitude"`
	TotalPaymentShipping float64         `gorm:"type:float;not null;default:0;" json:"total_payment_shipping"`
}
