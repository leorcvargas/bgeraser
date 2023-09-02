package processoutworker

import (
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/fx"
)

var Module = fx.Module("process_out_worker",
	fx.Provide(NewDispatcher),
	fx.Invoke(func(d *ProcessOutDispatcher) {
		log.Info("Starting worker")
		go d.Run()
	}),
)
