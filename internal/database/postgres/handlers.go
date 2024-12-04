package postgres

import (
	"WST_lab4_server/internal/models"
	"errors"
	"gorm.io/gorm"
)

type PersonWrapper struct {
	Person models.Person
}

func GetAllPersons() ([]models.Person, error) {
	var persons []models.Person
	err := db.Find(&persons).Error
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func GetPerson(id uint) (*models.Person, error) {
	var person models.Person

	err := db.First(&person, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &person, nil
}

func CheckPersonByID(id int) (bool, error) {
	var person models.Person
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&person).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if person.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (person *PersonWrapper) UpdatePerson() error {

	result := db.Model(person).Updates(models.Person{
		Name:      person.Person.Name,
		Surname:   person.Person.Surname,
		Age:       person.Person.Age,
		Email:     person.Person.Email,
		Telephone: person.Person.Telephone,
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
