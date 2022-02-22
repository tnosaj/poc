package httpTransport

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tnosaj/poc/public-api/internals"
	"github.com/tnosaj/poc/public-api/transport"
)

type HttpSyncTransport struct {
	client *http.Client
}

func NewHTTPSyncTransport(settings internals.SyncTransportSettings) transport.SyncTransport {
	return &HttpSyncTransport{
		client: &http.Client{Timeout: time.Duration(settings.Timeout)},
	}
}

func (t *HttpSyncTransport) Get(url string) ([]byte, error) {
	logrus.Debugf("Http transport Send")
	return []byte{}, nil
}
func (t *HttpSyncTransport) Post(url string, payload []byte) ([]byte, error) {
	logrus.Debugf("Http transport Listen")
	return []byte{}, nil
}
