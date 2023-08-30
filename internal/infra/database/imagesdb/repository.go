package imagesdb

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/ent"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	domainerrors "github.com/leorcvargas/bgeraser/internal/domain/errors"
)

type PostgresImageRepository struct {
	db  *ent.Client
	ctx context.Context
}

func (p *PostgresImageRepository) Save(image *entities.Image) error {
	_, err := p.db.Image.Create().
		SetID(image.ID).
		SetSize(image.Size).
		SetFormat(image.Format).
		SetOriginalFilename(image.OriginalFilename).
		SetCreatedAt(image.CreatedAt).
		SetUpdatedAt(image.UpdatedAt).
		Save(p.ctx)

	return err
}

func (p *PostgresImageRepository) Find(id string) (*entities.Image, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	result, err := p.db.Image.Get(p.ctx, parsedID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domainerrors.ErrImageNotFound
		}

		return nil, err
	}

	deletedAt := time.Time{}
	if result.DeletedAt != nil {
		deletedAt = *result.DeletedAt
	}

	entity := entities.NewImage(
		result.ID,
		result.Format,
		result.Size,
		result.OriginalFilename,
		result.CreatedAt,
		result.UpdatedAt,
		deletedAt,
	)

	return entity, nil
}

func (p *PostgresImageRepository) SaveProcess(process *entities.ImageProcess) error {
	_, err := p.db.ImageProcess.Create().
		SetID(process.ID).
		SetImageID(process.ImageID).
		SetKind(string(process.Kind)).
		Save(p.ctx)

	return err
}

func NewPostgresImageRepository(db *ent.Client) *PostgresImageRepository {
	return &PostgresImageRepository{db: db, ctx: context.Background()}
}
