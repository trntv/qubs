package broker

type Link interface {
	Send(*Message) error
	Close() error
	IsClosed() bool
}
