package metrics

import "github.com/prometheus/client_golang/prometheus"

// Membuat Counter untuk mengukur jumlah request
var HTTPRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "status"},
)

// Membuat Histogram untuk mengukur latensi request
var HTTPRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Histogram of HTTP request duration in seconds",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"method", "status"},
)

func InitMetrics() {
	// Register metrics ke Prometheus
	prometheus.MustRegister(HTTPRequestsTotal)
	prometheus.MustRegister(HTTPRequestDuration)
}
