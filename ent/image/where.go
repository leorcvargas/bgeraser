// Code generated by ent, DO NOT EDIT.

package image

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldID, id))
}

// Format applies equality check predicate on the "format" field. It's identical to FormatEQ.
func Format(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldFormat, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int64) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldSize, v))
}

// OriginalFilename applies equality check predicate on the "original_filename" field. It's identical to OriginalFilenameEQ.
func OriginalFilename(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldOriginalFilename, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldDeletedAt, v))
}

// FormatEQ applies the EQ predicate on the "format" field.
func FormatEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldFormat, v))
}

// FormatNEQ applies the NEQ predicate on the "format" field.
func FormatNEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldFormat, v))
}

// FormatIn applies the In predicate on the "format" field.
func FormatIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldFormat, vs...))
}

// FormatNotIn applies the NotIn predicate on the "format" field.
func FormatNotIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldFormat, vs...))
}

// FormatGT applies the GT predicate on the "format" field.
func FormatGT(v string) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldFormat, v))
}

// FormatGTE applies the GTE predicate on the "format" field.
func FormatGTE(v string) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldFormat, v))
}

// FormatLT applies the LT predicate on the "format" field.
func FormatLT(v string) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldFormat, v))
}

// FormatLTE applies the LTE predicate on the "format" field.
func FormatLTE(v string) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldFormat, v))
}

// FormatContains applies the Contains predicate on the "format" field.
func FormatContains(v string) predicate.Image {
	return predicate.Image(sql.FieldContains(FieldFormat, v))
}

// FormatHasPrefix applies the HasPrefix predicate on the "format" field.
func FormatHasPrefix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasPrefix(FieldFormat, v))
}

// FormatHasSuffix applies the HasSuffix predicate on the "format" field.
func FormatHasSuffix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasSuffix(FieldFormat, v))
}

// FormatEqualFold applies the EqualFold predicate on the "format" field.
func FormatEqualFold(v string) predicate.Image {
	return predicate.Image(sql.FieldEqualFold(FieldFormat, v))
}

// FormatContainsFold applies the ContainsFold predicate on the "format" field.
func FormatContainsFold(v string) predicate.Image {
	return predicate.Image(sql.FieldContainsFold(FieldFormat, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int64) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int64) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int64) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int64) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int64) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int64) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int64) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int64) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldSize, v))
}

// OriginalFilenameEQ applies the EQ predicate on the "original_filename" field.
func OriginalFilenameEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldOriginalFilename, v))
}

// OriginalFilenameNEQ applies the NEQ predicate on the "original_filename" field.
func OriginalFilenameNEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldOriginalFilename, v))
}

// OriginalFilenameIn applies the In predicate on the "original_filename" field.
func OriginalFilenameIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldOriginalFilename, vs...))
}

// OriginalFilenameNotIn applies the NotIn predicate on the "original_filename" field.
func OriginalFilenameNotIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldOriginalFilename, vs...))
}

// OriginalFilenameGT applies the GT predicate on the "original_filename" field.
func OriginalFilenameGT(v string) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldOriginalFilename, v))
}

// OriginalFilenameGTE applies the GTE predicate on the "original_filename" field.
func OriginalFilenameGTE(v string) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldOriginalFilename, v))
}

// OriginalFilenameLT applies the LT predicate on the "original_filename" field.
func OriginalFilenameLT(v string) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldOriginalFilename, v))
}

// OriginalFilenameLTE applies the LTE predicate on the "original_filename" field.
func OriginalFilenameLTE(v string) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldOriginalFilename, v))
}

// OriginalFilenameContains applies the Contains predicate on the "original_filename" field.
func OriginalFilenameContains(v string) predicate.Image {
	return predicate.Image(sql.FieldContains(FieldOriginalFilename, v))
}

// OriginalFilenameHasPrefix applies the HasPrefix predicate on the "original_filename" field.
func OriginalFilenameHasPrefix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasPrefix(FieldOriginalFilename, v))
}

// OriginalFilenameHasSuffix applies the HasSuffix predicate on the "original_filename" field.
func OriginalFilenameHasSuffix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasSuffix(FieldOriginalFilename, v))
}

// OriginalFilenameEqualFold applies the EqualFold predicate on the "original_filename" field.
func OriginalFilenameEqualFold(v string) predicate.Image {
	return predicate.Image(sql.FieldEqualFold(FieldOriginalFilename, v))
}

// OriginalFilenameContainsFold applies the ContainsFold predicate on the "original_filename" field.
func OriginalFilenameContainsFold(v string) predicate.Image {
	return predicate.Image(sql.FieldContainsFold(FieldOriginalFilename, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Image {
	return predicate.Image(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Image {
	return predicate.Image(sql.FieldNotNull(FieldDeletedAt))
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ImagesTable, ImagesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newImagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasImageProcesses applies the HasEdge predicate on the "image_processes" edge.
func HasImageProcesses() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ImageProcessesTable, ImageProcessesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImageProcessesWith applies the HasEdge predicate on the "image_processes" edge with a given conditions (other predicates).
func HasImageProcessesWith(preds ...predicate.ImageProcess) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newImageProcessesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		p(s.Not())
	})
}