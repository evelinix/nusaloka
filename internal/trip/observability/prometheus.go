package observability

import "github.com/prometheus/client_golang/prometheus"

func InitMetrics() {
	prometheus.MustRegister(TripRequestCount)
	prometheus.MustRegister(TripRequestDuration)
	prometheus.MustRegister(TripErrorCount)
}