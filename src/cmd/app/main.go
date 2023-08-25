package main

import (
	"log"
	"users-segments-service/config"
	"users-segments-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config setup error: %s", err)
	}

	app.Run(cfg)
}
