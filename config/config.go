package config

import (
	"WST_lab4_server/internal/models"
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

// HTTPServerConfig contains the configuration for starting the server
type HTTPServerConfig struct {
	RunMode      string        `yaml:"runMode"`
	BindAddr     string        `yaml:"bindAddr"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

// DatabaseConfig contains the configuration for connecting to the database
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslMode"`
}

var (
	config               Config
	GeneralServerSetting = &GeneralServerConfig{}
	HTTPServerSetting    = &HTTPServerConfig{}
	DatabaseSetting      = &DatabaseConfig{}
)

// Init initializes the server configuration
func Init() {
	file, err := os.Open("config/note.yaml")
	if err != nil {
		log.Fatal("error opening file config", zap.Error(err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("error closing file config", zap.Error(err))
		}
	}(file)
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal("error decoding file config", zap.Error(err))
	}
	*GeneralServerSetting = config.GeneralServer
	*HTTPServerSetting = config.HTTPServer
	*DatabaseSetting = config.Database

}
