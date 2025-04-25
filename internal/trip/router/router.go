package router

import (
	"fmt"
	"net/http"

	"github.com/evelinix/nusaloka/internal/trip/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupTripRouter(r chi.Router) {
	// Middleware untuk logging
	r.Use(middleware.Logger) // Log setiap request
	r.Use(middleware.Recoverer) // Tangani panic

	
	// Root route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Trip Service API")
	})
	
	// Route untuk metrics
	r.Handle("/metrics", http.HandlerFunc(handler.PrometheusHandler))
}
