package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	return &Config{
		AppName: os.Getenv("APP_NAME"),
		AppPort: os.Getenv("APP_PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),

		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
