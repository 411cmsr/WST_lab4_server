package sqldb

import (
	//"WST_lab4_server/internal/database"
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

func (s *Database) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		database: s,
	}
	return s.userRepository
}
