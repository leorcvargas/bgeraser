package queues

import (
	"fmt"
	"time"

	"github.com/adjust/rmq/v5"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
)

const (
	RemoveBackgroundQueueName = "remove_background"
	maxConsumers              = 2
)

type removeBackgroundQueue struct {
	queue    rmq.Queue
	consumer *removeBackgroundQueueConsumer
}

func (r *removeBackgroundQueue) Publish(payload *entities.ImageProcess) error {
	payloadBytes, err := sonic.Marshal(payload)
	if err != nil {
		return err
	}

	return r.queue.PublishBytes(payloadBytes)
}

func (r *removeBackgroundQueue) StartConsumers() error {
	err := r.queue.StartConsuming(1, time.Second)
	if err != nil {
		log.Fatalf("failed to start queue consumption: %v", err)
		return nil
	}

	for i := 0; i < maxConsumers; i++ {
		_, err = r.queue.AddConsumer(
			fmt.Sprintf("%s-%d", RemoveBackgroundQueueName, i),
			r.consumer,
		)

		if err != nil {
			log.Fatalf("failed to attach consumer to queue: %v", err)
			return err
		}
	}

	return nil
}

func (r *removeBackgroundQueue) handleConsumeErr(
	err error,
	delivery rmq.Delivery,
) {
	log.Errorf("consumer error: %v", err)

	if rejectErr := delivery.Reject(); rejectErr != nil {
		log.Errorf("failed to reject delivery: %v", rejectErr)
	}
}

func NewRemoveBackgroundQueue(
	conn QueueConnection,
	consumer *removeBackgroundQueueConsumer,
) *removeBackgroundQueue {
	openConn, err := conn.Open()
	if err != nil {
		log.Fatalf("failed to open queue connection: %v", err)
		return nil
	}

	queue, err := openConn.OpenQueue(RemoveBackgroundQueueName)
	if err != nil {
		log.Fatalf(
			"failed to open queue %s: %v",
			RemoveBackgroundQueueName,
			err,
		)
		return nil
	}

	return &removeBackgroundQueue{queue, consumer}
}
