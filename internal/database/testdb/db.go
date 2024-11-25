package testdb

import (
	"WST_lab4_server/internal/database"
	"WST_lab4_server/internal/models"
)

type Database struct {
	userRepository *UserRepository
}

func New() *Database {
	return &Database{}
}

func (s *Database) User() database.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		database: s,
		users:    make(map[int]*models.User),
	}
	return s.userRepository
}
