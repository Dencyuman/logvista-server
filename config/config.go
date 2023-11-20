package config

import (
	"log"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

func newConfig() *Config {
	cfg := &Config{
		Host:     getEnv("HOST", "localhost"),
		Port:     getEnv("PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Dbname:   getEnv("DB_NAME", "logvista"),
		Sslmode:  getEnv("DB_SSLMODE", "disable"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func logAndExit(format string, v ...interface{}) {
	log.Printf(format, v...)
	os.Exit(1)
}

var AppConfig = newConfig()
