package images

import "go.uber.org/fx"

var Module = fx.Module("images",
	fx.Provide(NewCreate),
	fx.Provide(NewSave),
)
