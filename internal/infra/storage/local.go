package storage

import (
	"fmt"
	"os"

	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type LocalImageStorage struct {
	config *config.Config
}

func (s *LocalImageStorage) Write(filename string, content []byte) error {
	path := fmt.Sprintf("%s/%s", s.config.Storage.LocalPath, filename)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(content); err != nil {
		return err
	}

	return nil
}

func NewLocalImageStorage(config *config.Config) *LocalImageStorage {
	return &LocalImageStorage{
		config: config,
	}
}
