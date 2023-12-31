// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/ent/image"
	"github.com/leorcvargas/bgeraser/ent/imageprocess"
)

// ImageProcessCreate is the builder for creating a ImageProcess entity.
type ImageProcessCreate struct {
	config
	mutation *ImageProcessMutation
	hooks    []Hook
}

// SetImageID sets the "image_id" field.
func (ipc *ImageProcessCreate) SetImageID(u uuid.UUID) *ImageProcessCreate {
	ipc.mutation.SetImageID(u)
	return ipc
}

// SetResultID sets the "result_id" field.
func (ipc *ImageProcessCreate) SetResultID(u uuid.UUID) *ImageProcessCreate {
	ipc.mutation.SetResultID(u)
	return ipc
}

// SetNillableResultID sets the "result_id" field if the given value is not nil.
func (ipc *ImageProcessCreate) SetNillableResultID(u *uuid.UUID) *ImageProcessCreate {
	if u != nil {
		ipc.SetResultID(*u)
	}
	return ipc
}

// SetKind sets the "kind" field.
func (ipc *ImageProcessCreate) SetKind(s string) *ImageProcessCreate {
	ipc.mutation.SetKind(s)
	return ipc
}

// SetFinishedAt sets the "finished_at" field.
func (ipc *ImageProcessCreate) SetFinishedAt(t time.Time) *ImageProcessCreate {
	ipc.mutation.SetFinishedAt(t)
	return ipc
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (ipc *ImageProcessCreate) SetNillableFinishedAt(t *time.Time) *ImageProcessCreate {
	if t != nil {
		ipc.SetFinishedAt(*t)
	}
	return ipc
}

// SetErroredAt sets the "errored_at" field.
func (ipc *ImageProcessCreate) SetErroredAt(t time.Time) *ImageProcessCreate {
	ipc.mutation.SetErroredAt(t)
	return ipc
}

// SetNillableErroredAt sets the "errored_at" field if the given value is not nil.
func (ipc *ImageProcessCreate) SetNillableErroredAt(t *time.Time) *ImageProcessCreate {
	if t != nil {
		ipc.SetErroredAt(*t)
	}
	return ipc
}

// SetErrorReason sets the "error_reason" field.
func (ipc *ImageProcessCreate) SetErrorReason(s string) *ImageProcessCreate {
	ipc.mutation.SetErrorReason(s)
	return ipc
}

// SetNillableErrorReason sets the "error_reason" field if the given value is not nil.
func (ipc *ImageProcessCreate) SetNillableErrorReason(s *string) *ImageProcessCreate {
	if s != nil {
		ipc.SetErrorReason(*s)
	}
	return ipc
}

// SetID sets the "id" field.
func (ipc *ImageProcessCreate) SetID(u uuid.UUID) *ImageProcessCreate {
	ipc.mutation.SetID(u)
	return ipc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ipc *ImageProcessCreate) SetNillableID(u *uuid.UUID) *ImageProcessCreate {
	if u != nil {
		ipc.SetID(*u)
	}
	return ipc
}

// SetOriginID sets the "origin" edge to the Image entity by ID.
func (ipc *ImageProcessCreate) SetOriginID(id uuid.UUID) *ImageProcessCreate {
	ipc.mutation.SetOriginID(id)
	return ipc
}

// SetOrigin sets the "origin" edge to the Image entity.
func (ipc *ImageProcessCreate) SetOrigin(i *Image) *ImageProcessCreate {
	return ipc.SetOriginID(i.ID)
}

// SetResult sets the "result" edge to the Image entity.
func (ipc *ImageProcessCreate) SetResult(i *Image) *ImageProcessCreate {
	return ipc.SetResultID(i.ID)
}

// Mutation returns the ImageProcessMutation object of the builder.
func (ipc *ImageProcessCreate) Mutation() *ImageProcessMutation {
	return ipc.mutation
}

// Save creates the ImageProcess in the database.
func (ipc *ImageProcessCreate) Save(ctx context.Context) (*ImageProcess, error) {
	ipc.defaults()
	return withHooks(ctx, ipc.sqlSave, ipc.mutation, ipc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ipc *ImageProcessCreate) SaveX(ctx context.Context) *ImageProcess {
	v, err := ipc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ipc *ImageProcessCreate) Exec(ctx context.Context) error {
	_, err := ipc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipc *ImageProcessCreate) ExecX(ctx context.Context) {
	if err := ipc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ipc *ImageProcessCreate) defaults() {
	if _, ok := ipc.mutation.ID(); !ok {
		v := imageprocess.DefaultID()
		ipc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ipc *ImageProcessCreate) check() error {
	if _, ok := ipc.mutation.ImageID(); !ok {
		return &ValidationError{Name: "image_id", err: errors.New(`ent: missing required field "ImageProcess.image_id"`)}
	}
	if _, ok := ipc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "ImageProcess.kind"`)}
	}
	if v, ok := ipc.mutation.Kind(); ok {
		if err := imageprocess.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "ImageProcess.kind": %w`, err)}
		}
	}
	if _, ok := ipc.mutation.OriginID(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required edge "ImageProcess.origin"`)}
	}
	return nil
}

func (ipc *ImageProcessCreate) sqlSave(ctx context.Context) (*ImageProcess, error) {
	if err := ipc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ipc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ipc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ipc.mutation.id = &_node.ID
	ipc.mutation.done = true
	return _node, nil
}

func (ipc *ImageProcessCreate) createSpec() (*ImageProcess, *sqlgraph.CreateSpec) {
	var (
		_node = &ImageProcess{config: ipc.config}
		_spec = sqlgraph.NewCreateSpec(imageprocess.Table, sqlgraph.NewFieldSpec(imageprocess.FieldID, field.TypeUUID))
	)
	if id, ok := ipc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ipc.mutation.Kind(); ok {
		_spec.SetField(imageprocess.FieldKind, field.TypeString, value)
		_node.Kind = value
	}
	if value, ok := ipc.mutation.FinishedAt(); ok {
		_spec.SetField(imageprocess.FieldFinishedAt, field.TypeTime, value)
		_node.FinishedAt = &value
	}
	if value, ok := ipc.mutation.ErroredAt(); ok {
		_spec.SetField(imageprocess.FieldErroredAt, field.TypeTime, value)
		_node.ErroredAt = &value
	}
	if value, ok := ipc.mutation.ErrorReason(); ok {
		_spec.SetField(imageprocess.FieldErrorReason, field.TypeString, value)
		_node.ErrorReason = &value
	}
	if nodes := ipc.mutation.OriginIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   imageprocess.OriginTable,
			Columns: []string{imageprocess.OriginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ImageID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ipc.mutation.ResultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   imageprocess.ResultTable,
			Columns: []string{imageprocess.ResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ResultID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ImageProcessCreateBulk is the builder for creating many ImageProcess entities in bulk.
type ImageProcessCreateBulk struct {
	config
	builders []*ImageProcessCreate
}

// Save creates the ImageProcess entities in the database.
func (ipcb *ImageProcessCreateBulk) Save(ctx context.Context) ([]*ImageProcess, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ipcb.builders))
	nodes := make([]*ImageProcess, len(ipcb.builders))
	mutators := make([]Mutator, len(ipcb.builders))
	for i := range ipcb.builders {
		func(i int, root context.Context) {
			builder := ipcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImageProcessMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ipcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ipcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ipcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ipcb *ImageProcessCreateBulk) SaveX(ctx context.Context) []*ImageProcess {
	v, err := ipcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ipcb *ImageProcessCreateBulk) Exec(ctx context.Context) error {
	_, err := ipcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipcb *ImageProcessCreateBulk) ExecX(ctx context.Context) {
	if err := ipcb.Exec(ctx); err != nil {
		panic(err)
	}
}
