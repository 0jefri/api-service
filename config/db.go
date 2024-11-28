package config

import (
	"fmt"
	"sync"
	"time"

	ordershiping "github.com/api-service/api/order_shiping"
	"github.com/api-service/api/seeder"
	"github.com/api-service/api/shiping"
	statusorder "github.com/api-service/api/status_order"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", Cfg.Database.Host, Cfg.Database.Username, Cfg.Database.Password, Cfg.Database.Dbname, Cfg.Database.Port)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	once.Do(func() {
		DB = Db
		fmt.Println("Successfully Connected To Database!")
	})
}

func SyncDB() {
	if err := DB.AutoMigrate(&shiping.Shiping{}, &ordershiping.OrderShipping{}, &statusorder.StatusOder{}); err != nil {
		fmt.Printf("AutoMigrate error: %s\n", err)
		panic(err)
	} else {
		fmt.Println("Database migrated successfully!")
	}

	seeder.SeedOrderShipping(DB)
	// if err := seeder.SeedOrderShipping(DB); err != nil {
	// 	panic("Failed to seed data")
	// }
}
