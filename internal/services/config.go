package services

import (
	"WST_lab4_server/internal/database"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	DataBase *database.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8092",
		LogLevel: "debug",
		DataBase: database.NewConfig(),
	}
}
