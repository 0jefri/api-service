package seeder

import (
	"log"

	ordershiping "github.com/api-service/api/order_shiping"
	"github.com/api-service/api/shiping"
	"gorm.io/gorm"
)

func SeedOrderShipping(db *gorm.DB) {
	data := []ordershiping.OrderShipping{
		{
			OrderID:     "order-009",
			EcommerceID: "1232112",
			Shipping: shiping.Shiping{
				ID:   "154c0255-198f-40e3-af94-61d28531b3ae",
				Name: "JNT",
			},
			OriginLongitude:      "106.84513",
			DestinationLatitude:  "-6.20876",
			DestinationLongitude: "107.62111",
			TotalPaymentShipping: 180000.50,
		},
	}

	for _, order := range data {
		if err := db.Create(&order).Error; err != nil {
			log.Println("Failed to seed order shipping data:", err)
		}
	}
}
