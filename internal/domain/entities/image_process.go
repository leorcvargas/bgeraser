package entities

import (
	"time"

	"github.com/google/uuid"
	domainerrors "github.com/leorcvargas/bgeraser/internal/domain/errors"
)

type ImageProcessKind string

const (
	ImageProcessKindRemoveBackground = ImageProcessKind("REMOVE_BACKGROUND")
)

type ImageProcess struct {
	ID          uuid.UUID
	OriginID    uuid.UUID
	ResultID    uuid.UUID
	Kind        ImageProcessKind
	FinishedAt  time.Time
	ErroredAt   time.Time
	ErrorReason string
}

func (i *ImageProcess) SetError(err error) {
	i.ErroredAt = time.Now()
	i.ErrorReason = err.Error()
}

func (i *ImageProcess) SetFinish(resultID uuid.UUID) error {
	if resultID == uuid.Nil {
		return domainerrors.ErrImageProcessEmptyResultID
	}

	i.FinishedAt = time.Now()
	i.ResultID = resultID

	return nil
}

func (i *ImageProcess) Failed() bool {
	return !i.ErroredAt.IsZero()
}

func (i *ImageProcess) Finished() bool {
	return !i.FinishedAt.IsZero()
}

func CreateImageProcess(
	originID uuid.UUID,
	kind ImageProcessKind,
) *ImageProcess {
	return &ImageProcess{
		ID:       uuid.New(),
		OriginID: originID,
		Kind:     kind,
	}
}

func NewImageProcess(
	id uuid.UUID,
	originID uuid.UUID,
	resultID uuid.UUID,
	kind ImageProcessKind,
	finishedAt time.Time,
	erroredAt time.Time,
	errorReason string,
) *ImageProcess {
	return &ImageProcess{
		ID:          id,
		OriginID:    originID,
		ResultID:    resultID,
		Kind:        kind,
		FinishedAt:  finishedAt,
		ErroredAt:   erroredAt,
		ErrorReason: errorReason,
	}
}
