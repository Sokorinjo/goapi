package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser   string
	DBPass   string
	DBHost   string
	DBPort   string
	DBName   string
	PORT  int
}

var AppConfig *Config

// This runs automatically BEFORE main()
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env not found, using system environment variables")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	AppConfig = &Config{
		DBUser:   os.Getenv("DB_USER"),
		DBPass:   os.Getenv("DB_PASS"),
		DBHost:   os.Getenv("DB_HOST"),
		DBPort:   os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		PORT:  port,
	}
}