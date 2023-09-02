package processoutworker

import (
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type ProcessOutDispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan images.ProcessOutJob
	jobQueue   chan images.ProcessOutJob
	repository images.Repository
	config     *config.Config
	maxWorkers int
}

func (d *ProcessOutDispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, d.repository)
		worker.Start()
	}

	go d.dispatch()
}

func (d *ProcessOutDispatcher) dispatch() {
	for job := range d.jobQueue {
		go func(job images.ProcessOutJob) {
			jobChannel := <-d.WorkerPool
			jobChannel <- job
		}(job)
	}
}

func NewDispatcher(jobQueue images.ProcessOutJobQueue, repository images.Repository, config *config.Config) *ProcessOutDispatcher {
	maxWorkers := MaxWorker

	pool := make(chan chan images.ProcessOutJob, maxWorkers)

	return &ProcessOutDispatcher{
		WorkerPool: pool,
		maxWorkers: maxWorkers,
		jobQueue:   jobQueue,
		repository: repository,
		config:     config,
	}
}
