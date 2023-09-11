package queues

import (
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"go.uber.org/fx"
)

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
