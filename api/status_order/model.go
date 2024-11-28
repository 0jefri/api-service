package statusorder

type StatusOder struct {
	ID             string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	OrderShipingID string `gorm:"type:varchar(100);not null" json:"order_shiping_id"`
	Status         string `gorm:"type:varchar(100)" json:"status"`
	Location       string `gorm:"type:varchar(100);not null" json:"location"`
}
