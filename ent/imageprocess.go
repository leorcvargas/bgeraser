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
	"github.com/leorcvargas/bgeraser/ent/imageprocess"
)

// ImageProcess is the model entity for the ImageProcess schema.
type ImageProcess struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ImageID holds the value of the "image_id" field.
	ImageID uuid.UUID `json:"image_id,omitempty"`
	// ResultID holds the value of the "result_id" field.
	ResultID uuid.UUID `json:"result_id,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind string `json:"kind,omitempty"`
	// FinishedAt holds the value of the "finished_at" field.
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	// ErroredAt holds the value of the "errored_at" field.
	ErroredAt *time.Time `json:"errored_at,omitempty"`
	// ErrorReason holds the value of the "error_reason" field.
	ErrorReason *string `json:"error_reason,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImageProcessQuery when eager-loading is set.
	Edges        ImageProcessEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ImageProcessEdges holds the relations/edges for other nodes in the graph.
type ImageProcessEdges struct {
	// Origin holds the value of the origin edge.
	Origin *Image `json:"origin,omitempty"`
	// Result holds the value of the result edge.
	Result *Image `json:"result,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OriginOrErr returns the Origin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageProcessEdges) OriginOrErr() (*Image, error) {
	if e.loadedTypes[0] {
		if e.Origin == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: image.Label}
		}
		return e.Origin, nil
	}
	return nil, &NotLoadedError{edge: "origin"}
}

// ResultOrErr returns the Result value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageProcessEdges) ResultOrErr() (*Image, error) {
	if e.loadedTypes[1] {
		if e.Result == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: image.Label}
		}
		return e.Result, nil
	}
	return nil, &NotLoadedError{edge: "result"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ImageProcess) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case imageprocess.FieldKind, imageprocess.FieldErrorReason:
			values[i] = new(sql.NullString)
		case imageprocess.FieldFinishedAt, imageprocess.FieldErroredAt:
			values[i] = new(sql.NullTime)
		case imageprocess.FieldID, imageprocess.FieldImageID, imageprocess.FieldResultID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ImageProcess fields.
func (ip *ImageProcess) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case imageprocess.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ip.ID = *value
			}
		case imageprocess.FieldImageID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field image_id", values[i])
			} else if value != nil {
				ip.ImageID = *value
			}
		case imageprocess.FieldResultID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field result_id", values[i])
			} else if value != nil {
				ip.ResultID = *value
			}
		case imageprocess.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				ip.Kind = value.String
			}
		case imageprocess.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				ip.FinishedAt = new(time.Time)
				*ip.FinishedAt = value.Time
			}
		case imageprocess.FieldErroredAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field errored_at", values[i])
			} else if value.Valid {
				ip.ErroredAt = new(time.Time)
				*ip.ErroredAt = value.Time
			}
		case imageprocess.FieldErrorReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error_reason", values[i])
			} else if value.Valid {
				ip.ErrorReason = new(string)
				*ip.ErrorReason = value.String
			}
		default:
			ip.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ImageProcess.
// This includes values selected through modifiers, order, etc.
func (ip *ImageProcess) Value(name string) (ent.Value, error) {
	return ip.selectValues.Get(name)
}

// QueryOrigin queries the "origin" edge of the ImageProcess entity.
func (ip *ImageProcess) QueryOrigin() *ImageQuery {
	return NewImageProcessClient(ip.config).QueryOrigin(ip)
}

// QueryResult queries the "result" edge of the ImageProcess entity.
func (ip *ImageProcess) QueryResult() *ImageQuery {
	return NewImageProcessClient(ip.config).QueryResult(ip)
}

// Update returns a builder for updating this ImageProcess.
// Note that you need to call ImageProcess.Unwrap() before calling this method if this ImageProcess
// was returned from a transaction, and the transaction was committed or rolled back.
func (ip *ImageProcess) Update() *ImageProcessUpdateOne {
	return NewImageProcessClient(ip.config).UpdateOne(ip)
}

// Unwrap unwraps the ImageProcess entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ip *ImageProcess) Unwrap() *ImageProcess {
	_tx, ok := ip.config.driver.(*txDriver)
	if !ok {
		panic("ent: ImageProcess is not a transactional entity")
	}
	ip.config.driver = _tx.drv
	return ip
}

// String implements the fmt.Stringer.
func (ip *ImageProcess) String() string {
	var builder strings.Builder
	builder.WriteString("ImageProcess(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ip.ID))
	builder.WriteString("image_id=")
	builder.WriteString(fmt.Sprintf("%v", ip.ImageID))
	builder.WriteString(", ")
	builder.WriteString("result_id=")
	builder.WriteString(fmt.Sprintf("%v", ip.ResultID))
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(ip.Kind)
	builder.WriteString(", ")
	if v := ip.FinishedAt; v != nil {
		builder.WriteString("finished_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := ip.ErroredAt; v != nil {
		builder.WriteString("errored_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := ip.ErrorReason; v != nil {
		builder.WriteString("error_reason=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// ImageProcesses is a parsable slice of ImageProcess.
type ImageProcesses []*ImageProcess