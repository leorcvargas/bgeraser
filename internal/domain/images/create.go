package images

import (
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
)

type Create struct{}

type CreateInput struct {
	Filename string
	Format   string
	Size     int64
}

func (c *Create) Exec(input CreateInput) *entities.Image {
	image := entities.CreateImage(
		input.Filename,
		input.Format,
		input.Size,
	)

	return image
}

func NewCreate() *Create {
	return &Create{}
}
