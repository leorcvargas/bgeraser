package repositories

import (
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
)

type ImageProcessRepository interface {
	SaveProcess(process *entities.ImageProcess) error
	UpdateProcessOnError(process *entities.ImageProcess) error
	UpdateProcessOnSuccess(process *entities.ImageProcess) error
	FindProcess(processID uuid.UUID) (*entities.ImageProcess, error)
}
