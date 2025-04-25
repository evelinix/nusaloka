package observability

import (
	"github.com/prometheus/client_golang/prometheus"
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
