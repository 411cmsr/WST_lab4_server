package database

type Database interface {
	User() UserRepository
}
