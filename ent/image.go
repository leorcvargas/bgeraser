// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/ent/image"
)

// Image is the model entity for the Image schema.
type Image struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Format holds the value of the "format" field.
	Format string `json:"format,omitempty"`
	// Size holds the value of the "size" field.
	Size int64 `json:"size,omitempty"`
	// OriginalFilename holds the value of the "original_filename" field.
	OriginalFilename string `json:"original_filename,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImageQuery when eager-loading is set.
	Edges        ImageEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ImageEdges holds the relations/edges for other nodes in the graph.
type ImageEdges struct {
	// Images holds the value of the images edge.
	Images []*Image `json:"images,omitempty"`
	// ImageProcesses holds the value of the image_processes edge.
	ImageProcesses []*ImageProcess `json:"image_processes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading.
func (e ImageEdges) ImagesOrErr() ([]*Image, error) {
	if e.loadedTypes[0] {
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// ImageProcessesOrErr returns the ImageProcesses value or an error if the edge
// was not loaded in eager-loading.
func (e ImageEdges) ImageProcessesOrErr() ([]*ImageProcess, error) {
	if e.loadedTypes[1] {
		return e.ImageProcesses, nil
	}
	return nil, &NotLoadedError{edge: "image_processes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Image) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case image.FieldSize:
			values[i] = new(sql.NullInt64)
		case image.FieldFormat, image.FieldOriginalFilename, image.FieldURL:
			values[i] = new(sql.NullString)
		case image.FieldCreatedAt, image.FieldUpdatedAt, image.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case image.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Image fields.
func (i *Image) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case image.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case image.FieldFormat:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field format", values[j])
			} else if value.Valid {
				i.Format = value.String
			}
		case image.FieldSize:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[j])
			} else if value.Valid {
				i.Size = value.Int64
			}
		case image.FieldOriginalFilename:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field original_filename", values[j])
			} else if value.Valid {
				i.OriginalFilename = value.String
			}
		case image.FieldURL:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[j])
			} else if value.Valid {
				i.URL = value.String
			}
		case image.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case image.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case image.FieldDeletedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[j])
			} else if value.Valid {
				i.DeletedAt = new(time.Time)
				*i.DeletedAt = value.Time
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Image.
// This includes values selected through modifiers, order, etc.
func (i *Image) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryImages queries the "images" edge of the Image entity.
func (i *Image) QueryImages() *ImageQuery {
	return NewImageClient(i.config).QueryImages(i)
}

// QueryImageProcesses queries the "image_processes" edge of the Image entity.
func (i *Image) QueryImageProcesses() *ImageProcessQuery {
	return NewImageClient(i.config).QueryImageProcesses(i)
}

// Update returns a builder for updating this Image.
// Note that you need to call Image.Unwrap() before calling this method if this Image
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Image) Update() *ImageUpdateOne {
	return NewImageClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Image entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Image) Unwrap() *Image {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Image is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Image) String() string {
	var builder strings.Builder
	builder.WriteString("Image(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("format=")
	builder.WriteString(i.Format)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", i.Size))
	builder.WriteString(", ")
	builder.WriteString("original_filename=")
	builder.WriteString(i.OriginalFilename)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(i.URL)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := i.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Images is a parsable slice of Image.
type Images []*Image
