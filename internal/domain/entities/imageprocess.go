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
	FinishedAt  *time.Time       `json:"finishedAt"`
	ResultID    *uuid.UUID       `json:"resultId"`
	ErroredAt   *time.Time       `json:"erroredAt"`
	ErrorReason *string          `json:"errorReason"`
	Image       *Image           `json:"image"`
	Result      *Image           `json:"result"`
	Kind        ImageProcessKind `json:"kind"`
	ID          uuid.UUID        `json:"id"`
	ImageID     uuid.UUID        `json:"imageId"`
}

func (i *ImageProcess) SetError(err error) {
	errorReason := err.Error()
	now := time.Now()

	i.ErroredAt = &now
	i.ErrorReason = &errorReason
}

func (i *ImageProcess) SetFinish(resultID uuid.UUID) error {
	if resultID == uuid.Nil {
		return domainerrors.ErrImageProcessEmptyResultID
	}

	now := time.Now()

	i.FinishedAt = &now
	i.ResultID = &resultID

	return nil
}

func (i *ImageProcess) Failed() bool {
	return !i.ErroredAt.IsZero()
}

func (i *ImageProcess) Finished() bool {
	return !i.FinishedAt.IsZero()
}

func CreateImageProcess(
	image *Image,
	kind ImageProcessKind,
) *ImageProcess {
	return &ImageProcess{
		ID:      uuid.New(),
		ImageID: image.ID,
		Image:   image,
		Kind:    kind,
	}
}

func NewImageProcess(
	id uuid.UUID,
	imageID uuid.UUID,
	resultID *uuid.UUID,
	kind ImageProcessKind,
	finishedAt *time.Time,
	erroredAt *time.Time,
	errorReason *string,
	image *Image,
	result *Image,
) *ImageProcess {
	return &ImageProcess{
		ID:          id,
		ImageID:     imageID,
		ResultID:    resultID,
		Kind:        kind,
		FinishedAt:  finishedAt,
		ErroredAt:   erroredAt,
		ErrorReason: errorReason,
		Image:       image,
		Result:      result,
	}
}
