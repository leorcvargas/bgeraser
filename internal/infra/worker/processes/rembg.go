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
	imageProcess entities.ImageProcess
}

func (r *RemoveBackgroundProcess) Exec() (*entities.ImageProcess, error) {
	r.imageProcess.StartProcess()

	if err := r.runCommand(); err != nil {
		return nil, r.handleError(err)
	}

	if err := r.finish(); err != nil {
		return nil, r.handleError(err)
	}

	return &r.imageProcess, nil
}

func (r *RemoveBackgroundProcess) handleError(err error) error {
	r.imageProcess.SetError(err)

	return err
}

func (r *RemoveBackgroundProcess) finish() error {
	workingdir, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.Open(workingdir + "/" + r.config.Storage.LocalPath + "/" + r.imageProcess.Result.Filename())
	if err != nil {
		return err
	}
	defer file.Close()

	stat, statErr := file.Stat()
	if statErr != nil {
		return err
	}

	return r.imageProcess.FinishProcess(stat.Name(), stat.Size())
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

	workingdir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cmd := &exec.Cmd{
		Path: "/bin/sh",
		Args: []string{
			"-c",
			workingdir + "/scripts/rembg/run.sh",
			r.imageProcess.Image.Filename(),
			r.imageProcess.Result.Filename(),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	log.Infof("command build: %s", strings.Join(cmd.Args, " "))

	return cmd, nil
}

func NewRemoveBackgroundProcess(imageProcess entities.ImageProcess, config *config.Config) *RemoveBackgroundProcess {
	return &RemoveBackgroundProcess{imageProcess: imageProcess, config: config}
}
