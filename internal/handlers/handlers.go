package handlers

import (
	"WST_lab4_server/internal/database/postgres"
	"WST_lab4_server/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllPersonsHandler(context *gin.Context) {
	persons, err := postgres.GetAllPersons()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch persons. Try again later."})
		return
	}
	context.JSON(http.StatusOK, persons)
}
func GetPersonHandler(context *gin.Context) {
	personId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse person id."})
		return
	}

	person, err := postgres.CheckPersonByID(uint(personId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch person."})
		return
	}
	context.JSON(http.StatusOK, person)
}

func AddPerson(context *gin.Context) {

	var person models.Person
	err := context.ShouldBindJSON(&person)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "person": person})
}
func UpdatePerson(context *gin.Context) {
	personId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	var updatedPerson postgres.PersonWrapper
	err = context.ShouldBindJSON(&updatedPerson)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedPerson.Person.ID = uint(personId)
	err = updatedPerson.UpdatePerson()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update person."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Person updated successfully!"})
}

func DeletePerson(context *gin.Context) {
	personId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse person id."})
		return
	}
	person, err := postgres.GetPerson(uint(personId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the person ."})
		return
	}
	if person == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Person not found."})
		return
	}
	wrapper := &postgres.PersonWrapper{Person: person}

	err = wrapper.DeletePerson()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Delete Person"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
