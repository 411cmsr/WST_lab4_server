package config

import (
	"WST_lab4_server/internal/models"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	GeneralServer GeneralServerConfig `yaml:"generalServer"`
	HTTPServer    HTTPServerConfig    `yaml:"httpServer"`
	Database      DatabaseConfig      `yaml:"database"`
}

type GeneralServerConfig struct {
	Env      string          `yaml:"env" env-required:"true"`
	LogLevel string          `yaml:"logLevel" env-default:"debug"`
	DataSet  []models.Person `yaml:"persons"`
}

var GeneralServerSetting = &GeneralServerConfig{}

// HTTPServerConfig contains the configuration for starting the server
type HTTPServerConfig struct {
	RunMode      string        `yaml:"runMode"`
	BindAddr     string        `yaml:"bindAddr"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

var HTTPServerSetting = &HTTPServerConfig{}

// DatabaseConfig contains the configuration for connecting to the database
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslMode"`
}

var DatabaseSetting = &DatabaseConfig{}

// Init initializes the server configuration
func Init() {
	file, err := os.Open("config/note.yaml")
	fmt.Println("File1111111\n\n")
	if err != nil {
		log.Fatal("Failed to initialize config", zap.Error(err))
	}
	fmt.Println("File12222222\n\n")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	fmt.Println("File33333333\n\n")
	var config Config
	decoder := yaml.NewDecoder(file)
	fmt.Println("File444444\n\n")
	if err := decoder.Decode(&config); err != nil {
		log.Fatal("Decode file config error:", zap.Error(err))
	}
}
