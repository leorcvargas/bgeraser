package images

import (
	"github.com/leorcvargas/bgeraser/internal/domain/repositories"
	"go.uber.org/fx"
)

var Module = fx.Module("imagesdb",
	fx.Provide(
		fx.Annotate(
			NewPostgresImageRepository,
			fx.As(new(repositories.ImageRepository)),
			fx.As(new(repositories.ImageProcessRepository)),
		),
	),
)
