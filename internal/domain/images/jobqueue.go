package images

import "github.com/leorcvargas/bgeraser/internal/domain/entities"

type Job struct {
	Payload entities.ImageProcess
}

type JobQueue chan Job

func NewJobQueue() JobQueue {
	maxQueue := 1
	return make(JobQueue, maxQueue)
}
