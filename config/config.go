package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Sslmode  string
}

func newConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Dbname:   os.Getenv("DBNAME"),
		Port:     os.Getenv("PORT"),
		Sslmode:  os.Getenv("SSLMODE"),
	}
}

var AppConfig = newConfig()
