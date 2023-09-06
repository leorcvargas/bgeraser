package storage

import (
	"bytes"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type S3Storage struct {
	sess   *session.Session
	config *config.Config
}

func (s *S3Storage) Bucket() string {
	return s.config.Storage.Bucket
}

func (s *S3Storage) Get(key string) ([]byte, error) {
	downloader := s3manager.NewDownloader(s.sess)

	input := &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket()),
		Key:    aws.String(key),
	}

	buf := &aws.WriteAtBuffer{}
	if _, err := downloader.Download(buf, input); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *S3Storage) Set(key string, val []byte, _ time.Duration) error {
	return s.upload(key, val)
}

func (s *S3Storage) Upload(key string, val []byte) error {
	return s.upload(key, val)
}

func (s *S3Storage) Delete(key string) error {
	return errors.New("not implemented")
}

func (s *S3Storage) Reset() error {
	return errors.New("not implemented")
}

func (s *S3Storage) Close() error {
	return errors.New("not implemented")
}

func (s *S3Storage) upload(key string, val []byte) error {
	uploader := s3manager.NewUploader(s.sess)
	reader := bytes.NewReader(val)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Bucket()),
		Key:    aws.String(key),
		Body:   reader,
		ACL:    aws.String("public-read"),
	})

	return err
}

func NewS3Storage(config *config.Config) images.Storage {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWS.Region),
	})
	if err != nil {
		log.Fatalw("failed to start s3 session", err)
	}

	return &S3Storage{sess, config}
}
