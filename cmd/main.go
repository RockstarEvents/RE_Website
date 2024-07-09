package main

import (
	"eventPlanner/internal/app"
	"eventPlanner/internal/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	cfg := config.Config{}

	app.RunApp(cfg)
}
