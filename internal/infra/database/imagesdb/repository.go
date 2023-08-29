package imagesdb

import (
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"gorm.io/gorm"
)

// images.Repository implementation for Postgres.
type PostgresImageRepository struct {
	db *gorm.DB
}

func (p *PostgresImageRepository) Save(image *entities.Image) error {
	model := &Model{
		ID:               image.ID,
		Format:           image.Format,
		Size:             image.Size,
		OriginalFilename: image.OriginalFilename,
		SavedAt:          image.SavedAt,
	}

	return p.db.Create(model).Error
}

func NewPostgresImageRepository(db *gorm.DB) *PostgresImageRepository {
	return &PostgresImageRepository{db: db}
}
