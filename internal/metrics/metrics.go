package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// RequestCounter tracks the number of requests
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "custom_request_count",
			Help: "Total number of requests received by the server",
		},
		[]string{"method", "status"},
	)

	// RequestLatency tracks the latency of requests
	RequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "custom_request_latency_seconds",
			Help:    "Latency of requests handled by the server",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)
