package testdb

import (
	"WST_lab4_server/internal/database"
	"WST_lab4_server/internal/models"
)

type UserRepository struct {
	database *Database
	users    map[int]*models.User
}

func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}
func (r *UserRepository) Find(id int) (*models.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, database.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, database.ErrRecordNotFound
}
