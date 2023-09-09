package processinworker

import (
	"github.com/leorcvargas/bgeraser/internal/domain/imageprocesses"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type ProcessInDispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool  chan chan imageprocesses.ProcessInJob
	jobQueue    imageprocesses.ProcessInJobQueue
	outJobQueue imageprocesses.ProcessOutJobQueue
	config      *config.Config
	storage     images.Storage
	maxWorkers  int
}

func (d *ProcessInDispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(
			d.WorkerPool,
			d.config,
			d.outJobQueue,
			d.storage,
		)
		worker.Start()
	}

	go d.dispatch()
}

func (d *ProcessInDispatcher) dispatch() {
	for job := range d.jobQueue {
		go func(job imageprocesses.ProcessInJob) {
			jobChannel := <-d.WorkerPool
			jobChannel <- job
		}(job)
	}
}

func NewDispatcher(
	jobQueue imageprocesses.ProcessInJobQueue,
	outJobQueue imageprocesses.ProcessOutJobQueue,
	config *config.Config,
	storage images.Storage,
) *ProcessInDispatcher {
	maxWorkers := MaxWorker

	pool := make(chan chan imageprocesses.ProcessInJob, maxWorkers)

	return &ProcessInDispatcher{
		WorkerPool:  pool,
		maxWorkers:  maxWorkers,
		jobQueue:    jobQueue,
		outJobQueue: outJobQueue,
		config:      config,
		storage:     storage,
	}
}
