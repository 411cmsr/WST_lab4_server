package main

import (
	"WST_lab4_server/config"
	//"WST_lab4_server/config"
	"WST_lab4_server/internal/database/postgres"
	"WST_lab4_server/internal/httpserver/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	//"WST_lab4_server/internal/services"
	//"flag"
	//"fmt"
	//"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	postgres.Init()
	httpServer := gin.Default()

	routes.Init(httpServer)

	err := httpServer.Run(":8088")
	if err != nil {
		fmt.Println(err)
		return
	}
}
