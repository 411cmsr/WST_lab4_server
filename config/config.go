package config

import (
	"WST_lab4_server/internal/models"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Env      string `yaml:"env" env-required:"true"`
	LogLevel string `yaml:"log_level" env-default:"debug"`
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Port     int    `yaml:"port"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
	Httpserver struct {
		BindAddr string `yaml:"bind_addr"`
	} `yaml:"httpserver"`
	Persons []models.Person `yaml:"persons"`
}

// /Server
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

// Database
type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}

func Init() {
	file, err := os.Open("config")
	if err != nil {
		//return nil, err
		fmt.Println("open file error:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		//return nil, err
		fmt.Println("Decode error:", err)
	}
	//return &config, nil
}
