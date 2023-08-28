package images

import "github.com/leorcvargas/bgeraser/internal/domain/entities"

type Repository interface {
	Save(image *entities.Image) error
}
