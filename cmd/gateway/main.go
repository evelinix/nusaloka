package main

import (
	"log"
	"net/http"

	"github.com/evelinix/nusaloka/internal/gateway/config"
	"github.com/evelinix/nusaloka/internal/gateway/observability"
	"github.com/evelinix/nusaloka/internal/gateway/router"
	sharedMiddleware "github.com/evelinix/nusaloka/internal/shared/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	config.LoadConfig()

	port := config.GatewayConfig.Port

	if port == "" {
		port = "8080"
	}

	// Setup router dan route
	r := chi.NewRouter()

	// Setup observability metrics
	observability.InitMetrics()

	// Middleware untuk logging, CORS, dll
	r.Use(sharedMiddleware.LoggingMiddleware)

	// Setup routes untuk service reverse proxy
	router.SetupGatewayRoute(r)

	// Jalankan HTTP server
	log.Println("Gateway server started on port", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}