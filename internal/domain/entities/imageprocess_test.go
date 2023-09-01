package entities_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	domainerrors "github.com/leorcvargas/bgeraser/internal/domain/errors"
)

func TestNewImageProcess(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	originID := uuid.New()
	resultID := uuid.New()
	kind := entities.ImageProcessKindRemoveBackground
	finishedAt := time.Now()
	erroredAt := time.Time{}
	errorReason := ""

	got := entities.NewImageProcess(
		id,
		originID,
		&resultID,
		kind,
		&finishedAt,
		&erroredAt,
		&errorReason,
		nil,
		nil,
	)

	if got.ID != id {
		t.Errorf("expected ID to be %v, got %v", id, got.ID)
	}

	if got.ImageID != originID {
		t.Errorf("expected ImageID to be %v, got %v", originID, got.ImageID)
	}

	if got.ResultID != &resultID {
		t.Errorf("expected ResultID to be %v, got %v", resultID, got.ResultID)
	}

	if got.Kind != kind {
		t.Errorf("expected Kind to be %v, got %v", kind, got.Kind)
	}

	if got.FinishedAt != &finishedAt {
		t.Errorf("expected FinishedAt to be %v, got %v", finishedAt, got.FinishedAt)
	}

	if got.ErroredAt != &erroredAt {
		t.Errorf("expected ErroredAt to be %v, got %v", erroredAt, got.ErroredAt)
	}

	if got.ErrorReason != &errorReason {
		t.Errorf("expected ErrorReason to be %v, got %v", errorReason, got.ErrorReason)
	}
}

func TestCreateImageProcess(t *testing.T) {
	t.Parallel()

	origin := &entities.Image{ID: uuid.New()}
	kind := entities.ImageProcessKindRemoveBackground

	got := entities.CreateImageProcess(
		origin,
		kind,
	)

	if got.ID == uuid.Nil {
		t.Errorf("expected ID to be a valid UUID, got %v", got.ID)
	}

	if got.ImageID != origin.ID {
		t.Errorf("expected ImageID to be %v, got %v", origin.ID, got.ImageID)
	}

	if got.ResultID != nil {
		t.Errorf("expected ResultID to be an empty UUID, got %v", got.ResultID)
	}

	if got.Kind != kind {
		t.Errorf("expected Kind to be %v, got %v", kind, got.Kind)
	}

	if got.FinishedAt != nil {
		t.Errorf("expected FinishedAt to be zero, got %v", got.FinishedAt)
	}

	if got.ErroredAt != nil {
		t.Errorf("expected ErroredAt to be zero, got %v", got.ErroredAt)
	}

	if got.ErrorReason != nil {
		t.Errorf("expected ErrorReason to be empty, got %v", got.ErrorReason)
	}
}

func TestImageProcess_SetFinish(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	originID := uuid.New()
	resultID := uuid.New()
	kind := entities.ImageProcessKindRemoveBackground
	finishedAt := time.Time{}
	erroredAt := time.Time{}
	errorReason := ""

	got := entities.NewImageProcess(
		id,
		originID,
		&resultID,
		kind,
		&finishedAt,
		&erroredAt,
		&errorReason,
		nil,
		nil,
	)

	err := got.SetFinish(resultID)
	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}

	if got.FinishedAt.IsZero() {
		t.Errorf("expected FinishedAt to be not zero, got %v", got.FinishedAt)
	}

	if got.ErroredAt != &erroredAt {
		t.Errorf("expected ErroredAt to be %v, got %v", erroredAt, got.ErroredAt)
	}

	if got.ErrorReason != &errorReason {
		t.Errorf("expected ErrorReason to be %v, got %v", errorReason, got.ErrorReason)
	}
}

func TestImageProcess_SetFinish_Error(t *testing.T) {
	t.Parallel()

	resultID := uuid.Nil
	expected := domainerrors.ErrImageProcessEmptyResultID

	sut := entities.ImageProcess{}

	err := sut.SetFinish(resultID)
	if !errors.Is(err, expected) {
		t.Errorf("expected err to be %v, got %v", expected, err)
	}
}

func TestImageProcess_SetError(t *testing.T) {
	t.Parallel()

	err := errors.New("some error")

	sut := entities.ImageProcess{}
	sut.SetError(err)

	errorReason := err.Error()

	if sut.ErroredAt == nil || sut.ErroredAt.IsZero() {
		t.Errorf("expected ErroredAt to be not zero, got %v", sut.ErroredAt)
	}

	if *sut.ErrorReason != errorReason {
		t.Errorf("expected ErrorReason to be %v, got %v", errorReason, *sut.ErrorReason)
	}
}

func TestImageProcess_Failed(t *testing.T) {
	t.Parallel()

	sut := entities.ImageProcess{ErroredAt: &time.Time{}}

	got := sut.Failed()
	if got {
		t.Errorf("expected Failed to return false, got %v", got)
	}

	now := time.Now()
	sut.ErroredAt = &now

	got = sut.Failed()
	if !got {
		t.Errorf("expected Failed to return true, got %v", got)
	}
}

func TestImageProcess_Finished(t *testing.T) {
	t.Parallel()

	sut := entities.ImageProcess{FinishedAt: &time.Time{}}

	got := sut.Finished()
	if got {
		t.Errorf("expected Finished to return false, got %v", got)
	}

	now := time.Now()
	sut.FinishedAt = &now

	got = sut.Finished()
	if !got {
		t.Errorf("expected Finished to return true, got %v", got)
	}
}
