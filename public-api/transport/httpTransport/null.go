package httpTransport

import (
	"github.com/sirupsen/logrus"
	"github.com/tnosaj/poc/public-api/transport"
)

type NullSyncTransport struct {
}

func NewNullSyncTransport() transport.SyncTransport {
	return &NullSyncTransport{}
}

func (t *NullSyncTransport) Get(url string) ([]byte, error) {
	logrus.Infof("Null transport GET: %s", url)
	return []byte{}, nil
}
func (t *NullSyncTransport) Post(url string, payload []byte) ([]byte, error) {
	logrus.Infof("Null transport POST: %s", url)
	return []byte{}, nil
}
