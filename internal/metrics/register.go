package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// RegisterMetrics registers all custom metrics with the Prometheus registry
func RegisterMetrics() {
	prometheus.MustRegister(RequestCounter)
	prometheus.MustRegister(RequestLatency)
}
