package processes

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type RemoveBackgroundProcess struct {
	config       *config.Config
	imageProcess entities.ImageProcess
	storage      images.Storage
}

func (r *RemoveBackgroundProcess) Exec() (*entities.ImageProcess, error) {
	log.Debugw("starting remove background process", r.imageProcess)
	r.imageProcess.StartProcess()

	if err := r.downloadImage(); err != nil {
		return nil, r.handleError(err)
	}

	if err := r.runCommand(); err != nil {
		return nil, r.handleError(err)
	}

	if err := r.uploadResultImage(); err != nil {
		return nil, r.handleError(err)
	}

	if err := r.finish(); err != nil {
		return nil, r.handleError(err)
	}

	go r.cleanup()

	return &r.imageProcess, nil
}

func (r *RemoveBackgroundProcess) downloadImage() error {
	content, err := r.storage.Get(r.imageProcess.Image.Filename())
	if err != nil {
		return err
	}

	workingdir, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.Create(
		workingdir + "/" + r.config.Storage.LocalPath + "/" + r.imageProcess.Image.Filename(),
	)
	if err != nil {
		return err
	}

	_, err = file.Write(content)

	return err
}

func (r *RemoveBackgroundProcess) uploadResultImage() error {
	workingdir, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.ReadFile(
		workingdir + "/" + r.config.Storage.LocalPath + "/" + r.imageProcess.Result.Filename(),
	)
	if err != nil {
		return err
	}

	if err = r.storage.Upload(r.imageProcess.Result.Filename(), file); err != nil {
		return err
	}

	return err
}

func (r *RemoveBackgroundProcess) handleError(err error) error {
	r.imageProcess.SetError(err)

	return err
}

func (r *RemoveBackgroundProcess) finish() error {
	resultFileLocalPath, err := r.buildFileLocalPath(
		r.imageProcess.Result.Filename(),
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
		r.imageProcess.Result.Filename(),
	)
	if err := r.imageProcess.FinishProcess(stat.Name(), stat.Size(), url); err != nil {
		return err
	}

	return nil
}

func (r *RemoveBackgroundProcess) cleanup() {
	filenames := []string{
		r.imageProcess.Image.Filename(),
		r.imageProcess.Result.Filename(),
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

func (r *RemoveBackgroundProcess) runCommand() error {
	command, err := r.buildCommand()
	if err != nil {
		return err
	}

	err = command.Start()
	if err != nil {
		return nil
	}

	return command.Wait()
}

func (r *RemoveBackgroundProcess) buildCommand() (*exec.Cmd, error) {
	if r.imageProcess.Result == nil {
		return nil, errors.New("RemoveBackgroundProcess.resultImage is empty")
	}

	workdir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cmd := &exec.Cmd{
		Path: "/bin/sh",
		Args: []string{
			"-c",
			workdir + "/scripts/rembg/run.sh",
			r.imageProcess.Image.Filename(),
			r.imageProcess.Result.Filename(),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	return cmd, nil
}

func (r *RemoveBackgroundProcess) buildFileLocalPath(
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

func NewRemoveBackgroundProcess(
	imageProcess entities.ImageProcess,
	config *config.Config,
	storage images.Storage,
) *RemoveBackgroundProcess {
	return &RemoveBackgroundProcess{
		imageProcess: imageProcess,
		config:       config,
		storage:      storage,
	}
}
