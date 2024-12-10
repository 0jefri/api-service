package router

import (
	"fmt"
	"net/http"
	"os"

	ordershiping "github.com/api-service/api/order_shiping"
	"github.com/api-service/api/report"
	"github.com/api-service/api/shiping"
	statusorder "github.com/api-service/api/status_order"
	"github.com/api-service/config"
	"github.com/api-service/manager"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func SetupRouter(router *gin.Engine) error {

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	// Pastikan folder laporan tersedia
	err := os.MkdirAll("reports", os.ModePerm)
	if err != nil {
		fmt.Println("Gagal membuat folder reports:", err)
		return err
	}

	db := config.DB

	// Inisialisasi service
	orderRepo := ordershiping.NewOrderShippingRepository(db)
	reportService := report.NewReportService(orderRepo)

	// Menjalankan cron job untuk pembuatan laporan
	c := cron.New()
	err = c.AddFunc("@every 5s", func() {
		if err := reportService.GenerateOrderReport(); err != nil {
			fmt.Println("Gagal membuat laporan:", err)
		}
	})
	if err != nil {
		fmt.Println("Gagal menambahkan cron job:", err)
		return err
	}
	c.Start()

	// Agar program tidak langsung selesai
	// select {}

	shipingHandler := shiping.NewShipingHandler(repoManager.ShipingService())
	orderShipingHandler := ordershiping.NewOrderShippingHandler(repoManager.OrderShipingService())
	statusOrderHandler := statusorder.NewHandler(repoManager.StatusService())

	v1 := router.Group("/api/v1")
	{
		eCommerce := v1.Group("/e-commerce")
		{
			shiping := eCommerce.Group("/shiping")
			{
				shiping.POST("/register-shiping", shipingHandler.AddShiping)
				shiping.GET("/list-shiping", shipingHandler.ListShipings)
				shiping.GET("/:id", shipingHandler.GetShipingById)
				shiping.GET("/cost", shipingHandler.CalculateShippingCost)
			}
			orderShiping := eCommerce.Group("/order")
			{
				orderShiping.POST("/create", orderShipingHandler.CreateOrder)
			}
			statusOrder := eCommerce.Group("/status")
			{
				statusOrder.POST("/create", statusOrderHandler.CreateStatusOrder)
				statusOrder.GET("/list-status", statusOrderHandler.GetStatusOrders)
			}
		}
		http.HandleFunc("/trigger-report", report.TriggerReportHandler(reportService))
	}

	return router.Run()

}
