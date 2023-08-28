package controllers

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
)

var (
	ErrInvalidFileType = errors.New("Invalid file type")
	ErrFileTooLarge    = errors.New("File too large")
)

type ImagesController struct {
	uploader *images.Uploader
}

func (i *ImagesController) Upload(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("images")
	if err != nil {
		log.Errorw("Error reading file from request:", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	if err = i.validateUpload(fileHeader); err != nil {
		log.Errorw("Error validating file:", err)
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(ErrResponse{Message: err.Error()})
	}

	image, err := i.uploader.Upload(fileHeader)
	if err != nil {
		log.Errorw("Error uploading file:", err)
		return c.
			Status(http.StatusInternalServerError).
			JSON(InternalServerErrResponse)
	}

	return c.
		Status(http.StatusOK).
		JSON(Response{Data: image})
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
	uploader *images.Uploader,
) *ImagesController {
	return &ImagesController{
		uploader: uploader,
	}
}
