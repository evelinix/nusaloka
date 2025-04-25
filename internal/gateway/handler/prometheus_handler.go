package handler

import (
	"net/http"

	"github.com/evelinix/nusaloka/internal/gateway/observability"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PrometheusHandler serves the Prometheus metrics endpoint
func PrometheusHandler(w http.ResponseWriter, r *http.Request) {
	// Record the metrics if necessary
	observability.RequestCount.WithLabelValues(r.Method, r.URL.Path, "200").Inc() // Example: counting successful requests

	// Serve the metrics endpoint
	promhttp.Handler().ServeHTTP(w, r)
}
