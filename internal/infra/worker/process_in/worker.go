package processinworker

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
	"github.com/leorcvargas/bgeraser/internal/infra/worker/processes"
)

var MaxWorker = 4

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool          chan chan images.ProcessInJob
	ProcessInJobChannel chan images.ProcessInJob
	quit                chan bool
	config              *config.Config
	repository          images.Repository
	outQueue            images.ProcessOutJobQueue
}

func (w Worker) Start() {
	dataCh := make(chan images.ProcessInJob)

	go w.bootstrap(dataCh)

	go func() {
		for data := range dataCh {
			log.Debugf("worker received data", data)
			process := processes.NewRemoveBackgroundProcess(data.Payload, w.config)

			imageProcess, err := process.Exec()

			w.outQueue <- images.ProcessOutJob{
				Err:     err,
				Payload: *imageProcess,
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func (w Worker) bootstrap(dataCh chan images.ProcessInJob) {
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
	workerPool chan chan images.ProcessInJob,
	repository images.Repository,
	config *config.Config,
	outQueue images.ProcessOutJobQueue,
) Worker {
	return Worker{
		WorkerPool:          workerPool,
		ProcessInJobChannel: make(chan images.ProcessInJob),
		quit:                make(chan bool),
		repository:          repository,
		config:              config,
		outQueue:            outQueue,
	}
}
