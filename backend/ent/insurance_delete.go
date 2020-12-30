// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/b6109868/app/ent/insurance"
	"github.com/b6109868/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// InsuranceDelete is the builder for deleting a Insurance entity.
type InsuranceDelete struct {
	config
	hooks      []Hook
	mutation   *InsuranceMutation
	predicates []predicate.Insurance
}

// Where adds a new predicate to the delete builder.
func (id *InsuranceDelete) Where(ps ...predicate.Insurance) *InsuranceDelete {
	id.predicates = append(id.predicates, ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *InsuranceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(id.hooks) == 0 {
		affected, err = id.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InsuranceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			id.mutation = mutation
			affected, err = id.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(id.hooks) - 1; i >= 0; i-- {
			mut = id.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, id.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (id *InsuranceDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *InsuranceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: insurance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: insurance.FieldID,
			},
		},
	}
	if ps := id.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, id.driver, _spec)
}

// InsuranceDeleteOne is the builder for deleting a single Insurance entity.
type InsuranceDeleteOne struct {
	id *InsuranceDelete
}

// Exec executes the deletion query.
func (ido *InsuranceDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{insurance.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *InsuranceDeleteOne) ExecX(ctx context.Context) {
	ido.id.ExecX(ctx)
}
