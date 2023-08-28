package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Image struct {
	ID               uuid.UUID `json:"id"`
	Format           string    `json:"format"`
	Size             int64     `json:"size"`
	OriginalFilename string    `json:"originalFilename"`
	SavedAt          time.Time `json:"savedAt"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	DeletedAt        time.Time `json:"deletedAt"`
}

func (i *Image) Extension() string {
	return i.Format[6:]
}

func (i *Image) Filename() string {
	return fmt.Sprintf("%s.%s", i.ID.String(), i.Extension())
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

func NewImage(
	id uuid.UUID,
	format string,
	size int64,
	originId string,
	savedAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) *Image {
	return &Image{
		ID:        id,
		Format:    format,
		Size:      size,
		SavedAt:   savedAt,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
