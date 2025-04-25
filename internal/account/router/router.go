package router

import (
	"net/http"
	"time"

	"github.com/evelinix/nusaloka/internal/account/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func SetupAccountRouter(r chi.Router) {

	r.Use(middleware.RealIP)
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Account Service is up 🚀"))
	})

	r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Auth Endpoint"))
	})

	r.Post("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login Endpoint"))
	})

	r.Post("/auth/register", handler.RegisterHandler)

	r.Post("/auth/forgot-password", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Forgot Password Endpoint"))
	})

	r.Get("/account", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Account Endpoint"))
	})
	

	r.Post("/account/update-password", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Update Password Endpoint"))
	})

	r.Post("/account/update-avatar", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Update Avatar Endpoint"))
	})

	r.Get("/referal", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Referal Endpoint"))
	})

	r.Get("/webauthn", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WebAuthn Endpoint"))
	})

	r.Get("/webauthn/start-registration", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WebAuthn Start Registration Endpoint"))
	})

	r.Get("/webauthn/finish-registration", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WebAuthn Finish Registration Endpoint"))
	})

	// Metrics endpoint
	r.Handle("/metrics", http.HandlerFunc(handler.PrometheusHandler))

}
