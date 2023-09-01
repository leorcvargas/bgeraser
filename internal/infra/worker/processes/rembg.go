package processes

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type RemoveBackgroundProcess struct {
	config       *config.Config
	ImageProcess *entities.ImageProcess
	ResultImage  *entities.Image
}

func (r *RemoveBackgroundProcess) Exec() error {
	r.setResultImage()

	if err := r.runCommand(); err != nil {
		return r.handleError(err)
	}

	if err := r.finish(); err != nil {
		return r.handleError(err)
	}

	return nil
}

func (r *RemoveBackgroundProcess) handleError(err error) error {
	r.ImageProcess.SetError(err)

	return err
}

func (r *RemoveBackgroundProcess) finish() error {
	workingdir, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.Open(workingdir + "/" + r.config.Storage.LocalPath + "/" + r.ResultImage.Filename())
	if err != nil {
		return err
	}
	defer file.Close()

	stat, statErr := file.Stat()
	if statErr != nil {
		return err
	}

	r.ResultImage.SetStatInfo(stat.Name(), stat.Size())
	r.ImageProcess.SetFinish(r.ResultImage.ID)

	return nil
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

func (r *RemoveBackgroundProcess) setResultImage() {
	r.ResultImage = entities.CreateResultImage("image/png")
}

func (r *RemoveBackgroundProcess) buildCommand() (*exec.Cmd, error) {
	if r.ResultImage == nil {
		return nil, errors.New("RemoveBackgroundProcess.resultImage is empty")
	}

	workingdir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cmd := &exec.Cmd{
		Path: "/bin/sh",
		Args: []string{
			"-c",
			workingdir + "/scripts/rembg/run.sh",
			r.ImageProcess.Image.Filename(),
			r.ResultImage.Filename(),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	log.Infof("command build: %s", strings.Join(cmd.Args, " "))

	return cmd, nil
}

func NewRemoveBackgroundProcess(imageProcess *entities.ImageProcess, config *config.Config) *RemoveBackgroundProcess {
	return &RemoveBackgroundProcess{ImageProcess: imageProcess, config: config}
}
