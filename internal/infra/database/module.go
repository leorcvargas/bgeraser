package database

import (
	"github.com/leorcvargas/bgeraser/internal/infra/database/imagesdb"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("database",
	fx.Provide(NewPostgresDatabase),
	fx.Options(imagesdb.Module),
	fx.Invoke(func(_ *gorm.DB) {}),
)
