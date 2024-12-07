package metrics

import (
	"errors"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Namespace for all custom metrics
	Namespace = "whrmi"
)

// Metrics defines the interface for custom metric operations
type Metrics interface {
	IncrementCounter(name string, labels map[string]string) error
	ObserveLatency(name string, duration time.Duration, labels map[string]string) error
	RegisterCounter(name, help string, labelKeys []string) error
	RegisterHistogram(name, help string, labelKeys []string) error
}

type prometheusMetrics struct {
	counters   map[string]*prometheus.CounterVec
	histograms map[string]*prometheus.HistogramVec
	mu         sync.RWMutex
}

// NewPrometheusMetrics initializes a Prometheus-based metrics implementation
func NewPrometheusMetrics() Metrics {
	return &prometheusMetrics{
		counters:   make(map[string]*prometheus.CounterVec),
		histograms: make(map[string]*prometheus.HistogramVec),
	}
}

func (m *prometheusMetrics) mergeLabels(customLabels map[string]string) map[string]string {
	mergedLabels := GetGlobalLabels() // Get global labels
	for k, v := range customLabels {
		mergedLabels[k] = v
	}
	return mergedLabels
}

// mergeGlobalLabels adds global label keys to custom label keys for registration
func mergeGlobalLabels(customLabels []string) []string {
	globalLabels := GetGlobalLabels() // Global labels defined elsewhere
	keys := make([]string, 0, len(globalLabels)+len(customLabels))
	for k := range globalLabels {
		keys = append(keys, k)
	}
	return append(keys, customLabels...)
}

func (m *prometheusMetrics) IncrementCounter(name string, labels map[string]string) error {
	m.mu.RLock()
	counter, exists := m.counters[name]
	m.mu.RUnlock()
	if !exists {
		return errors.New("counter not registered")
	}
	mergedLabels := m.mergeLabels(labels)
	counter.With(mergedLabels).Inc()
	return nil
}

func (m *prometheusMetrics) ObserveLatency(name string, duration time.Duration, labels map[string]string) error {
	m.mu.RLock()
	histogram, exists := m.histograms[name]
	m.mu.RUnlock()
	if !exists {
		return errors.New("histogram not registered")
	}
	mergedLabels := m.mergeLabels(labels)
	histogram.With(mergedLabels).Observe(duration.Seconds())
	return nil
}

// Updated RegisterCounter method
func (m *prometheusMetrics) RegisterCounter(name, help string, labelKeys []string) error {
	mergedKeys := mergeGlobalLabels(labelKeys)
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:      name,
			Namespace: Namespace,
			Help:      help,
		},
		mergedKeys,
	)
	m.counters[name] = counter
	prometheus.MustRegister(counter)
	return nil
}

// Updated RegisterHistogram method
func (m *prometheusMetrics) RegisterHistogram(name, help string, labelKeys []string) error {
	mergedKeys := mergeGlobalLabels(labelKeys)
	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      name,
			Namespace: Namespace,
			Help:      help,
			Buckets:   prometheus.DefBuckets,
		},
		mergedKeys,
	)
	m.histograms[name] = histogram
	prometheus.MustRegister(histogram)
	return nil
}
