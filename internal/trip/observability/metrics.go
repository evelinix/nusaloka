package observability

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Define metrics
var (
	// Counter to track the number of incoming requests to the Trip service
	TripRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "trip_service_requests_total",
			Help: "Total number of requests received by the Trip service.",
		},
		[]string{"method", "route", "status_code"},
	)

	// Histogram to measure the duration of requests to the Trip service
	TripRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "trip_service_duration_seconds",
			Help:    "Histogram of request durations for the Trip service.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route"},
	)

	// Counter to track the number of errors in the Trip service
	TripErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "trip_service_errors_total",
			Help: "Total number of errors in the Trip service.",
		},
		[]string{"method", "route"},
	)
)