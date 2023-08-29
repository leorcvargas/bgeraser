package controllers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

var (
	ErrInvalidFileType = errors.New("invalid file type")
	ErrFileTooLarge    = errors.New("file too large")
)

type ImagesController struct {
	create *images.Create
	save   *images.Save
	config *config.Config
}

func (i *ImagesController) Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	files := form.File["images"]

	if len(files) == 0 {
		return c.Status(http.StatusBadRequest).JSON(ErrResponse{Message: "Missing files"})
	}

	result := make([]entities.Image, 0)

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

		localPath := fmt.Sprintf("%s/%s", i.config.Storage.LocalPath, image.Filename())
		if err := c.SaveFile(file, localPath); err != nil {
			return err
		}

		if err := i.save.Exec(image); err != nil {
			return err
		}

		result = append(result, *image)
	}

	return c.Status(http.StatusOK).JSON(Response{Data: result})
}

func (i *ImagesController) validateUpload(formFile *multipart.FileHeader) error {
	contentType := formFile.Header.Get("Content-Type")
	size := formFile.Size

	log.Infof("Validating file upload with content type %s and size %d", contentType, size)

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
) *ImagesController {
	return &ImagesController{
		config: config,
		create: create,
		save:   save,
	}
}
