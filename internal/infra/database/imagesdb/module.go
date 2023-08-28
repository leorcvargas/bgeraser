package imagesdb

import (
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"go.uber.org/fx"
)

var Module = fx.Module("imagesdb",
	fx.Provide(
		fx.Annotate(
			NewPostgresImageRepository,
			fx.As(new(images.Repository)),
		),
	),
)
