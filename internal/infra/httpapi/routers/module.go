package routers

import "go.uber.org/fx"

func asRouter(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Router)),
		fx.ResultTags(`group:"routers"`),
	)
}

var Module = fx.Module("routers",
	fx.Provide(
		asRouter(NewPingRouter),
		asRouter(NewImagesRouter),
		asRouter(NewImageProcessesRouter),
		fx.Annotate(
			MakeRouter,
			fx.ParamTags(`group:"routers"`),
		),
	),
)
