package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leorcvargas/bgeraser/internal/infra/httpapi/controllers"
)

type ImageProcessesRouter struct {
	controller *controllers.ImageProcessesController
}

func (i *ImageProcessesRouter) Load(r *fiber.App) {
	r.Post("/images/:id/process/:kind", i.controller.CreateProcess)
	r.Get("/images/process/:id", i.controller.GetProcess)
}

func NewImageProcessesRouter(
	controller *controllers.ImageProcessesController,
) *ImageProcessesRouter {
	return &ImageProcessesRouter{
		controller: controller,
	}
}
