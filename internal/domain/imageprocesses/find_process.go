package imageprocesses

import (
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"github.com/leorcvargas/bgeraser/internal/domain/repositories"
)

type FindProcess struct {
	repository repositories.ImageProcessRepository
}

func (f *FindProcess) Get(processID uuid.UUID) (*entities.ImageProcess, error) {
	return f.repository.FindProcess(processID)
}

func NewFindProcess(
	repository repositories.ImageProcessRepository,
) *FindProcess {
	return &FindProcess{
		repository: repository,
	}
}
