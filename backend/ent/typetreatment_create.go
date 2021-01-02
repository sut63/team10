// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/treatment"
	"github.com/team10/app/ent/typetreatment"
)

// TypetreatmentCreate is the builder for creating a Typetreatment entity.
type TypetreatmentCreate struct {
	config
	mutation *TypetreatmentMutation
	hooks    []Hook
}

// SetType sets the type field.
func (tc *TypetreatmentCreate) SetType(s string) *TypetreatmentCreate {
	tc.mutation.SetType(s)
	return tc
}

// AddTreatmentIDs adds the treatment edge to Treatment by ids.
func (tc *TypetreatmentCreate) AddTreatmentIDs(ids ...int) *TypetreatmentCreate {
	tc.mutation.AddTreatmentIDs(ids...)
	return tc
}

// AddTreatment adds the treatment edges to Treatment.
func (tc *TypetreatmentCreate) AddTreatment(t ...*Treatment) *TypetreatmentCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTreatmentIDs(ids...)
}

// Mutation returns the TypetreatmentMutation object of the builder.
func (tc *TypetreatmentCreate) Mutation() *TypetreatmentMutation {
	return tc.mutation
}

// Save creates the Typetreatment in the database.
func (tc *TypetreatmentCreate) Save(ctx context.Context) (*Typetreatment, error) {
	if _, ok := tc.mutation.GetType(); !ok {
		return nil, &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	if v, ok := tc.mutation.GetType(); ok {
		if err := typetreatment.TypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	var (
		err  error
		node *Typetreatment
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypetreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TypetreatmentCreate) SaveX(ctx context.Context) *Typetreatment {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TypetreatmentCreate) sqlSave(ctx context.Context) (*Typetreatment, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TypetreatmentCreate) createSpec() (*Typetreatment, *sqlgraph.CreateSpec) {
	var (
		t     = &Typetreatment{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: typetreatment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typetreatment.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typetreatment.FieldType,
		})
		t.Type = value
	}
	if nodes := tc.mutation.TreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   typetreatment.TreatmentTable,
			Columns: []string{typetreatment.TreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: treatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return t, _spec
}
