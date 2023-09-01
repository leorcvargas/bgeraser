package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Image struct {
	CreatedAt        time.Time  `json:"-"`
	UpdatedAt        time.Time  `json:"-"`
	DeletedAt        *time.Time `json:"-"`
	Format           string     `json:"format"`
	OriginalFilename string     `json:"originalFilename"`
	Size             int64      `json:"size"`
	ID               uuid.UUID  `json:"id"`
}

func (i *Image) Extension() string {
	return i.Format[6:]
}

func (i *Image) Filename() string {
	return fmt.Sprintf("%s.%s", i.ID.String(), i.Extension())
}

func (i *Image) SetStatInfo(name string, size int64) {
	i.Size = size
	i.OriginalFilename = name
	i.UpdatedAt = time.Now()
}

func CreateImage(
	originalFilename string,
	format string,
	size int64,
) *Image {
	now := time.Now()

	return &Image{
		ID:               uuid.New(),
		Format:           format,
		Size:             size,
		OriginalFilename: originalFilename,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}

func CreateResultImage(format string) *Image {
	now := time.Now()

	return &Image{
		ID:        uuid.New(),
		Format:    format,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewImage(
	id uuid.UUID,
	format string,
	size int64,
	originalFilename string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *Image {
	return &Image{
		ID:               id,
		Format:           format,
		Size:             size,
		OriginalFilename: originalFilename,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		DeletedAt:        deletedAt,
	}
}
