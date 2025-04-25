package handler

import (
	"net/http"

	"github.com/evelinix/nusaloka/internal/account/observability"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PrometheusHandler exposes metrics for Prometheus
func PrometheusHandler(w http.ResponseWriter, r *http.Request) {
	observability.AccountRequestsTotal.WithLabelValues(r.Method, r.URL.Path).Inc()
	promhttp.Handler().ServeHTTP(w, r)
}
