package testdb

import (
	"WST_lab4_server/internal/database"
	"WST_lab4_server/internal/models"
)

type UserRepository struct {
	database *Database
	users    map[string]*models.User
}

func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	r.users[u.Email] = u
	u.ID = len(r.users)
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, database.ErrRecordNotFound
	}
	return u, nil

}
