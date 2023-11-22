package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	Dbname     string
	DbSslmode  string
	ServerPort string
	ViteApiUrl string
}

func newConfig() *Config {
	// .env ファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := &Config{
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbUser:     getEnv("DB_USER", "postgres"),
		DbPassword: getEnv("DB_PASSWORD", "postgres"),
		Dbname:     getEnv("DB_NAME", "logvista"),
		DbSslmode:  getEnv("DB_SSLMODE", "disable"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		ViteApiUrl: getEnv("VITE_API_URL", "http://localhost:8080/api/v1"),
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
