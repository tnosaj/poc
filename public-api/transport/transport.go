package transport

type AsyncTransport interface {
	Send(string, []byte) error
	Listen() error
}

type SyncTransport interface {
	Get(string) ([]byte, error)
	Post(string, []byte) ([]byte, error)
}
