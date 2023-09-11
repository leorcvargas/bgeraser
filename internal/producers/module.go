package producers

import (
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"go.uber.org/fx"
)

var Module = fx.Module("producers", fx.Provide(
	fx.Annotate(
		NewRemoveBackgroundProducer,
		fx.As(new(Producer[*entities.ImageProcess])),
	),
))
