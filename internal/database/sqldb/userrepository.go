package sqldb

import (
	"WST_lab4_server/internal/database"
	"WST_lab4_server/internal/models"
	"database/sql"
)

type UserRepository struct {
	database *Database
}

func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.database.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)

}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}
	if err := r.database.db.QueryRow(
		"SELECT id , email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, database.ErrRecordNotFound
			
		}
		return nil, err
	}
	return u, nil
}
