package imagesdb

import (
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model

	ID               uuid.UUID `gorm:"primaryKey, type:uuid"`
	Format           string    `gorm:"type:varchar(16); not null"`
	Size             int64     `gorm:"not null"`
	OriginalFilename string    `gorm:"type:varchar(255); not null"`
}

func (Model) TableName() string {
	return "images"
}

func NewModel(image *entities.Image) *Model {
	return &Model{
		ID:               image.ID,
		Format:           image.Format,
		Size:             image.Size,
		OriginalFilename: image.OriginalFilename,
	}
}

func NewModelFromDomain(image *entities.Image) *Model {
	return &Model{
		ID:               image.ID,
		Format:           image.Format,
		Size:             image.Size,
		OriginalFilename: image.OriginalFilename,
	}
}
