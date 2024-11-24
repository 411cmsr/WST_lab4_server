package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	config *Config
	db     *sql.DB
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

//var db *gorm.DB
//
//func InitDB(configFile string) error {
//	var err error
//	configuration, err := config.LoadConfig(configFile)
//
//	if err != nil {
//		log.Fatalf("error loading config: %v", err)
//	}
//
//	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
//		configuration.Database.Host,
//		configuration.Database.User,
//		configuration.Database.Password,
//		configuration.Database.Name,
//		configuration.Database.Port,
//		configuration.Database.SSLMode)
//
//	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		return err
//	}
//
//	return db.AutoMigrate(&models.Person{})
//}
//
//func UpdateDB(configFile string) error {
//	var err error
//	configuration, err := config.LoadConfig(configFile)
//	if err != nil {
//		log.Fatalf("error loading config: %v", err)
//	}
//	db.Exec("DELETE FROM people")
//	result := db.Create(&configuration.Persons)
//	if result.Error != nil {
//		return result.Error
//	}
//	return err
//}
//
//func GetDB() *gorm.DB {
//	return db
//}
