package images

import (
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
)

type Repository interface {
	Save(image *entities.Image) error
	Find(id uuid.UUID) (*entities.Image, error)
	SaveProcess(process *entities.ImageProcess) error
	FindProcess(imageID, processID uuid.UUID) (*entities.ImageProcess, error)
}
