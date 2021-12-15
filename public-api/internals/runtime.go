package internals

import (
	"github.com/tnosaj/poc/public-api/transport"
	"github.com/tnosaj/poc/public-api/transport/http"
	"github.com/tnosaj/poc/public-api/transport/queues"
)

type Runtime struct {
	Metrics Metrics
	Async   transport.AsyncTransport
	Sync    transport.SyncTransport
}

func NewRuntime(asyncTransport, syncTransport string, metrics Metrics) Runtime {

	return Runtime{Metrics: metrics,
		Async: getAsyncTransport(asyncTransport),
		Sync:  getSyncTransport(syncTransport),
	}
}

func getAsyncTransport(s string) transport.AsyncTransport {
	switch s {
	case "nullqueue":
		return queues.NewNullAsyncTransport()
	}
	return queues.NewNullAsyncTransport()
}

func getSyncTransport(s string) transport.SyncTransport {
	switch s {
	case "nullhttp":
		return http.NewNullSyncTransport()
	}
	return http.NewNullSyncTransport()
}
