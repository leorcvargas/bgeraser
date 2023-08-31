package images

import (
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
)

type FindProcess struct {
	repository Repository
}

func (f *FindProcess) Get(imageID, processID uuid.UUID) (*entities.ImageProcess, error) {
	return f.repository.FindProcess(imageID, processID)
}

func NewFindProcess(repository Repository) *FindProcess {
	return &FindProcess{
		repository: repository,
	}
}
