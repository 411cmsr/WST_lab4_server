package config

import (
	"WST_lab4_server/internal/models"
	"gopkg.in/yaml.v3"
	"os"
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

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
