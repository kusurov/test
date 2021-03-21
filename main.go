package main

import (
	"awesomeProject2/pkg"
	"flag"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configHandler-path", "configs/config.yml", "Путь для файла конфигурации")
}

func main() {
	flag.Parse()

	if err := src.Start(configPath); err != nil {
		log.Fatal(err)
	}
}