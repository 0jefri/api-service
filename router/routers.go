package router

import (
	ordershiping "github.com/api-service/api/order_shiping"
	"github.com/api-service/api/shiping"
	"github.com/api-service/config"
	"github.com/api-service/manager"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) error {

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	shipingHandler := shiping.NewShipingHandler(repoManager.ShipingService())
	orderShipingHandler := ordershiping.NewOrderShippingHandler(repoManager.OrderShipingService())

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
		}
	}

	return router.Run()

}
