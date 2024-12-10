package report

import ordershiping "github.com/api-service/api/order_shiping"

type DataStorage struct{}

func NewDataStorage() *DataStorage {
	return &DataStorage{}
}

func (s *DataStorage) GetOrders() []ordershiping.OrderShipping {
	return []ordershiping.OrderShipping{
		{
			OrderID:              "ORD123",
			EcommerceID:          "E001",
			ShippingID:           "SHIP001",
			OriginLongitude:      "106.816666",
			DestinationLatitude:  "-6.200000",
			DestinationLongitude: "106.850000",
			TotalPaymentShipping: 50000.00,
		},
		{
			OrderID:              "ORD124",
			EcommerceID:          "E002",
			ShippingID:           "SHIP002",
			OriginLongitude:      "107.616666",
			DestinationLatitude:  "-7.150000",
			DestinationLongitude: "107.850000",
			TotalPaymentShipping: 75000.00,
		},
	}
}
