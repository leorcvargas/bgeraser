package images

import (
	"mime/multipart"

	"github.com/leorcvargas/bgeraser/internal/domain/entities"
)

type Uploader struct {
	storage    Storage
	repository Repository
}

func (u *Uploader) Upload(fileHeader *multipart.FileHeader) (*entities.Image, error) {
	image := entities.CreateImage(
		fileHeader.Filename,
		fileHeader.Header.Get("Content-Type"),
		fileHeader.Size,
	)

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fileBytes []byte
	_, err = file.Read([]byte(fileBytes))
	if err != nil {
		return nil, err
	}

	if err = u.storage.Write(image.Filename(), fileBytes); err != nil {
		return nil, err
	}

	return image, nil
}

func NewUploader(
	storage Storage,
	repository Repository,
) *Uploader {
	return &Uploader{
		storage:    storage,
		repository: repository,
	}
}
