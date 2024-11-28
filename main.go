package main

import (
	"github.com/api-service/config"
	"github.com/api-service/router"
)

func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
}
func main() {
	router.Server().Run()

}
