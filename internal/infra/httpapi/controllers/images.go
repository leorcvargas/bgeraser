package controllers

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

var (
	ErrInvalidFileType = errors.New("invalid file type")
	ErrFileTooLarge    = errors.New("file too large")
)

type ImagesController struct {
	create  *images.Create
	save    *images.Save
	config  *config.Config
	storage fiber.Storage
}

func (i *ImagesController) Create(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	files := form.File["images"]

	if len(files) == 0 {
		return c.Status(http.StatusBadRequest).
			JSON(ErrResponse{Message: "Missing files"})
	}

	result := make([]string, 0)

	for _, file := range files {
		if err = i.validateUpload(file); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(ErrResponse{
				Message: err.Error(),
			})
		}

		input := images.CreateInput{
			Filename: file.Filename,
			Format:   file.Header["Content-Type"][0],
			Size:     file.Size,
		}
		image := i.create.Exec(input)

		saveFileErr := c.SaveFileToStorage(file, image.Filename(), i.storage)
		if saveFileErr != nil {
			log.Errorw("error saving file to disk: %w", saveFileErr)
			return c.Status(http.StatusInternalServerError).
				JSON(InternalServerErrResponse)
		}

		saveErr := i.save.Exec(image)
		if saveErr != nil {
			log.Errorw("error saving file info: %w", saveErr)
			return c.Status(http.StatusInternalServerError).
				JSON(InternalServerErrResponse)
		}

		result = append(result, image.ID.String())
	}

	return c.Status(http.StatusOK).JSON(Response{Data: result})
}

func (i *ImagesController) validateUpload(
	formFile *multipart.FileHeader,
) error {
	contentType := formFile.Header.Get("Content-Type")
	size := formFile.Size

	log.Infof(
		"Validating file upload with content type %s and size %d",
		contentType,
		size,
	)

	if contentType != "image/png" && contentType != "image/jpeg" {
		return ErrInvalidFileType
	}

	if size > 5*1024*1024 {
		return ErrFileTooLarge
	}

	return nil
}

func NewImagesController(
	config *config.Config,
	create *images.Create,
	save *images.Save,
	storage fiber.Storage,
) *ImagesController {
	return &ImagesController{
		config:  config,
		create:  create,
		save:    save,
		storage: storage,
	}
}
