package images

import "github.com/leorcvargas/bgeraser/internal/domain/entities"

type ProcessOutJob struct {
	Err     error
	Payload entities.ImageProcess
}

type ProcessOutJobQueue chan ProcessOutJob

func NewProcessOutJobQueue() ProcessOutJobQueue {
	maxQueue := 1
	return make(ProcessOutJobQueue, maxQueue)
}
