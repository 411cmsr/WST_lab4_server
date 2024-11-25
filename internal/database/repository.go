package database

import "WST_lab4_server/internal/models"

type UserRepository interface {
	Create(*models.User) error
	Find(int) (*models.User, error)
	FindByEmail(string) (*models.User, error)
}
