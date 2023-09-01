package imagesdb

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
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

func (p *PostgresImageRepository) UpdateProcessOnError(process *entities.ImageProcess) error {
	_, err := p.db.ImageProcess.
		UpdateOneID(process.ID).
		SetErrorReason(*process.ErrorReason).
		SetErroredAt(*process.ErroredAt).
		Save(p.ctx)

	if ent.IsNotFound(err) {
		return domainerrors.ErrImageProcessNotFound
	}

	return err
}

func (p *PostgresImageRepository) UpdateProcessOnSuccess(process *entities.ImageProcess, result *entities.Image) error {
	_, err := p.db.Image.Create().
		SetID(result.ID).
		SetSize(result.Size).
		SetFormat(result.Format).
		SetOriginalFilename(result.OriginalFilename).
		SetCreatedAt(result.CreatedAt).
		SetUpdatedAt(result.UpdatedAt).
		Save(p.ctx)
	if err != nil {
		log.Debug(result)
		log.Error(err)
		return err
	}

	_, err = p.db.ImageProcess.
		UpdateOneID(process.ID).
		SetResultID(result.ID).
		SetFinishedAt(*process.FinishedAt).
		Save(p.ctx)

	if ent.IsNotFound(err) {
		return domainerrors.ErrImageProcessNotFound
	}

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (p *PostgresImageRepository) FindProcess(processID uuid.UUID,
) (*entities.ImageProcess, error) {
	// result, err := p.db.ImageProcess.Get(p.ctx, processID)
	result, err := p.db.ImageProcess.Query().WithOrigin().WithResult().Where(imageprocess.ID(processID)).First(p.ctx)
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
		p.mapImageToDomain(result.Edges.Origin),
		p.mapImageToDomain(result.Edges.Result),
	)

	return entity, nil
}

func (p *PostgresImageRepository) mapImageToDomain(ent *ent.Image) *entities.Image {
	entity := entities.NewImage(
		ent.ID,
		ent.Format,
		ent.Size,
		ent.OriginalFilename,
		ent.CreatedAt,
		ent.UpdatedAt,
		ent.DeletedAt,
	)
	return entity
}

func NewPostgresImageRepository(db *ent.Client) *PostgresImageRepository {
	return &PostgresImageRepository{db: db, ctx: context.Background()}
}
