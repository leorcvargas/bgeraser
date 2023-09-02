package processinworker

import (
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/fx"
)

var Module = fx.Module("process_in_worker",
	fx.Provide(NewDispatcher),
	fx.Invoke(func(d *ProcessInDispatcher) {
		log.Info("Starting worker")
		go d.Run()
	}),
)
