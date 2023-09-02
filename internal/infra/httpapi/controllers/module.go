package controllers

import "go.uber.org/fx"

var Module = fx.Module("controllers",
	fx.Provide(NewPingController),
	fx.Provide(NewImagesController),
)
