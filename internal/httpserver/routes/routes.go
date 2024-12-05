package routes

import (
	"WST_lab4_server/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Init(httpserver *gin.Engine) {
	//router := gin.New()
	httpserver.Use(gin.Recovery())
	httpserver.Use(gin.Logger())

	apiv1 := httpserver.Group("/api/v1")
	//apiv1.GET("/persons", handlers.FindPerson)
	apiv1.POST("/persons", handlers.AddPerson)
	apiv1.GET("/persons/list", handlers.GetAllPersonsHandler)
	apiv1.GET("/person/:id", handlers.GetPersonHandler)
	apiv1.PUT("/person/:id", handlers.UpdatePerson)
	apiv1.DELETE("/person/:id", handlers.DeletePerson)

}
