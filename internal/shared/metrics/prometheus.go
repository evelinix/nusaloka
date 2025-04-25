package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metrik untuk jumlah total permintaan yang diterima oleh aplikasi
var (
	TotalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_requests_received",
			Help: "Total number of requests received by the application",
		},
		[]string{"service", "status"},
	)

	TotalErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_errors",
			Help: "Total number of errors occurred in the application",
		},
		[]string{"service", "error_type"},
	)
)

func init() {
	// Mendaftarkan metrik ke Prometheus
	prometheus.MustRegister(TotalRequests)
	prometheus.MustRegister(TotalErrors)
}
