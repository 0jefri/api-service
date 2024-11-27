package shiping

type Shiping struct {
	ID   string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Name string `gorm:"type:varchar(255);not null" json:"name" binding:"required,alphanum"`
}
