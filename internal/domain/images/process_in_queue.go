package images

import "github.com/leorcvargas/bgeraser/internal/domain/entities"

type ProcessInJob struct {
	Payload entities.ImageProcess
}

type ProcessInJobQueue chan ProcessInJob

func NewProcessInJobQueue() ProcessInJobQueue {
	maxQueue := 1
	return make(ProcessInJobQueue, maxQueue)
}
