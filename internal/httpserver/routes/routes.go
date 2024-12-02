package routes

import (
	"WST_lab4_server/internal/handlers"
	"github.com/gin-gonic/gin"
)

type PersonRouteHandler struct {
	personHandler handlers.PersonHandler
}

func NewRoutePersonHandler(personHandler handlers.PersonHandler) PersonRouteHandler {
	return PersonRouteHandler{personHandler}
}

func (ph *PersonRouteHandler) PersonRoute(rg *gin.RouterGroup) {

	router := rg.Group("persons")
	router.GET("/", ph.personHandler.FindPerson)
	router.POST("/", ph.personHandler.AddPerson)
	router.GET("/list", ph.personHandler.GetAllPersons)
	router.GET("/:personId", ph.personHandler.GetPerson)
	router.PUT("/:personId", ph.personHandler.UpdatePerson)
	router.DELETE("/:personId", ph.personHandler.DeletePerson)
}
