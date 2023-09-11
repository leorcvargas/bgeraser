package queues

import (
	"github.com/adjust/rmq/v5"
)

type Queue[Payload interface{}] interface {
	Publish(payload Payload) error
	StartConsumers() error
}

type QueueConnection interface {
	Open() (rmq.Connection, error)
	Close()
}
