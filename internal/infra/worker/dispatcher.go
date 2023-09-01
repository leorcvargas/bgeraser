package worker

import "github.com/leorcvargas/bgeraser/internal/domain/images"

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan images.Job
	jobQueue   chan images.Job
	maxWorkers int
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for job := range d.jobQueue {
		go func(job images.Job) {
			jobChannel := <-d.WorkerPool
			jobChannel <- job
		}(job)
	}
}

func NewDispatcher(jobQueue images.JobQueue) *Dispatcher {
	maxWorkers := MaxWorker

	pool := make(chan chan images.Job, maxWorkers)

	return &Dispatcher{
		WorkerPool: pool,
		maxWorkers: maxWorkers,
		jobQueue:   jobQueue,
	}
}
