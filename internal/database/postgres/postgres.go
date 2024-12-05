package postgres

import (
	"WST_lab4_server/config"
	//"WST_lab4_server/internal/logging"
	"WST_lab4_server/internal/models"
	//"WST_lab4_server/internal/services"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func Init() {
	//services.InitializeLogger()

	dsn1 := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.DatabaseSetting.Host,
		config.DatabaseSetting.User,
		config.DatabaseSetting.Password,
		config.DatabaseSetting.Name,
		config.DatabaseSetting.Port,
		config.DatabaseSetting.SSLMode)
	fmt.Println("Database", dsn1)

	var err error
	var logLevel logger.LogLevel

	switch config.GeneralServerSetting.LogLevel {
	case "fatal":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info", "debug":
		logLevel = logger.Info
	default:
		logLevel = logger.Info // Значение по умолчанию
	}
	dsn := fmt.Sprintf("host=127.0.0.1 user=pguser password=pgpassword dbname=wstbd port=5432 sslmode=disable")
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
	//	config.DatabaseSetting.Host,
	//	config.DatabaseSetting.User,
	//	config.DatabaseSetting.Password,
	//	config.DatabaseSetting.Name,
	//	config.DatabaseSetting.Port,
	//	config.DatabaseSetting.SSLMode)
	//fmt.Println("Database", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	//logging.Logger.Info("Database connection established successfully.")
	/////
	err = db.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatalf("error creating table: %v", err)
	}
	//logging.Logger.Info("Migration completed successfully.")
	/////
	db.Exec("DELETE FROM people")
	///
	//result := db.Create(&config.GeneralServerSetting.DataSet)
	//if result.Error != nil {
	//	log.Fatalf("error creating table: %v", result.Error)
	//}
	//logging.Logger.Info("Database updated successfully.")

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
