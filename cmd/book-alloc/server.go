package main

import (
	"book-alloc/api/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if "" == env {
		env = "local"
	}
	if err := godotenv.Load(".env." + env); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	r := router.Initialize()
	r.Run()
}
