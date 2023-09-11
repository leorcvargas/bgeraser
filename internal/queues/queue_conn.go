package queues

import (
	"fmt"

	"github.com/adjust/rmq/v5"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type queueConnection struct {
	config *config.Config
	conn   rmq.Connection
}

func (q *queueConnection) Open() (rmq.Connection, error) {
	if q.conn != nil {
		return q.conn, nil
	}

	return q.connect()
}

func (q *queueConnection) Close() {
	finishCh := q.conn.StopAllConsuming()
	<-finishCh
}

func (q *queueConnection) connect() (rmq.Connection, error) {
	addr := fmt.Sprintf("%s:%s", q.config.Queues.Host, q.config.Queues.Port)

	errCh := make(chan error)

	conn, err := rmq.OpenConnection("bgeraser", "tcp", addr, 1, errCh)
	if err != nil {
		close(errCh)
		return nil, err
	}
	q.conn = conn

	go q.handleAsyncConnectionError(errCh)

	return q.conn, nil
}

func (queueConnection) handleAsyncConnectionError(errCh chan error) {
	for err := range errCh {
		log.Error(err)
	}
}

func NewQueueConnection(config *config.Config) QueueConnection {
	return &queueConnection{config: config}
}
