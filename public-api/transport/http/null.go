package http

import (
	"github.com/sirupsen/logrus"
	"github.com/tnosaj/poc/public-api/transport"
)

type NullSyncTransport struct {
}

func NewNullSyncTransport() transport.SyncTransport {
	return &NullSyncTransport{}
}

func (t *NullSyncTransport) Get() error {
	logrus.Infof("Null transport Send")
	return nil
}
func (t *NullSyncTransport) Post() error {
	logrus.Infof("Null transport Listen")
	return nil
}
