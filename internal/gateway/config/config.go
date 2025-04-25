package config

import (
	"log"
	"os"
)

type Config struct {
	Port           string
	AccountService string
	TripService    string
	FinanceService string
	MediaService   string
	RealtimeService string
	JWTSecret      string
}

var GatewayConfig *Config

func LoadConfig() {
	GatewayConfig = &Config{
		Port:            getEnv("GATEWAY_PORT", "8080"),
		AccountService:  getEnv("ACCOUNT_SERVICE_URL", "http://account_service:9002"),
		TripService:     getEnv("TRIP_SERVICE_URL", "http://trip_service:9003"),
		FinanceService:  getEnv("FINANCE_SERVICE_URL", "http://finance_service:9004"),
		MediaService:    getEnv("MEDIA_SERVICE_URL", "http://media_service:9005"),
		RealtimeService: getEnv("REALTIME_SERVICE_URL", "http://realtime_service:9007"),
		JWTSecret:       getEnv("JWT_SECRET", "your-secret-key"), // Optional, kalau verif JWT di gateway
	}
	log.Println("✅ Gateway config loaded")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
