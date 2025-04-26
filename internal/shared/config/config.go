package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)


type Config struct {
	GatewayPort string
	AccountPort string
	TripPort string
	FinancePort string
	AdminPort string
	MediaPort string
	RealTimePort string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	RedisHost string
	RedisPort string
	RedisPass string
	JwtSecret string
	JwtPublicPath string
	JwtPrivatePath string
}

var AppConfig *Config

func LoadConfig() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{
        Out:        os.Stdout,
        TimeFormat: "Monday, 02/01/2006 15:04:05 MST",
    }).With().Timestamp().Logger()

	log.Info().Msg("[global] Loading Global config")

	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("[global] Error loading .env file")
	}

	AppConfig = &Config{
		GatewayPort: getEnv("GATEWAY_PORT", "9000"),
		AccountPort: getEnv("ACCOUNT_PORT", "9001"),
		TripPort: getEnv("TRIP_PORT", "9002"),
		FinancePort: getEnv("FINANCE_PORT", "9003"),
		AdminPort: getEnv("ADMIN_PORT", "9004"),
		MediaPort: getEnv("MEDIA_PORT", "9005"),
		RealTimePort: getEnv("REALTIME_PORT", "9006"),
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
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}