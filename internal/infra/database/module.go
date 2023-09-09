package database

import (
	"github.com/leorcvargas/bgeraser/internal/infra/database/images"
	"go.uber.org/fx"
)

var Module = fx.Module("database",
	fx.Provide(NewEntClient),
	fx.Options(images.Module),
)
