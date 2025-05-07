package config

import (
	"os"

	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_PORT     string
	POSTGRES_HOST     string
	IMGBB_API_KEY     string
}

func LoadConfig() *Config {
	// load .env into os.Getenv
	if err := godotenv.Load("../../deploy/.env"); err != nil {
		log.Println("no .env file found, reading from environment:", err)
	}
	// Load environment variables from .env file

	cfg := &Config{
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		IMGBB_API_KEY:     os.Getenv("IMGBB_API_KEY"),
	}

	return cfg
}
