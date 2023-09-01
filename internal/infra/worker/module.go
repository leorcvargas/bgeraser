package worker

import (
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/fx"
)

var Module = fx.Module("worker",
	fx.Provide(NewDispatcher),
	fx.Invoke(func(d *Dispatcher) {
		log.Info("Starting worker")
		go d.Run()
	}),
)
