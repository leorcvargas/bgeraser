package processinworker

import (
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type ProcessInDispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool  chan chan images.ProcessInJob
	jobQueue    images.ProcessInJobQueue
	outJobQueue images.ProcessOutJobQueue
	repository  images.Repository
	config      *config.Config
	maxWorkers  int
}

func (d *ProcessInDispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, d.repository, d.config, d.outJobQueue)
		worker.Start()
	}

	go d.dispatch()
}

func (d *ProcessInDispatcher) dispatch() {
	for job := range d.jobQueue {
		go func(job images.ProcessInJob) {
			jobChannel := <-d.WorkerPool
			jobChannel <- job
		}(job)
	}
}

func NewDispatcher(jobQueue images.ProcessInJobQueue, outJobQueue images.ProcessOutJobQueue, repository images.Repository, config *config.Config) *ProcessInDispatcher {
	maxWorkers := MaxWorker

	pool := make(chan chan images.ProcessInJob, maxWorkers)

	return &ProcessInDispatcher{
		WorkerPool:  pool,
		maxWorkers:  maxWorkers,
		jobQueue:    jobQueue,
		outJobQueue: outJobQueue,
		repository:  repository,
		config:      config,
	}
}
