package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PingController struct{}

func (*PingController) Ping(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Pong!"})
}

func NewPingController() *PingController {
	return &PingController{}
}
