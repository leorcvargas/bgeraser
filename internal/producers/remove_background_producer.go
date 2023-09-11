package producers

import (
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"github.com/leorcvargas/bgeraser/internal/queues"
)

type removeBackgroundProducer struct {
	queue queues.Queue[*entities.ImageProcess]
}

func (r *removeBackgroundProducer) Send(payload *entities.ImageProcess) error {
	return r.queue.Publish(payload)
}

func NewRemoveBackgroundProducer(
	queue queues.Queue[*entities.ImageProcess],
) Producer[*entities.ImageProcess] {
	return &removeBackgroundProducer{queue}
}
