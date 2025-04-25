package main

import (
	"log"
	"net/http"
	"os"

	"github.com/evelinix/nusaloka/internal/trip/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	router.SetupTripRouter(r)
	
	port := os.Getenv("TRIP_SERVICE_PORT")
	if port == "" {
		port = "9002" // default fallback biar gak crash
	}

	log.Printf("🚀 Starting Account Service on port : %s ", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("❌ Failed to Account Service crashed: %v", err)
	}
}