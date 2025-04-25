package router

import (
	"net/http"

	"github.com/evelinix/nusaloka/internal/gateway/handler"
	"github.com/evelinix/nusaloka/internal/gateway/middleware"
	"github.com/go-chi/chi/v5"
)

// SetupGatewayRoute mendefinisikan routes untuk layanan gateway
func SetupGatewayRoute(r chi.Router) {

	// Endpoint untuk Prometheus metrics
	r.Handle("/metrics", http.HandlerFunc(handler.PrometheusHandler))

	// Setup route untuk Auth (login,register)
	r.Route("/api/auth", func(r chi.Router) {
		r.Handle("/auth/*", handler.NewReverseProxy("http://account_service:9001")) // Proxy ke Account Service
	})
	// Setup route untuk Account Service (me,update-password,update-avatar,update-profile)
	r.Route("/api/account", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware) // Middleware autentikasi JWT
		r.Handle("/account/*", handler.NewReverseProxy("http://account_service:9001")) // Proxy ke Account Service
	})

	// Setup route untuk WebAuthn Service (WebAutn start-register,finish-register,start-login,finish-login)
	r.Route("/api/webauthn", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware) // Middleware autentikasi JWT
		r.Handle("/webauthn/*", handler.NewReverseProxy("http://account_service:9001")) // Proxy ke Account Service
	})

	// Setup route untuk Trip Service (trip, itinerary, matching, booking, hotel, ticket, restorant, review )
	r.Route("/api/trip", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware) // Middleware autentikasi JWT
		r.Handle("/*", handler.NewReverseProxy("http://trip_service:9003")) // Proxy ke Trip Service
	})

	// Tambahkan route lainnya sesuai kebutuhan
}
