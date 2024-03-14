package main

import (
	"log"
	"simple-proxy-images/internal/app"
)

func main() {
	log.Println("Init app")

	application := app.NewApp(":3000")

	application.Run()
}
