package images

import "github.com/leorcvargas/bgeraser/internal/domain/entities"

type Save struct {
	repository Repository
}

func (s *Save) Exec(image *entities.Image) error {
	return s.repository.Save(image)
}

func NewSave(repository Repository) *Save {
	return &Save{
		repository: repository,
	}
}
