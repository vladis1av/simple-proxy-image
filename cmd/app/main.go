package main

import (
	"flag"
	"log"
	"simple-proxy-image/internal/app"
	"simple-proxy-image/internal/config"
	"simple-proxy-image/internal/config/env"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	log.Println("Init app")

	flag.Parse()

	log.Println("Load config")
	err := config.Load(configPath)
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	httpConfig, err := env.NewHTTPConfig()
	if err != nil {
		log.Fatal("failed to get http config: ", err)
	}

	application := app.NewApp(httpConfig.Address())

	application.Run()
}
