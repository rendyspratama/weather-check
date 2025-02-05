package config

import (
	"github.com/joho/godotenv"
	
	"log"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}
}
