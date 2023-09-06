package storage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"go.uber.org/fx"
)

var Module = fx.Module("storage",
	fx.Provide(
		fx.Annotate(
			NewS3Storage,
			fx.As(new(fiber.Storage)),
			fx.As(new(images.Storage)),
		),
	),
)
