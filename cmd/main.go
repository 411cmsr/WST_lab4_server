package main

import (
	"WST_lab4_server/config"
	"WST_lab4_server/internal/database/postgres"
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

}
