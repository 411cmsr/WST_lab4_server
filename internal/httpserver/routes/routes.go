package routes

import (
	"WST_lab4_server/internal/handlers"
	"github.com/gin-gonic/gin"
)

//type PersonRouteHandler struct {
//	personHandler handlers.PersonHandler
//}
//
//func NewRoutePersonHandler(personHandler handlers.PersonHandler) PersonRouteHandler {
//	return PersonRouteHandler{personHandler}
//}
//
//func (ph *PersonRouteHandler) PersonRoute(rg *gin.RouterGroup) {
//
//	router := rg.Group("persons")
//	router.GET("/", ph.personHandler.FindPerson)
//	router.POST("/", ph.personHandler.AddPerson)
//	router.GET("/list", ph.personHandler.GetAllPersons)
//	router.GET("/:id", ph.personHandler.GetPerson)
//	router.PUT("/:id", ph.personHandler.UpdatePerson)
//	router.DELETE("/:id", ph.personHandler.DeletePerson)
//}

func RegisterRoutes(server *gin.Engine) {

	router := server.Group("persons")
	router.GET("/", handlers.GetPerson)

}
