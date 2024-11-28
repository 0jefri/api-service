package seeder

import (
	"log"

	ordershiping "github.com/api-service/api/order_shiping"
	"gorm.io/gorm"
)

func SeedOrderShipping(db *gorm.DB) {
	data := []ordershiping.OrderShipping{
		{
			OrderID:              "order-001-shope-123",
			EcommerceID:          "ecommerce-uuid-002",
			ShippingID:           "2ca3934a-0bb1-4595-9e66-7e31a92e8a41",
			OriginLongitude:      "106.816666",
			DestinationLatitude:  "-6.200000",
			DestinationLongitude: "107.620000",
			TotalPaymentShipping: 150000.00,
		},
	}

	for _, order := range data {
		if err := db.Create(&order).Error; err != nil {
			log.Println("Failed to seed order shipping data:", err)
		}
	}
}
