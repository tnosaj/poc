package internals

import "github.com/prometheus/client_golang/prometheus"

type Settings struct {
	Debug bool

	AsyncTransport string
	Port           string
	SyncTransport  string

	Timeout int
}

// Metrics contsins all metric types
type Metrics struct {
	RequestDuration *prometheus.HistogramVec
	ErrorRequests   *prometheus.CounterVec
}
