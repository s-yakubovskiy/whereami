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

func (m *prometheusMetrics) IncrementCounter(name string, labels map[string]string) error {
	m.mu.RLock()
	counter, exists := m.counters[name]
	m.mu.RUnlock()
	if !exists {
		return errors.New("counter not registered")
	}
	counter.With(labels).Inc()
	return nil
}

func (m *prometheusMetrics) ObserveLatency(name string, duration time.Duration, labels map[string]string) error {
	m.mu.RLock()
	histogram, exists := m.histograms[name]
	m.mu.RUnlock()
	if !exists {
		return errors.New("histogram not registered")
	}
	histogram.With(labels).Observe(duration.Seconds())
	return nil
}

func (m *prometheusMetrics) RegisterCounter(name, help string, labelKeys []string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.counters[name]; exists {
		return errors.New("counter already registered")
	}
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      name,
			Help:      help,
		},
		labelKeys,
	)
	m.counters[name] = counter
	prometheus.MustRegister(counter)
	return nil
}

func (m *prometheusMetrics) RegisterHistogram(name, help string, labelKeys []string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.histograms[name]; exists {
		return errors.New("histogram already registered")
	}
	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: Namespace,
			Name:      name,
			Help:      help,
			Buckets:   prometheus.DefBuckets,
		},
		labelKeys,
	)
	m.histograms[name] = histogram
	prometheus.MustRegister(histogram)
	return nil
}
