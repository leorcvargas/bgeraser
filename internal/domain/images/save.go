package images

import (
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"github.com/leorcvargas/bgeraser/internal/domain/repositories"
)

type Save struct {
	repository repositories.ImageRepository
}

func (s *Save) Exec(image *entities.Image) error {
	return s.repository.Save(image)
}

func NewSave(repository repositories.ImageRepository) *Save {
	return &Save{
		repository: repository,
	}
}
