package main

import (
	"flag"
	"kusurovAPI/internal/app"
	"log"
)

var (
	configPath  string
	loggingPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yml", "Путь для файла конфигурации")
	flag.StringVar(&loggingPath, "logging-path", "logs/", "Путь до папки с логами")
}

func main() {
	flag.Parse()

	if err := app.Start(configPath, loggingPath); err != nil {
		log.Fatal(err)
	}
}
