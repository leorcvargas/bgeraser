package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leorcvargas/bgeraser/internal/infra/httpapi/controllers"
)

type PingRouter struct {
	controller *controllers.PingController
}

func (p *PingRouter) Load(r *fiber.App) {
	r.Get("/ping", p.controller.Ping)
}

func NewPingRouter(controller *controllers.PingController) *PingRouter {
	return &PingRouter{
		controller: controller,
	}
}
