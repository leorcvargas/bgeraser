package producers

type Producer[Payload interface{}] interface {
	Send(payload Payload) error
}
