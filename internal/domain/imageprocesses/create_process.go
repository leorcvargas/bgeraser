package imageprocesses

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"github.com/leorcvargas/bgeraser/internal/domain/repositories"
	"github.com/leorcvargas/bgeraser/internal/producers"
)

type CreateProcess struct {
	repository      repositories.ImageProcessRepository
	imageRepository repositories.ImageRepository
	producer        producers.Producer[*entities.ImageProcess]
}

func (c *CreateProcess) Exec(
	id uuid.UUID,
	kind entities.ImageProcessKind,
) (*entities.ImageProcess, error) {
	image, err := c.imageRepository.Find(id)
	if err != nil {
		return nil, err
	}

	process := entities.CreateImageProcess(
		image,
		entities.ImageProcessKindRemoveBackground,
	)

	errCh := make(chan error, 2)
	defer close(errCh)

	go c.save(process, errCh)
	go c.publish(process, errCh)

	for i := 0; i < 2; i++ {
		err = <-errCh
		if err != nil {
			return nil, err
		}
	}

	return process, nil
}

func (c *CreateProcess) save(process *entities.ImageProcess, errCh chan error) {
	err := c.repository.SaveProcess(process)
	if err != nil {
		log.Errorf("failed to save process: %v", err)
	}

	errCh <- err
}

func (c *CreateProcess) publish(
	process *entities.ImageProcess,
	errCh chan error,
) {
	err := c.producer.Send(process)
	if err != nil {
		log.Errorf("failed to enqueue process: %v", err)
	}

	errCh <- err
}

func NewCreateProcess(
	repository repositories.ImageProcessRepository,
	imageRepository repositories.ImageRepository,
	producer producers.Producer[*entities.ImageProcess],
) *CreateProcess {
	return &CreateProcess{
		repository:      repository,
		imageRepository: imageRepository,
		producer:        producer,
	}
}
