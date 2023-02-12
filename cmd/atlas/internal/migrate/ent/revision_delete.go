// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/internal"
	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/predicate"
	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/revision"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RevisionDelete is the builder for deleting a Revision entity.
type RevisionDelete struct {
	config
	hooks    []Hook
	mutation *RevisionMutation
}

// Where appends a list predicates to the RevisionDelete builder.
func (rd *RevisionDelete) Where(ps ...predicate.Revision) *RevisionDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *RevisionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, RevisionMutation](ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *RevisionDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *RevisionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(revision.Table, sqlgraph.NewFieldSpec(revision.FieldID, field.TypeString))
	_spec.Node.Schema = rd.schemaConfig.Revision
	ctx = internal.NewSchemaConfigContext(ctx, rd.schemaConfig)
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rd.mutation.done = true
	return affected, err
}

// RevisionDeleteOne is the builder for deleting a single Revision entity.
type RevisionDeleteOne struct {
	rd *RevisionDelete
}

// Where appends a list predicates to the RevisionDelete builder.
func (rdo *RevisionDeleteOne) Where(ps ...predicate.Revision) *RevisionDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *RevisionDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{revision.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *RevisionDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}
