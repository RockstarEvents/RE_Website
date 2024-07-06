package main

import (
	"eventPlanner/internal/app"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	port := parseConfig()
	app.RunApp(port)
}

func parseConfig() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("REST_PORT")
	if port == "" {
		log.Fatal("REST_PORT is not set in the environment")
	}
	return port
}
