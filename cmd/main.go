package main

import (
	"WST_lab4_server/internal/database/postgres"
	"WST_lab4_server/internal/services"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

var flagConfig = flag.String("conf", "pc", "path to the config file (ps, vm or note")

func main() {
	flag.Parse()
	services.InitializeLogger()

	configFile := "config/" + *flagConfig + ".yaml"
	fmt.Println(configFile)
	postgres.New(configFile)

	router := gin.Default()
	router.GET("/api/v1", func(c *gin.Context) {})

}
