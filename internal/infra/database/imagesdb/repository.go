package imagesdb

import (
	"context"

	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/ent"
	"github.com/leorcvargas/bgeraser/ent/imageprocess"
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

func (p *PostgresImageRepository) Find(id uuid.UUID) (*entities.Image, error) {
	result, err := p.db.Image.Get(p.ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domainerrors.ErrImageNotFound
		}

		return nil, err
	}

	entity := entities.NewImage(
		result.ID,
		result.Format,
		result.Size,
		result.OriginalFilename,
		result.CreatedAt,
		result.UpdatedAt,
		result.DeletedAt,
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

func (p *PostgresImageRepository) FindProcess(
	imageID, processID uuid.UUID,
) (*entities.ImageProcess, error) {
	result, err := p.db.ImageProcess.
		Query().
		Where(
			imageprocess.And(
				imageprocess.ID(processID),
				imageprocess.ImageIDEQ(imageID),
			),
		).
		Only(p.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domainerrors.ErrImageProcessNotFound
		}

		return nil, err
	}

	var resultID *uuid.UUID
	if result.ResultID != uuid.Nil {
		resultID = &result.ResultID
	}

	entity := entities.NewImageProcess(
		result.ID,
		result.ImageID,
		resultID,
		entities.ImageProcessKind(result.Kind),
		result.FinishedAt,
		result.ErroredAt,
		result.ErrorReason,
	)

	return entity, nil
}

func NewPostgresImageRepository(db *ent.Client) *PostgresImageRepository {
	return &PostgresImageRepository{db: db, ctx: context.Background()}
}
