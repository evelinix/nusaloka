package handler

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Contoh metrik: Hitung jumlah request yang diterima oleh shared service
	sharedRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "shared_requests_total",
			Help: "Total number of requests to the Shared service",
		},
		[]string{"method", "route"},
	)
)

func init() {
	// Daftarkan metrik pada Prometheus
	prometheus.MustRegister(sharedRequestsTotal)
}

// PrometheusHandler akan menangani endpoint /metrics untuk Shared service
func PrometheusHandler(w http.ResponseWriter, r *http.Request) {
	// Tambahkan logika untuk mengupdate metrik jika perlu
	sharedRequestsTotal.WithLabelValues(r.Method, r.URL.Path).Inc()

	// Panggil handler Prometheus untuk melayani data metrik
	promhttp.Handler().ServeHTTP(w, r)
}

func SetupSharedMetrics(r *http.ServeMux) {
	// Menambahkan route untuk metrik
	r.Handle("/metrics", http.HandlerFunc(PrometheusHandler))
	log.Println("Shared Prometheus metrics endpoint initialized.")
}
