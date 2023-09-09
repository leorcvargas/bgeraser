package processoutworker

import (
	"github.com/leorcvargas/bgeraser/internal/domain/imageprocesses"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type ProcessOutDispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan imageprocesses.ProcessOutJob
	jobQueue   chan imageprocesses.ProcessOutJob
	repository imageprocesses.Repository
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
		go func(job imageprocesses.ProcessOutJob) {
			jobChannel := <-d.WorkerPool
			jobChannel <- job
		}(job)
	}
}

func NewDispatcher(
	jobQueue imageprocesses.ProcessOutJobQueue,
	repository imageprocesses.Repository,
	config *config.Config,
) *ProcessOutDispatcher {
	maxWorkers := MaxWorker

	pool := make(chan chan imageprocesses.ProcessOutJob, maxWorkers)

	return &ProcessOutDispatcher{
		WorkerPool: pool,
		maxWorkers: maxWorkers,
		jobQueue:   jobQueue,
		repository: repository,
		config:     config,
	}
}
