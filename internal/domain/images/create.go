package images

import (
	"fmt"

	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type Create struct {
	config *config.Config
}

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
	url := fmt.Sprintf("%s/%s", c.config.Storage.BucketURL, image.Filename())
	image.SetURL(url)

	return image
}

func NewCreate(config *config.Config) *Create {
	return &Create{config}
}
