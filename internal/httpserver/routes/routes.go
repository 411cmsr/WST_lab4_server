package routes

import (
	"WST_lab4_server/internal/database/postgres"
	"github.com/gin-gonic/gin"
)

func Init(httpserver *gin.Engine) {
	//router := gin.New()
	httpserver.Use(gin.Recovery())
	httpserver.Use(gin.Logger())

	apiv1 := httpserver.Group("/api/v1")
	apiv1.GET("/persons", postgres.FindPerson)
	apiv1.POST("/persons", postgres.AddPerson)
	apiv1.GET("/persons/list", postgres.GetAllPersons)
	apiv1.GET("/person/:id", postgres.GetPerson)
	apiv1.PUT("/person/:id", postgres.UpdatePerson)
	apiv1.DELETE("/person/:id", postgres.DeletePerson)

}
