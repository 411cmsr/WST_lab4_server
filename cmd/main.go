package main

import (
	"WST_lab4_server/config"
	"WST_lab4_server/internal/database/postgres"
	"WST_lab4_server/internal/handlers"
	"WST_lab4_server/internal/httpserver/routes"
	"WST_lab4_server/internal/services"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

var flagConfig = flag.String("conf", "pc", "path to the config file (ps, vm or note")

func init() {
	config.Init()
}

func main() {
	flag.Parse()
	services.InitializeLogger()

	configFile := "config/" + *flagConfig + ".yaml"
	fmt.Println(configFile)
	postgres.New(configFile)
	//
	//PersonHandler = handlers.NewPersonHandler(postgres.DB)
	//PersonRouteHandler = routes.NewRoutePersonHandler(PersonHandler)
	//
	//gin.SetMode(gin.TestMode)
	//server := gin.Default()
	//server.StaticFile("/favicon.ico", "./favicon.ico")
	//
	//router := server.Group("/api/v1")
	//
	//router.GET("/healthchecker", func(ctx *gin.Context) {
	//	message := "Welcome to Test"
	//	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	//})
	//
	//PersonRouteHandler.PersonRoute(router)
	//log.Fatal(server.Run(":8084"))

	gin.SetMode(config.Server.RunMode)
	routesInit := routes.Init()
}
