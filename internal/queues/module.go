package queues

import (
	"sync"

	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"go.uber.org/fx"
)

var once sync.Once

var Module = fx.Module("queues",
	fx.Provide(
		fx.Annotate(
			NewRemoveBackgroundQueue,
			fx.As(new(Queue[*entities.ImageProcess])),
		),
	),
	fx.Provide(
		fx.Annotate(
			NewQueueConnection,
			fx.As(new(QueueConnection)),
		),
	),
	fx.Provide(NewRemoveBackgroundQueueConsumer),
)
