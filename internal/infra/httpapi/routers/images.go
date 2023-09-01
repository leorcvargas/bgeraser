package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leorcvargas/bgeraser/internal/infra/httpapi/controllers"
)

type ImagesRouter struct {
	controller *controllers.ImagesController
}

func (i *ImagesRouter) Load(r *fiber.App) {
	r.Post("/images", i.controller.Create)
	r.Post("/images/:id/process/:kind", i.controller.CreateProcess)
	r.Get("/images/process/:id", i.controller.GetProcess)
}

func NewImagesRouter(controller *controllers.ImagesController) *ImagesRouter {
	return &ImagesRouter{
		controller: controller,
	}
}
