package metrics

import (
	"runtime"

	"github.com/google/wire"
)

func initializeGlobalLabels() {
	// SetGlobalLabel("environment", os.Getenv("ENVIRONMENT"))
	// SetGlobalLabel("region", os.Getenv("REGION"))
	SetGlobalLabel("service_name", "whereami")
	SetGlobalLabel("version", "v1.0.0")
	SetGlobalLabel("runtime", runtime.Version())
}

func ProvideMetrics() Metrics {
	// Initialize global labels
	initializeGlobalLabels()

	// Return the metrics provider
	return NewPrometheusMetrics()
}

var ProviderSet = wire.NewSet(ProvideMetrics)
