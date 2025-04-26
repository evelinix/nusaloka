package observability

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


var (
	AccountRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "account_requests_total",
			Help: "Total number of HTTP requests received by the Account service",
		},
		[]string{"method", "route"},
	)

	LoginDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "account_service_login_duration_seconds",
			Help:    "Duration of login requests to Account Service",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"status"},
	)

	LoginSuccess = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "account_login_success_total",
			Help: "Total number of successful login attempts",
		},
		[]string{"method"},
	)

	LoginFailure = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "account_login_failure_total",
			Help: "Total number of failed login attempts",
		},
		[]string{"method"},
	)

	RegistrationDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "account_service_registration_duration_seconds",
			Help:    "Duration of registration requests to Account Service",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"status"},
	)

	RegisterSuccess = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "account_register_success_total",
			Help: "Total number of successful register attempts",
		},
		[]string{"method"},
	)

	RegisterFailure = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "account_register_failure_total",
			Help: "Total number of failed register attempts",
		},
		[]string{"method"},
	)
)

func InitMetrics() {
	// Register metrics
	prometheus.MustRegister(AccountRequestsTotal)
	prometheus.MustRegister(LoginDuration)
	prometheus.MustRegister(LoginSuccess)
	prometheus.MustRegister(LoginFailure)
	prometheus.MustRegister(RegistrationDuration)
	prometheus.MustRegister(RegisterSuccess)
	prometheus.MustRegister(RegisterFailure)
}

func PrometheusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Increment the counter for the request method and path
		AccountRequestsTotal.WithLabelValues(c.Request.Method, c.Request.URL.Path).Inc()

		// Handle the Prometheus metrics request
		if c.Request.URL.Path == "/metrics" {
			promhttp.Handler().ServeHTTP(c.Writer, c.Request)
			return
		}

		c.Next()
	}
}
