package postgres

import (
	"WST_lab4_server/config"
	"WST_lab4_server/internal/models"
	"WST_lab4_server/internal/services"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func New(configFile string) {
	services.InitializeLogger()
	var err error
	///
	configuration, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	services.Logger.Info("Config uploaded successfully.")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		configuration.Database.Host,
		configuration.Database.User,
		configuration.Database.Password,
		configuration.Database.Name,
		configuration.Database.Port,
		configuration.Database.SSLMode)
	/////
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	services.Logger.Info("Database connection established successfully.")
	/////
	err = db.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatalf("error creating table: %v", err)
	}
	services.Logger.Info("Migration completed successfully.")
	/////
	db.Exec("DELETE FROM people")
	///
	result := db.Create(&configuration.Persons)
	if result.Error != nil {
		log.Fatalf("error creating table: %v", result.Error)
	}
	services.Logger.Info("Database updated successfully.")

	//////////////////////////////
	////////////////////////////////
	var results []models.Person
	if err := db.Find(&results).Error; err != nil {
		log.Fatalf("query failed: %v", err)
	}

	for _, record := range results {
		fmt.Println(record)

	}
	fmt.Println("database content in quantity:", len(results))
}
