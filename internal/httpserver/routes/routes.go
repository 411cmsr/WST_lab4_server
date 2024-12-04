package routes

import (
	"WST_lab4_server/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	apiv1 := router.Group("/api/v1")
	apiv1.GET("/persons", handlers.FindPerson)
	apiv1.POST("/persons", handlers.AddPerson)
	apiv1.GET("/persons/list", handlers.GetAllPersons)
	apiv1.GET("/person/:id", handlers.GetPerson)
	apiv1.PUT("/person/:id", handlers.UpdatePerson)
	apiv1.DELETE("/person/:id", handlers.DeletePerson)
	return router
}
