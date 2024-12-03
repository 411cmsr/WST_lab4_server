package handlers

import (
	"WST_lab4_server/internal/database/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
	//"WST_lab4_server/internal/database/postgres"
	"WST_lab4_server/internal/models"
)

//var logger *zap.Logger
//
//func init() {
//	var err error
//	logger, err = zap.NewDevelopment()
//	if err != nil {
//		panic(err)
//	}
//	defer func(logger *zap.Logger) {
//		err := logger.Sync()
//		if err != nil {
//
//		}
//	}(logger)
//}

type PersonHandler struct {
	DB *gorm.DB
}

func NewPersonHandler(DB *gorm.DB) PersonHandler {
	return PersonHandler{DB}
}

// FindPerson
func (pc *PersonHandler) FindPerson(ctx *gin.Context) {
	name := ctx.Query("name")
	surname := ctx.Query("surname")
	age := ctx.Query("age")
	email := ctx.Query("email")
	telephone := ctx.Query("telephone")

	var persons []models.Person
	query := pc.DB.Model(&models.Person{})

	var conditions []string
	var args []interface{}

	if name != "" {
		conditions = append(conditions, "name = ?")
		args = append(args, name)
	}
	if surname != "" {
		conditions = append(conditions, "surname = ?")
		args = append(args, surname)
	}
	if age != "" {
		conditions = append(conditions, "age = ?")
		args = append(args, age)
	}
	if email != "" {
		conditions = append(conditions, "email = ?")
		args = append(args, email)
	}
	if telephone != "" {
		conditions = append(conditions, "telephone = ?")
		args = append(args, telephone)
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), args...)
	}

	result := query.Find(&persons)
	if result.Error != nil || len(persons) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No person found with the given criteria"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": persons})
}

// AddPerson
func (ph *PersonHandler) AddPerson(ctx *gin.Context) {
	var payload *models.AddPersonRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newPerson := models.Person{
		Name:      payload.Name,
		Surname:   payload.Surname,
		Age:       payload.Age,
		Email:     payload.Email,
		Telephone: payload.Telephone,
	}

	result := ph.DB.Create(&newPerson)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Person with that email already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPerson})
}

// GetPerson
func (ph *PersonHandler) GetPerson(ctx *gin.Context) {

	personId := ctx.Param("personId")

	fmt.Println(personId)
	fmt.Println("GetPerson")

	var person models.Person
	var results []models.Person
	if err := postgres.DB.Find(&results).Error; err != nil {
		log.Fatalf("query failed: %v", err)
	}

	for _, record := range results {
		fmt.Println(record)

	}

	result := postgres.DB.First(&person, "id = ?", personId)
	fmt.Println("PERSONNNNNNNN", person)
	fmt.Println("ERREREREOROROROORORORO", result.Error)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No person with that id exists"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": person})
}

// GetAllPerson
func (ph *PersonHandler) GetAllPersons(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}
	offset := (intPage - 1) * intLimit

	var persons []models.Person
	result := ph.DB.Limit(intLimit).Offset(offset).Find(&persons)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": persons})
}

// UpdatePerson
func (ph *PersonHandler) UpdatePerson(ctx *gin.Context) {
	personId := ctx.Param("personId")
	var payload *models.UpdatePersonRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedPerson models.Person
	result := ph.DB.First(&updatedPerson, "id = ?", personId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No person with that email exists"})
		return
	}
	personToUpdate := models.Person{
		Name:      payload.Name,
		Surname:   payload.Surname,
		Age:       payload.Age,
		Email:     payload.Email,
		Telephone: payload.Telephone,
	}
	ph.DB.Model(&updatedPerson).Updates(personToUpdate)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPerson})
}

// DeletePerson    ----  check
func (ph *PersonHandler) DeletePerson(ctx *gin.Context) {
	personId := ctx.Param("personId")

	result := ph.DB.Delete(&models.Person{}, "id = ?", personId)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error occurred while deleting the person"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No person with that id exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
