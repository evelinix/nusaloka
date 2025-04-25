package observability

import "github.com/prometheus/client_golang/prometheus"

func InitMetrics() {
	// Register metrics with Prometheus
	prometheus.MustRegister(RequestCount)
}
