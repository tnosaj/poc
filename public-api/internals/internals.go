package internals

import "github.com/prometheus/client_golang/prometheus"

type Settings struct {
	Debug bool

	Port string

	AsyncTransportSettings AsyncTransportSettings
	SyncTransportSettings  SyncTransportSettings

	Timeout int

	Backends map[string]map[string]string
}

// Metrics contsins all metric types
type Metrics struct {
	RequestDuration *prometheus.HistogramVec
	ErrorRequests   *prometheus.CounterVec
}

type AsyncTransportSettings struct {
	Name string
}

type SyncTransportSettings struct {
	Name    string
	Timeout int
}
