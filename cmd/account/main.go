package main

import (
	"log"
	"net/http"

	"github.com/evelinix/nusaloka/internal/account/config"
	"github.com/evelinix/nusaloka/internal/account/observability"
	"github.com/evelinix/nusaloka/internal/account/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	config.LoadConfig()

	// Inisialisasi Prometheus metrics
	observability.InitMetrics()
	
	port := config.AccountConfig.Port
	if port == "" {
		port = "9001"
		log.Printf("✅ Setting Account Service with default port : %s ", port)
	}

	log.Printf("🚀 Starting Account Service on port : %s ", port)

	r := chi.NewRouter()

	router.SetupAccountRouter(r)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("❌ Failed to Account Service crashed: %v", err)
	}

}