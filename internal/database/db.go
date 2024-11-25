package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Database {
	return &Database{
		config: config,
	}
}

func (s *Database) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Database) Close() {
	s.db.Close()
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
