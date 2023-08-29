package entities_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
)

func TestCreateImage(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc             string
		originalFilename string
		format           string
		size             int64
	}{
		{
			desc:             "create a PNG image with random size",
			originalFilename: "fakepngimage.png",
			format:           "image/png",
			size:             1024 * 3,
		},
		{
			desc:             "create a JPEG image with random size",
			originalFilename: "fakepngimage.jpeg",
			format:           "image/jpeg",
			size:             1024 * 0.5,
		},
	}

	for _, testCase := range testCases {
		tC := testCase

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			image := entities.CreateImage(tC.originalFilename, tC.format, tC.size)

			if image == nil {
				t.Error("expected image to not be nil")
				return
			}

			if image.OriginalFilename != tC.originalFilename {
				t.Errorf(
					"expected image.OriginalFilename to match %s, got %s",
					tC.originalFilename,
					image.OriginalFilename,
				)
				return
			}

			if image.Size != tC.size {
				t.Errorf("expected image.Size to match %d, got %d", tC.size, image.Size)
			}

			if image.CreatedAt.IsZero() {
				t.Error("image.CreatedAt should not be zero")
			}

			if !image.DeletedAt.IsZero() {
				t.Error("image.DeletedAt should be zero")
			}
		})
	}
}

func TestNewImage(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	format := "application/jpeg"
	size := int64(3412894)
	createdAt := time.Now()
	updatedAt := time.Now()
	deletedAt := time.Time{}

	got := entities.NewImage(
		id,
		format,
		size,
		createdAt,
		updatedAt,
		deletedAt,
	)

	if got == nil {
		t.Error("expected NewImage to no return nil")
		return
	}

	if got.ID != id {
		t.Errorf("expected ID to match %s, got %s", id, got.ID)
	}

	if got.Format != format {
		t.Errorf("expected Format to match %s, got %s", format, got.Format)
	}

	if got.Size != size {
		t.Errorf("expected Size to match %d, got %d", size, got.Size)
	}

	if got.CreatedAt != createdAt {
		t.Errorf("expected CreatedAt to match %v, got %v", createdAt, got.CreatedAt)
	}

	if got.UpdatedAt != updatedAt {
		t.Errorf("expected UpdatedAt to match %v, got %v", updatedAt, got.UpdatedAt)
	}

	if !got.DeletedAt.IsZero() {
		t.Errorf("expected DeletedAt to be zero")
	}
}

func TestImage_Extension(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		sut      *entities.Image
		expected string
	}{
		{
			desc:     "return PNG file extension",
			sut:      entities.CreateImage("testfile.png", "image/png", int64(1024)),
			expected: "png",
		},
		{
			desc:     "return JPEG file extension",
			sut:      entities.CreateImage("testfile.jpeg", "image/jpeg", int64(1024)),
			expected: "jpeg",
		},
	}

	for _, testCase := range testCases {
		tC := testCase

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			got := tC.sut.Extension()

			if got != tC.expected {
				t.Errorf("expected file extension %s, got %s", tC.expected, got)
			}
		})
	}
}

func TestImage_Filename(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc  string
		sut   *entities.Image
		sufix string
	}{
		{
			desc:  "return filename for PNG file",
			sut:   entities.CreateImage("testfile.png", "image/png", int64(1024)),
			sufix: "png",
		},
		{
			desc:  "return filename for JPEG file",
			sut:   entities.CreateImage("testfile.jpeg", "image/jpeg", int64(1024)),
			sufix: "jpeg",
		},
	}

	for _, testCase := range testCases {
		tC := testCase

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			got := tC.sut.Filename()
			expected := fmt.Sprintf("%s.%s", tC.sut.ID, tC.sufix)

			if got != expected {
				t.Errorf("expected file extension %s, got %s", expected, got)
			}
		})
	}
}
