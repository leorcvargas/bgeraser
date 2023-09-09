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
}

func NewImagesRouter(controller *controllers.ImagesController) *ImagesRouter {
	return &ImagesRouter{
		controller: controller,
	}
}
