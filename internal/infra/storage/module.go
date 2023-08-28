package storage

import (
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"go.uber.org/fx"
)

var Module = fx.Module("storage",
	fx.Provide(
		fx.Annotate(
			NewLocalImageStorage,
			fx.As(new(images.Storage)),
		),
	),
)
