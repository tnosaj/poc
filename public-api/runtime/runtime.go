package runtime

import (
	"github.com/tnosaj/poc/public-api/internals"
	"github.com/tnosaj/poc/public-api/transport"
	"github.com/tnosaj/poc/public-api/transport/httpTransport"
	"github.com/tnosaj/poc/public-api/transport/queueTransport"
)

type RuntimeSettings struct {
	Metrics internals.Metrics
	Async   transport.AsyncTransport
	Sync    transport.SyncTransport
}

func NewRuntime(asyncTransport internals.AsyncTransportSettings, syncTransport internals.SyncTransportSettings, metrics internals.Metrics) RuntimeSettings {

	return RuntimeSettings{Metrics: metrics,
		Async: getAsyncTransport(asyncTransport),
		Sync:  getSyncTransport(syncTransport),
	}
}

func getAsyncTransport(s internals.AsyncTransportSettings) transport.AsyncTransport {
	switch s.Name {
	case "nullqueue":
		return queueTransport.NewNullAsyncTransport()
	}
	return queueTransport.NewNullAsyncTransport()
}

func getSyncTransport(s internals.SyncTransportSettings) transport.SyncTransport {
	switch s.Name {
	case "http":
		return httpTransport.NewHTTPSyncTransport(s)
	}
	return httpTransport.NewNullSyncTransport()
}
