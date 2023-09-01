package images

import (
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
)

type CreateProcess struct {
	repository Repository
	jobQueue   JobQueue
}

func (c *CreateProcess) Exec(
	id uuid.UUID,
	kind entities.ImageProcessKind,
) (*entities.ImageProcess, error) {
	image, err := c.repository.Find(id)
	if err != nil {
		return nil, err
	}

	process := entities.CreateImageProcess(image.ID, entities.ImageProcessKindRemoveBackground)

	if err := c.repository.SaveProcess(process); err != nil {
		return nil, err
	}

	c.jobQueue <- Job{Payload: process}

	return process, nil
}

func NewCreateProcess(repository Repository, jobQueue JobQueue) *CreateProcess {
	return &CreateProcess{
		repository: repository,
		jobQueue:   jobQueue,
	}
}
