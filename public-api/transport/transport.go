package transport

type AsyncTransport interface {
	Send() error
	Listen() error
}

type SyncTransport interface {
	Get() error
	Post() error
}
