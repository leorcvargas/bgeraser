package images

import "go.uber.org/fx"

var Module = fx.Module("images",
	fx.Provide(NewCreate),
	fx.Provide(NewSave),
	fx.Provide(NewCreateProcess),
	fx.Provide(NewFindProcess),
	fx.Provide(NewJobQueue),
)
