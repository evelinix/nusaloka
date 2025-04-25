package config

import (
	"log"
	"os"
)

type Config struct {
Port           string
}

var AccountConfig *Config

func LoadConfig() {
	AccountConfig = &Config{
		Port:            getEnv("ACCOUNT_SERVICE_PORT", "9001"),
	}
	log.Println("✅ Account config loaded")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}