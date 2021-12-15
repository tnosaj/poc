package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tnosaj/poc/public-api/internals"
)

// RegisterPrometheusMetrics setsup metrics and returns them
func RegisterPrometheusMetrics() internals.Metrics {

	RequestDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "request_duration_seconds",
		Help:    "Histogram for the runtime of a simple method function.",
		Buckets: prometheus.LinearBuckets(0.00, 0.002, 75),
	}, []string{"method"})

	ErrorReuests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "error_requests",
			Help: "The total number of failed requests",
		},
		[]string{"method"},
	)

	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(ErrorReuests)

	return internals.Metrics{RequestDuration: RequestDuration, ErrorRequests: ErrorReuests}
}
