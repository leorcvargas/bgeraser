package queues

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/adjust/rmq/v5"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/domain/repositories"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type removeBackgroundQueueConsumer struct {
	config     *config.Config
	storage    images.Storage
	repository repositories.ImageProcessRepository
}

func NewRemoveBackgroundQueueConsumer(
	config *config.Config,
	storage images.Storage,
	repository repositories.ImageProcessRepository,
) *removeBackgroundQueueConsumer {
	return &removeBackgroundQueueConsumer{config, storage, repository}
}

func (r *removeBackgroundQueueConsumer) Consume(delivery rmq.Delivery) {
	log.Infof("consuming: %v")

	var imageProcess entities.ImageProcess

	unmarshalErr := sonic.Unmarshal([]byte(delivery.Payload()), &imageProcess)
	if unmarshalErr != nil {
		log.Errorf("failed to unmarshal payload: %v", unmarshalErr)
		r.handleConsumeErr(unmarshalErr, delivery)

		return
	}

	log.Infow("payload unmarshalled:", imageProcess)

	execStepsErr := r.ExecSteps(imageProcess)
	if execStepsErr != nil {
		log.Errorf("failed to handle steps result: %v", execStepsErr)
		r.handleConsumeErr(execStepsErr, delivery)

		return
	}

	ackErr := delivery.Ack()
	if ackErr != nil {
		log.Errorf("failed to acknowledge item: %v", ackErr)
		r.handleConsumeErr(ackErr, delivery)

		return
	}
}

func (r *removeBackgroundQueueConsumer) ExecSteps(
	imageProcess entities.ImageProcess,
) error {
	steps := []func(imageProcess *entities.ImageProcess) error{
		r.startStep,
		r.downloadStep,
		r.removeBackgroundStep,
		r.uploadResultStep,
		r.finishStep,
	}

	var err error

	for stepNumber, step := range steps {
		log.Infof("executing step #%d", stepNumber)

		stepErr := step(&imageProcess)
		if stepErr != nil {
			log.Errorf("step #%d failed: %v", stepNumber, stepErr)
			imageProcess.SetError(stepErr)
			err = stepErr
			break
		}
	}

	if err != nil {
		log.Errorf("remove background process error caught: %v", err)
		return r.repository.UpdateProcessOnError(&imageProcess)
	}

	defer r.cleanup(&imageProcess)

	return r.repository.UpdateProcessOnSuccess(&imageProcess)
}

func (r *removeBackgroundQueueConsumer) startStep(
	imageProcess *entities.ImageProcess,
) error {
	imageProcess.StartProcess()
	return nil
}

func (r *removeBackgroundQueueConsumer) downloadStep(
	imageProcess *entities.ImageProcess,
) error {
	content, err := r.storage.Get(imageProcess.Image.Filename())
	if err != nil {
		return err
	}

	workingdir, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.Create(
		workingdir + "/" + r.config.Storage.LocalPath + "/" + imageProcess.Image.Filename(),
	)
	if err != nil {
		return err
	}

	_, err = file.Write(content)

	return err
}

func (r *removeBackgroundQueueConsumer) removeBackgroundStep(
	imageProcess *entities.ImageProcess,
) error {
	filenames := []string{
		imageProcess.Image.Filename(),
		imageProcess.Result.Filename(),
	}

	localPaths := make([]string, 0, len(filenames))
	for _, filename := range filenames {
		var err error
		localPath, err := r.buildFileLocalPath(filename)
		if err != nil {
			return err
		}

		localPaths = append(localPaths, localPath)
	}

	args := []string{"i"}
	args = append(args, localPaths...)

	cmd := exec.Command("rembg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (r *removeBackgroundQueueConsumer) uploadResultStep(
	imageProcess *entities.ImageProcess,
) error {
	workingdir, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.ReadFile(
		workingdir + "/" + r.config.Storage.LocalPath + "/" + imageProcess.Result.Filename(),
	)
	if err != nil {
		return err
	}

	err = r.storage.Upload(imageProcess.Result.Filename(), file)
	if err != nil {
		return err
	}

	return err
}

func (r *removeBackgroundQueueConsumer) finishStep(
	imageProcess *entities.ImageProcess,
) error {
	resultFileLocalPath, err := r.buildFileLocalPath(
		imageProcess.Result.Filename(),
	)
	if err != nil {
		return err
	}

	file, err := os.Open(resultFileLocalPath)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, statErr := file.Stat()
	if statErr != nil {
		return err
	}

	url := fmt.Sprintf(
		"%s/%s",
		r.config.Storage.BucketURL,
		imageProcess.Result.Filename(),
	)

	err = imageProcess.FinishProcess(stat.Name(), stat.Size(), url)
	if err != nil {
		return err
	}

	return nil
}

func (r *removeBackgroundQueueConsumer) cleanup(
	imageProcess *entities.ImageProcess,
) {
	filenames := []string{
		imageProcess.Image.Filename(),
		imageProcess.Result.Filename(),
	}

	localPaths := make([]string, 0, len(filenames))
	for _, filename := range filenames {
		localPath, err := r.buildFileLocalPath(filename)
		if err != nil {
			log.Warnw("cleanup failed to build file local path", err)
			continue
		}

		localPaths = append(localPaths, localPath)
	}

	for _, localPath := range localPaths {
		if err := os.Remove(localPath); err != nil {
			log.Warnf(
				"cleanup failed to delete local file: %s",
				localPath,
			)
		}
	}
}

func (r *removeBackgroundQueueConsumer) buildFileLocalPath(
	filename string,
) (string, error) {
	workdir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf(
		"%s/%s/%s",
		workdir,
		r.config.Storage.LocalPath,
		filename,
	)

	return path, nil
}

func (r *removeBackgroundQueueConsumer) handleStepsResult(
	result *entities.ImageProcess,
	err error,
) error {
	if err != nil {
		log.Errorf("remova background process error caught: %v", err)
		return r.repository.UpdateProcessOnError(result)
	}

	return r.repository.UpdateProcessOnSuccess(result)
}

func (r *removeBackgroundQueueConsumer) handleConsumeErr(
	err error,
	delivery rmq.Delivery,
) {
	log.Error(err)
	if rejectErr := delivery.Reject(); rejectErr != nil {
		log.Errorf("failed to reject delivery: %v", rejectErr)
	}
}
