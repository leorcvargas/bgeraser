package queues

import (
	"fmt"

	"github.com/adjust/rmq/v5"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
	"github.com/redis/go-redis/v9"
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
	errCh := make(chan error)

	addr := fmt.Sprintf(
		"%s:%s",
		q.config.Queues.Host,
		q.config.Queues.Port,
	)

	rc := redis.NewClient(&redis.Options{
		Addr:       addr,
		Username:   q.config.Queues.User,
		Password:   q.config.Queues.Password,
		DB:         1,
		ClientName: "bgeraser_client",
	})

	conn, err := rmq.OpenConnectionWithRedisClient(
		"bgeraser",
		rc,
		errCh,
	)
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
