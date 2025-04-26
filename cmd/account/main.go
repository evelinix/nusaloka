package main

import (
	"runtime"

	"github.com/evelinix/nusaloka/internal/account/config"
	"github.com/evelinix/nusaloka/internal/account/observability"
	"github.com/evelinix/nusaloka/internal/account/router"
	"github.com/evelinix/nusaloka/internal/shared/jwtutil"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	// Inisialisasi Gin router
	r := gin.Default()
	// Muat konfigurasi dari file
	config.LoadConfig()

	ConfigRuntime()

	// Inisialisasi koneksi ke database
	config.InitDatabase()
	config.AutoMigrateDatabase()
	config.InitRedis()

	// Inisialisasi kunci JWT (private & public) untuk autentikasi
	if err := jwtutil.InitKeys(); err != nil {
		log.Fatal().Err(err).Msg("[main] Failed to initialize JWT keys")
	}

	// Menentukan port untuk Account Service, jika tidak ada maka menggunakan default 9001
	port := config.AccountConfig.AccountPort
	if port == "" {
		port = "9001"
		log.Info().Str("port", port).Msg("[main] Setting Account Service with default port")
	}

	// Log informasi sebelum memulai layanan
	log.Info().Str("port", port).Msg("[main] Starting Account Service on port")

	// Inisialisasi Prometheus untuk observabilitas
	observability.InitMetrics()
	
	// Daftarkan middleware Prometheus untuk memonitor request
	// r.Use(sharedMiddleware.RateLimitMiddleware(config.RedisClient, 10, 1*time.Minute))
	r.Use(observability.PrometheusHandler())

	// Setup route untuk Account Service
	router.SetupAccountRouter(r, config.PostgresDB)
	router.SetupWebAuthnRoutes(r, config.PostgresDB)

	// Menjalankan server pada port yang sudah ditentukan
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Error().Err(err).Msg("[main] Failed to start Account Service")
	}
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	log.Info().Int("CPU", nuCPU).Msg("[main] Running with") // .Printf("Running with %d CPUs\n", nuCPU)
}