package main

import (
	"WST_lab4_server/config"
	"WST_lab4_server/internal/database/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	//"WST_lab4_server/internal/handlers"
	"WST_lab4_server/internal/httpserver/routes"
	//"WST_lab4_server/internal/services"
	//"flag"
	//"fmt"
	//"github.com/gin-gonic/gin"
)

func init() {
	config.Init()
}

func main() {

	//	services.InitializeLogger()
	postgres.Init()
	gin.SetMode(config.HTTPServerSetting.RunMode)
	routersInit := routes.Init()

	readTimeout := config.HTTPServerSetting.ReadTimeout
	writeTimeout := config.HTTPServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.HTTPServerSetting.BindAddr)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
