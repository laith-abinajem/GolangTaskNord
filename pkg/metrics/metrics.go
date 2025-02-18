package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// API request duration histogram
var APIRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "api_request_duration_seconds",
		Help:    "Histogram of response time for API requests",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"route", "method"},
)

// Transactions processed per second
var TransactionsProcessed = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "transactions_processed_total",
		Help: "Total number of transactions processed",
	},
	[]string{"status"},
)

// Cache hit/miss rates
var CacheHits = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "cache_hits_total",
		Help: "Total number of cache hits",
	},
)
var CacheMisses = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "cache_misses_total",
		Help: "Total number of cache misses",
	},
)

// Register all metrics
func RegisterMetrics() {
	prometheus.MustRegister(APIRequestDuration, TransactionsProcessed, CacheHits, CacheMisses)
}
