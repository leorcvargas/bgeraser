package images

import "time"

type Storage interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte, exp time.Duration) error
	Upload(key string, val []byte) error
	Delete(key string) error
	Reset() error
	Close() error
}
