package observability

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Define RequestCount metric globally
var RequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gateway_requests_total",
		Help: "Total number of requests processed by the Gateway",
	},
	[]string{"method", "route", "status"},
)

