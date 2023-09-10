package processinworker

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/imageprocesses"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
	"github.com/leorcvargas/bgeraser/internal/infra/worker/processes"
)

var MaxWorker = 4

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool          chan chan imageprocesses.ProcessInJob
	ProcessInJobChannel chan imageprocesses.ProcessInJob
	quit                chan bool
	config              *config.Config
	storage             images.Storage
	outQueue            imageprocesses.ProcessOutJobQueue
}

func (w Worker) Start() {
	dataCh := make(chan imageprocesses.ProcessInJob)

	go w.bootstrap(dataCh)

	go func() {
		for data := range dataCh {
			log.Debugf("worker received data", data)
			process := processes.NewRemoveBackgroundProcess(
				data.Payload,
				w.config,
				w.storage,
			)

			imageProcess, err := process.Exec()

			w.outQueue <- imageprocesses.ProcessOutJob{
				Err:     err,
				Payload: imageProcess,
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func (w Worker) bootstrap(dataCh chan imageprocesses.ProcessInJob) {
	for {
		w.WorkerPool <- w.ProcessInJobChannel

		select {
		case job := <-w.ProcessInJobChannel:
			dataCh <- job

		case <-w.quit:
			return
		}
	}
}

func NewWorker(
	workerPool chan chan imageprocesses.ProcessInJob,
	config *config.Config,
	outQueue imageprocesses.ProcessOutJobQueue,
	storage images.Storage,
) Worker {
	return Worker{
		WorkerPool:          workerPool,
		ProcessInJobChannel: make(chan imageprocesses.ProcessInJob),
		quit:                make(chan bool),
		config:              config,
		outQueue:            outQueue,
		storage:             storage,
	}
}
