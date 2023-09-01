package worker

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
)

var MaxWorker = 4

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan images.Job
	JobChannel chan images.Job
	quit       chan bool
}

func (w Worker) Start() {
	dataCh := make(chan images.Job)

	go w.bootstrap(dataCh)

	go func() {
		for data := range dataCh {
			log.Debugf("worker received data: %w", data)
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func (w Worker) bootstrap(dataCh chan images.Job) {
	for {
		w.WorkerPool <- w.JobChannel

		select {
		case job := <-w.JobChannel:
			dataCh <- job

		case <-w.quit:
			return
		}
	}
}

func NewWorker(workerPool chan chan images.Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan images.Job),
		quit:       make(chan bool),
	}
}
