package controllers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	domainerrors "github.com/leorcvargas/bgeraser/internal/domain/errors"
	"github.com/leorcvargas/bgeraser/internal/domain/imageprocesses"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type ImageProcessesController struct {
	config        *config.Config
	createProcess *imageprocesses.CreateProcess
	findProcess   *imageprocesses.FindProcess
	storage       fiber.Storage
}

func (i *ImageProcessesController) CreateProcess(c *fiber.Ctx) error {
	id := c.Params("id")
	kind := c.Params("kind")

	if kind != "REMOVE_BACKGROUND" {
		return c.Status(http.StatusBadRequest).JSON(ErrResponse{
			Message: "Invalid kind",
		})
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrResponse{
			Message: "Invalid id",
		})
	}

	process, err := i.createProcess.Exec(
		parsedID,
		entities.ImageProcessKindRemoveBackground,
	)
	if err != nil {
		if errors.Is(err, domainerrors.ErrImageNotFound) {
			return c.Status(http.StatusNotFound).JSON(ErrResponse{
				Message: err.Error(),
			})
		}

		log.Errorf("Error creating process: %s", err.Error())

		return c.Status(http.StatusInternalServerError).
			JSON(InternalServerErrResponse)
	}

	return c.Status(http.StatusOK).JSON(Response{Data: process.ID})
}

func (i *ImageProcessesController) GetProcess(c *fiber.Ctx) error {
	id := c.Params("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrResponse{
			Message: "Invalid id",
		})
	}

	process, err := i.findProcess.Get(parsedID)
	if err != nil {
		if errors.Is(err, domainerrors.ErrImageProcessNotFound) {
			return c.Status(http.StatusNotFound).JSON(ErrResponse{
				Message: err.Error(),
			})
		}

		log.Errorf("Error getting process: %s", err.Error())

		return c.Status(http.StatusInternalServerError).
			JSON(InternalServerErrResponse)
	}

	return c.Status(http.StatusOK).JSON(Response{Data: process})
}

func NewImageProcessesController(
	config *config.Config,
	create *images.Create,
	save *images.Save,
	createProcess *imageprocesses.CreateProcess,
	findProcess *imageprocesses.FindProcess,
	storage fiber.Storage,
) *ImageProcessesController {
	return &ImageProcessesController{
		config:        config,
		createProcess: createProcess,
		findProcess:   findProcess,
	}
}
