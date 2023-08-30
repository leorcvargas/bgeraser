// Code generated by ent, DO NOT EDIT.

package imageprocess

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldLTE(FieldID, id))
}

// ImageID applies equality check predicate on the "image_id" field. It's identical to ImageIDEQ.
func ImageID(v uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldImageID, v))
}

// ResultID applies equality check predicate on the "result_id" field. It's identical to ResultIDEQ.
func ResultID(v uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldResultID, v))
}

// Kind applies equality check predicate on the "kind" field. It's identical to KindEQ.
func Kind(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldKind, v))
}

// FinishedAt applies equality check predicate on the "finished_at" field. It's identical to FinishedAtEQ.
func FinishedAt(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldFinishedAt, v))
}

// ErroredAt applies equality check predicate on the "errored_at" field. It's identical to ErroredAtEQ.
func ErroredAt(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldErroredAt, v))
}

// ErrorReason applies equality check predicate on the "error_reason" field. It's identical to ErrorReasonEQ.
func ErrorReason(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldErrorReason, v))
}

// ImageIDEQ applies the EQ predicate on the "image_id" field.
func ImageIDEQ(v uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldImageID, v))
}

// ImageIDNEQ applies the NEQ predicate on the "image_id" field.
func ImageIDNEQ(v uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNEQ(FieldImageID, v))
}

// ImageIDIn applies the In predicate on the "image_id" field.
func ImageIDIn(vs ...uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIn(FieldImageID, vs...))
}

// ImageIDNotIn applies the NotIn predicate on the "image_id" field.
func ImageIDNotIn(vs ...uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotIn(FieldImageID, vs...))
}

// ResultIDEQ applies the EQ predicate on the "result_id" field.
func ResultIDEQ(v uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldResultID, v))
}

// ResultIDNEQ applies the NEQ predicate on the "result_id" field.
func ResultIDNEQ(v uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNEQ(FieldResultID, v))
}

// ResultIDIn applies the In predicate on the "result_id" field.
func ResultIDIn(vs ...uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIn(FieldResultID, vs...))
}

// ResultIDNotIn applies the NotIn predicate on the "result_id" field.
func ResultIDNotIn(vs ...uuid.UUID) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotIn(FieldResultID, vs...))
}

// ResultIDIsNil applies the IsNil predicate on the "result_id" field.
func ResultIDIsNil() predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIsNull(FieldResultID))
}

// ResultIDNotNil applies the NotNil predicate on the "result_id" field.
func ResultIDNotNil() predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotNull(FieldResultID))
}

// KindEQ applies the EQ predicate on the "kind" field.
func KindEQ(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldKind, v))
}

// KindNEQ applies the NEQ predicate on the "kind" field.
func KindNEQ(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNEQ(FieldKind, v))
}

// KindIn applies the In predicate on the "kind" field.
func KindIn(vs ...string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIn(FieldKind, vs...))
}

// KindNotIn applies the NotIn predicate on the "kind" field.
func KindNotIn(vs ...string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotIn(FieldKind, vs...))
}

// KindGT applies the GT predicate on the "kind" field.
func KindGT(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldGT(FieldKind, v))
}

// KindGTE applies the GTE predicate on the "kind" field.
func KindGTE(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldGTE(FieldKind, v))
}

// KindLT applies the LT predicate on the "kind" field.
func KindLT(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldLT(FieldKind, v))
}

// KindLTE applies the LTE predicate on the "kind" field.
func KindLTE(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldLTE(FieldKind, v))
}

// KindContains applies the Contains predicate on the "kind" field.
func KindContains(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldContains(FieldKind, v))
}

// KindHasPrefix applies the HasPrefix predicate on the "kind" field.
func KindHasPrefix(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldHasPrefix(FieldKind, v))
}

// KindHasSuffix applies the HasSuffix predicate on the "kind" field.
func KindHasSuffix(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldHasSuffix(FieldKind, v))
}

// KindEqualFold applies the EqualFold predicate on the "kind" field.
func KindEqualFold(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEqualFold(FieldKind, v))
}

// KindContainsFold applies the ContainsFold predicate on the "kind" field.
func KindContainsFold(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldContainsFold(FieldKind, v))
}

// FinishedAtEQ applies the EQ predicate on the "finished_at" field.
func FinishedAtEQ(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldFinishedAt, v))
}

// FinishedAtNEQ applies the NEQ predicate on the "finished_at" field.
func FinishedAtNEQ(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNEQ(FieldFinishedAt, v))
}

// FinishedAtIn applies the In predicate on the "finished_at" field.
func FinishedAtIn(vs ...time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIn(FieldFinishedAt, vs...))
}

// FinishedAtNotIn applies the NotIn predicate on the "finished_at" field.
func FinishedAtNotIn(vs ...time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotIn(FieldFinishedAt, vs...))
}

// FinishedAtGT applies the GT predicate on the "finished_at" field.
func FinishedAtGT(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldGT(FieldFinishedAt, v))
}

// FinishedAtGTE applies the GTE predicate on the "finished_at" field.
func FinishedAtGTE(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldGTE(FieldFinishedAt, v))
}

// FinishedAtLT applies the LT predicate on the "finished_at" field.
func FinishedAtLT(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldLT(FieldFinishedAt, v))
}

// FinishedAtLTE applies the LTE predicate on the "finished_at" field.
func FinishedAtLTE(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldLTE(FieldFinishedAt, v))
}

// FinishedAtIsNil applies the IsNil predicate on the "finished_at" field.
func FinishedAtIsNil() predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIsNull(FieldFinishedAt))
}

// FinishedAtNotNil applies the NotNil predicate on the "finished_at" field.
func FinishedAtNotNil() predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotNull(FieldFinishedAt))
}

// ErroredAtEQ applies the EQ predicate on the "errored_at" field.
func ErroredAtEQ(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldErroredAt, v))
}

// ErroredAtNEQ applies the NEQ predicate on the "errored_at" field.
func ErroredAtNEQ(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNEQ(FieldErroredAt, v))
}

// ErroredAtIn applies the In predicate on the "errored_at" field.
func ErroredAtIn(vs ...time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIn(FieldErroredAt, vs...))
}

// ErroredAtNotIn applies the NotIn predicate on the "errored_at" field.
func ErroredAtNotIn(vs ...time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotIn(FieldErroredAt, vs...))
}

// ErroredAtGT applies the GT predicate on the "errored_at" field.
func ErroredAtGT(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldGT(FieldErroredAt, v))
}

// ErroredAtGTE applies the GTE predicate on the "errored_at" field.
func ErroredAtGTE(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldGTE(FieldErroredAt, v))
}

// ErroredAtLT applies the LT predicate on the "errored_at" field.
func ErroredAtLT(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldLT(FieldErroredAt, v))
}

// ErroredAtLTE applies the LTE predicate on the "errored_at" field.
func ErroredAtLTE(v time.Time) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldLTE(FieldErroredAt, v))
}

// ErroredAtIsNil applies the IsNil predicate on the "errored_at" field.
func ErroredAtIsNil() predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIsNull(FieldErroredAt))
}

// ErroredAtNotNil applies the NotNil predicate on the "errored_at" field.
func ErroredAtNotNil() predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotNull(FieldErroredAt))
}

// ErrorReasonEQ applies the EQ predicate on the "error_reason" field.
func ErrorReasonEQ(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEQ(FieldErrorReason, v))
}

// ErrorReasonNEQ applies the NEQ predicate on the "error_reason" field.
func ErrorReasonNEQ(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNEQ(FieldErrorReason, v))
}

// ErrorReasonIn applies the In predicate on the "error_reason" field.
func ErrorReasonIn(vs ...string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIn(FieldErrorReason, vs...))
}

// ErrorReasonNotIn applies the NotIn predicate on the "error_reason" field.
func ErrorReasonNotIn(vs ...string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotIn(FieldErrorReason, vs...))
}

// ErrorReasonGT applies the GT predicate on the "error_reason" field.
func ErrorReasonGT(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldGT(FieldErrorReason, v))
}

// ErrorReasonGTE applies the GTE predicate on the "error_reason" field.
func ErrorReasonGTE(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldGTE(FieldErrorReason, v))
}

// ErrorReasonLT applies the LT predicate on the "error_reason" field.
func ErrorReasonLT(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldLT(FieldErrorReason, v))
}

// ErrorReasonLTE applies the LTE predicate on the "error_reason" field.
func ErrorReasonLTE(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldLTE(FieldErrorReason, v))
}

// ErrorReasonContains applies the Contains predicate on the "error_reason" field.
func ErrorReasonContains(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldContains(FieldErrorReason, v))
}

// ErrorReasonHasPrefix applies the HasPrefix predicate on the "error_reason" field.
func ErrorReasonHasPrefix(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldHasPrefix(FieldErrorReason, v))
}

// ErrorReasonHasSuffix applies the HasSuffix predicate on the "error_reason" field.
func ErrorReasonHasSuffix(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldHasSuffix(FieldErrorReason, v))
}

// ErrorReasonIsNil applies the IsNil predicate on the "error_reason" field.
func ErrorReasonIsNil() predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldIsNull(FieldErrorReason))
}

// ErrorReasonNotNil applies the NotNil predicate on the "error_reason" field.
func ErrorReasonNotNil() predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldNotNull(FieldErrorReason))
}

// ErrorReasonEqualFold applies the EqualFold predicate on the "error_reason" field.
func ErrorReasonEqualFold(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldEqualFold(FieldErrorReason, v))
}

// ErrorReasonContainsFold applies the ContainsFold predicate on the "error_reason" field.
func ErrorReasonContainsFold(v string) predicate.ImageProcess {
	return predicate.ImageProcess(sql.FieldContainsFold(FieldErrorReason, v))
}

// HasOrigin applies the HasEdge predicate on the "origin" edge.
func HasOrigin() predicate.ImageProcess {
	return predicate.ImageProcess(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OriginTable, OriginColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOriginWith applies the HasEdge predicate on the "origin" edge with a given conditions (other predicates).
func HasOriginWith(preds ...predicate.Image) predicate.ImageProcess {
	return predicate.ImageProcess(func(s *sql.Selector) {
		step := newOriginStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResult applies the HasEdge predicate on the "result" edge.
func HasResult() predicate.ImageProcess {
	return predicate.ImageProcess(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ResultTable, ResultColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResultWith applies the HasEdge predicate on the "result" edge with a given conditions (other predicates).
func HasResultWith(preds ...predicate.Image) predicate.ImageProcess {
	return predicate.ImageProcess(func(s *sql.Selector) {
		step := newResultStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ImageProcess) predicate.ImageProcess {
	return predicate.ImageProcess(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ImageProcess) predicate.ImageProcess {
	return predicate.ImageProcess(func(s *sql.Selector) {
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
func Not(p predicate.ImageProcess) predicate.ImageProcess {
	return predicate.ImageProcess(func(s *sql.Selector) {
		p(s.Not())
	})
}