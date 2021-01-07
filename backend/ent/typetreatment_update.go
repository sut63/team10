// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/treatment"
	"github.com/team10/app/ent/typetreatment"
)

// TypetreatmentUpdate is the builder for updating Typetreatment entities.
type TypetreatmentUpdate struct {
	config
	hooks      []Hook
	mutation   *TypetreatmentMutation
	predicates []predicate.Typetreatment
}

// Where adds a new predicate for the builder.
func (tu *TypetreatmentUpdate) Where(ps ...predicate.Typetreatment) *TypetreatmentUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetTypetreatment sets the Typetreatment field.
func (tu *TypetreatmentUpdate) SetTypetreatment(s string) *TypetreatmentUpdate {
	tu.mutation.SetTypetreatment(s)
	return tu
}

// AddTreatmentIDs adds the treatment edge to Treatment by ids.
func (tu *TypetreatmentUpdate) AddTreatmentIDs(ids ...int) *TypetreatmentUpdate {
	tu.mutation.AddTreatmentIDs(ids...)
	return tu
}

// AddTreatment adds the treatment edges to Treatment.
func (tu *TypetreatmentUpdate) AddTreatment(t ...*Treatment) *TypetreatmentUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTreatmentIDs(ids...)
}

// Mutation returns the TypetreatmentMutation object of the builder.
func (tu *TypetreatmentUpdate) Mutation() *TypetreatmentMutation {
	return tu.mutation
}

// RemoveTreatmentIDs removes the treatment edge to Treatment by ids.
func (tu *TypetreatmentUpdate) RemoveTreatmentIDs(ids ...int) *TypetreatmentUpdate {
	tu.mutation.RemoveTreatmentIDs(ids...)
	return tu
}

// RemoveTreatment removes treatment edges to Treatment.
func (tu *TypetreatmentUpdate) RemoveTreatment(t ...*Treatment) *TypetreatmentUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTreatmentIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TypetreatmentUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := tu.mutation.Typetreatment(); ok {
		if err := typetreatment.TypetreatmentValidator(v); err != nil {
			return 0, &ValidationError{Name: "Typetreatment", err: fmt.Errorf("ent: validator failed for field \"Typetreatment\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypetreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TypetreatmentUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TypetreatmentUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TypetreatmentUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TypetreatmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   typetreatment.Table,
			Columns: typetreatment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typetreatment.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Typetreatment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typetreatment.FieldTypetreatment,
		})
	}
	if nodes := tu.mutation.RemovedTreatmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TreatmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{typetreatment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TypetreatmentUpdateOne is the builder for updating a single Typetreatment entity.
type TypetreatmentUpdateOne struct {
	config
	hooks    []Hook
	mutation *TypetreatmentMutation
}

// SetTypetreatment sets the Typetreatment field.
func (tuo *TypetreatmentUpdateOne) SetTypetreatment(s string) *TypetreatmentUpdateOne {
	tuo.mutation.SetTypetreatment(s)
	return tuo
}

// AddTreatmentIDs adds the treatment edge to Treatment by ids.
func (tuo *TypetreatmentUpdateOne) AddTreatmentIDs(ids ...int) *TypetreatmentUpdateOne {
	tuo.mutation.AddTreatmentIDs(ids...)
	return tuo
}

// AddTreatment adds the treatment edges to Treatment.
func (tuo *TypetreatmentUpdateOne) AddTreatment(t ...*Treatment) *TypetreatmentUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTreatmentIDs(ids...)
}

// Mutation returns the TypetreatmentMutation object of the builder.
func (tuo *TypetreatmentUpdateOne) Mutation() *TypetreatmentMutation {
	return tuo.mutation
}

// RemoveTreatmentIDs removes the treatment edge to Treatment by ids.
func (tuo *TypetreatmentUpdateOne) RemoveTreatmentIDs(ids ...int) *TypetreatmentUpdateOne {
	tuo.mutation.RemoveTreatmentIDs(ids...)
	return tuo
}

// RemoveTreatment removes treatment edges to Treatment.
func (tuo *TypetreatmentUpdateOne) RemoveTreatment(t ...*Treatment) *TypetreatmentUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTreatmentIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (tuo *TypetreatmentUpdateOne) Save(ctx context.Context) (*Typetreatment, error) {
	if v, ok := tuo.mutation.Typetreatment(); ok {
		if err := typetreatment.TypetreatmentValidator(v); err != nil {
			return nil, &ValidationError{Name: "Typetreatment", err: fmt.Errorf("ent: validator failed for field \"Typetreatment\": %w", err)}
		}
	}

	var (
		err  error
		node *Typetreatment
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypetreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TypetreatmentUpdateOne) SaveX(ctx context.Context) *Typetreatment {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TypetreatmentUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TypetreatmentUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TypetreatmentUpdateOne) sqlSave(ctx context.Context) (t *Typetreatment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   typetreatment.Table,
			Columns: typetreatment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typetreatment.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Typetreatment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := tuo.mutation.Typetreatment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typetreatment.FieldTypetreatment,
		})
	}
	if nodes := tuo.mutation.RemovedTreatmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TreatmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	t = &Typetreatment{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{typetreatment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
