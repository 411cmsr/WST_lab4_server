package main

import (
	"WST_lab4_server/internal/services"
	"flag"
	"github.com/BurntSushi/toml"

	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "config/config.toml", "config path")
}

func main() {
	flag.Parse()

	config := services.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := services.Start(config); err != nil {
		log.Fatal(err)
	}
}
