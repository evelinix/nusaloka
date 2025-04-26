package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	AccountPort string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
	RedisHost   string
	RedisPort   string
	RedisPass   string
	JwtSecret   string
	JwtPublicPath string
	JwtPrivatePath string
	RPName string
	RPID string
	RPOrigin []string
}

var AccountConfig *Config

func LoadConfig() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "Monday, 02/01/2006 15:04:05 MST",
		}).With().Timestamp().Logger()

	log.Info().Msg("[config] Loading Account config")

	err := godotenv.Load()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("[config] Error loading .env file")
	}

	AccountConfig = &Config{
		AccountPort: getEnv("ACCOUNT_SERVICE_PORT", "9001"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", "postgres"),
		DBName: getEnv("DB_NAME", "apps"),
		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
		RedisPass: getEnv("REDIS_PASS", ""),
		JwtSecret: getEnv("JWT_SECRET", "secret"),
		JwtPublicPath: getEnv("JWT_PUBLIC_PATH", "public.pem"),
		JwtPrivatePath: getEnv("JWT_PRIVATE_PATH", "private.pem"),
		RPName: getEnv("RP_NAME", "Nusaloka"),
		RPID: getEnv("RP_ID", "localhost"),
		RPOrigin: getListEnv("RP_ORIGIN", "http://localhost:9001"),
	}

	log.Info().Msg("[config] .env file loaded successfully")

}

func getListEnv(key, fallback string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return strings.Split(fallback, ",")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Info().Str(key, fallback).Msg("[config] Using fallback value")
	return fallback
}