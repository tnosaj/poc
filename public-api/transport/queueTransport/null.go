package queueTransport

import (
	"github.com/sirupsen/logrus"
	"github.com/tnosaj/poc/public-api/transport"
)

type NullAsyncTransport struct {
}

func NewNullAsyncTransport() transport.AsyncTransport {
	return &NullAsyncTransport{}
}

func (t *NullAsyncTransport) Send(topic string, payload []byte) error {
	logrus.Infof("Null transport Send")
	return nil
}
func (t *NullAsyncTransport) Listen() error {
	logrus.Infof("Null transport Listen")
	return nil
}
