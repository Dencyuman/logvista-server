package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"db_user"`
	Password string `yaml:"db_password"`
	Dbname   string `yaml:"db_name"`
	Sslmode  string `yaml:"db_sslmode"`
}

func newConfig() *Config {
	cfg := &Config{}

	// config.yamlファイルを読み取る
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		logAndExit("config.yamlの読み込みに失敗しました\nError: %v", err)
	}

	// YAMLを構造体にアンマーシャル
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		logAndExit("config.yamlの解析に失敗しました\nError: %v", err)
	}

	return cfg
}

func logAndExit(format string, v ...interface{}) {
	log.Printf(format, v...)
	time.Sleep(5 * time.Second)
	os.Exit(1)
}

var AppConfig = newConfig()
