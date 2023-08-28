package images

import (
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
)

type Create struct {
	repository Repository
}

func (c *Create) Exec(filename string, format string, size int64) *entities.Image {
	image := entities.CreateImage(
		filename,
		format,
		size,
	)

	return image
}

func NewCreate(
	repository Repository,
) *Create {
	return &Create{
		repository: repository,
	}
}
