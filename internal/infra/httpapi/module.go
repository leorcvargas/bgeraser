package httpapi

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func startServer(*fasthttp.Server) {}

var Module = fx.Module("httpapi",
	fx.Provide(NewServer),
	fx.Invoke(startServer),
)
