package processoutworker

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
)

var MaxWorker = 4

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool           chan chan images.ProcessOutJob
	ProcessOutJobChannel chan images.ProcessOutJob
	quit                 chan bool
	repository           images.Repository
}

func (w Worker) Start() {
	dataCh := make(chan images.ProcessOutJob)

	go w.bootstrap(dataCh)

	go func() {
		for data := range dataCh {
			log.Debugf("process out worker received data", data)

			if data.Err != nil {
				log.Errorw("error while executing process", data.Err, data.Payload)

				err := w.repository.UpdateProcessOnError(&data.Payload)
				if err != nil {
					log.Warnw("failed to update process", data.Err, data.Payload)
				}

				return
			}

			err := w.repository.UpdateProcessOnSuccess(&data.Payload)
			if err != nil {
				log.Errorw("process out failed to update process", err)
			}

			log.Debug("process completed successfuly")
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func (w Worker) bootstrap(dataCh chan images.ProcessOutJob) {
	for {
		w.WorkerPool <- w.ProcessOutJobChannel

		select {
		case job := <-w.ProcessOutJobChannel:
			dataCh <- job

		case <-w.quit:
			return
		}
	}
}

func NewWorker(workerPool chan chan images.ProcessOutJob, repository images.Repository) Worker {
	return Worker{
		WorkerPool:           workerPool,
		ProcessOutJobChannel: make(chan images.ProcessOutJob),
		quit:                 make(chan bool),
		repository:           repository,
	}
}
