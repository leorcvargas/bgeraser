package worker

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
	"github.com/leorcvargas/bgeraser/internal/infra/worker/processes"
)

var MaxWorker = 4

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan images.Job
	JobChannel chan images.Job
	quit       chan bool
	config     *config.Config
	repository images.Repository
}

func (w Worker) Start() {
	dataCh := make(chan images.Job)

	go w.bootstrap(dataCh)

	go func() {
		for data := range dataCh {
			log.Debugf("worker received data", data)
			process := processes.NewRemoveBackgroundProcess(&data.Payload, w.config)

			err := process.Exec()
			if err != nil {
				log.Errorw("error while executing process", err, process)

				updateErr := w.repository.UpdateProcessOnError(process.ImageProcess)
				if updateErr != nil {
					log.Warnw("failed to update process", err, process)
				}

				return
			}

			updateErr := w.repository.UpdateProcessOnSuccess(
				process.ImageProcess,
				process.ResultImage,
			)
			if updateErr != nil {
				log.Errorw("failed to update process", err)
			}
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

func NewWorker(workerPool chan chan images.Job, repository images.Repository, config *config.Config) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan images.Job),
		quit:       make(chan bool),
		repository: repository,
		config:     config,
	}
}
