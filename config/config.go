package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host       string
	Port       string
	User       string
	Password   string
	Dbname     string
	Sslmode    string
	ServerPort string
}

func newConfig() *Config {
	// .env ファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := &Config{
		Host:       getEnv("HOST", "localhost"),
		Port:       getEnv("PORT", "5432"),
		User:       getEnv("DB_USER", "postgres"),
		Password:   getEnv("DB_PASSWORD", "postgres"),
		Dbname:     getEnv("DB_NAME", "logvista"),
		Sslmode:    getEnv("DB_SSLMODE", "disable"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	log.Printf("Config: %+v\n", cfg)

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
