package imageprocesses

import "go.uber.org/fx"

var Module = fx.Module("imageprocesses",
	fx.Provide(NewCreateProcess),
	fx.Provide(NewFindProcess),
)
