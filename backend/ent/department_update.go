// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/doctorinfo"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/predicate"
)

// DepartmentUpdate is the builder for updating Department entities.
type DepartmentUpdate struct {
	config
	hooks      []Hook
	mutation   *DepartmentMutation
	predicates []predicate.Department
}

// Where adds a new predicate for the builder.
func (du *DepartmentUpdate) Where(ps ...predicate.Department) *DepartmentUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDepartment sets the department field.
func (du *DepartmentUpdate) SetDepartment(s string) *DepartmentUpdate {
	du.mutation.SetDepartment(s)
	return du
}

// AddEdgesOfDepartment2doctorinfoIDs adds the EdgesOfDepartment2doctorinfo edge to Doctorinfo by ids.
func (du *DepartmentUpdate) AddEdgesOfDepartment2doctorinfoIDs(ids ...int) *DepartmentUpdate {
	du.mutation.AddEdgesOfDepartment2doctorinfoIDs(ids...)
	return du
}

// AddEdgesOfDepartment2doctorinfo adds the EdgesOfDepartment2doctorinfo edges to Doctorinfo.
func (du *DepartmentUpdate) AddEdgesOfDepartment2doctorinfo(d ...*Doctorinfo) *DepartmentUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddEdgesOfDepartment2doctorinfoIDs(ids...)
}

// AddEdgesOfHistorytakingIDs adds the EdgesOfHistorytaking edge to Historytaking by ids.
func (du *DepartmentUpdate) AddEdgesOfHistorytakingIDs(ids ...int) *DepartmentUpdate {
	du.mutation.AddEdgesOfHistorytakingIDs(ids...)
	return du
}

// AddEdgesOfHistorytaking adds the EdgesOfHistorytaking edges to Historytaking.
func (du *DepartmentUpdate) AddEdgesOfHistorytaking(h ...*Historytaking) *DepartmentUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return du.AddEdgesOfHistorytakingIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (du *DepartmentUpdate) Mutation() *DepartmentMutation {
	return du.mutation
}

// RemoveEdgesOfDepartment2doctorinfoIDs removes the EdgesOfDepartment2doctorinfo edge to Doctorinfo by ids.
func (du *DepartmentUpdate) RemoveEdgesOfDepartment2doctorinfoIDs(ids ...int) *DepartmentUpdate {
	du.mutation.RemoveEdgesOfDepartment2doctorinfoIDs(ids...)
	return du
}

// RemoveEdgesOfDepartment2doctorinfo removes EdgesOfDepartment2doctorinfo edges to Doctorinfo.
func (du *DepartmentUpdate) RemoveEdgesOfDepartment2doctorinfo(d ...*Doctorinfo) *DepartmentUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveEdgesOfDepartment2doctorinfoIDs(ids...)
}

// RemoveEdgesOfHistorytakingIDs removes the EdgesOfHistorytaking edge to Historytaking by ids.
func (du *DepartmentUpdate) RemoveEdgesOfHistorytakingIDs(ids ...int) *DepartmentUpdate {
	du.mutation.RemoveEdgesOfHistorytakingIDs(ids...)
	return du
}

// RemoveEdgesOfHistorytaking removes EdgesOfHistorytaking edges to Historytaking.
func (du *DepartmentUpdate) RemoveEdgesOfHistorytaking(h ...*Historytaking) *DepartmentUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return du.RemoveEdgesOfHistorytakingIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DepartmentUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Department(); ok {
		if err := department.DepartmentValidator(v); err != nil {
			return 0, &ValidationError{Name: "department", err: fmt.Errorf("ent: validator failed for field \"department\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DepartmentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DepartmentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DepartmentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DepartmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Department(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldDepartment,
		})
	}
	if nodes := du.mutation.RemovedEdgesOfDepartment2doctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EdgesOfDepartment2doctorinfoTable,
			Columns: []string{department.EdgesOfDepartment2doctorinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctorinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.EdgesOfDepartment2doctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EdgesOfDepartment2doctorinfoTable,
			Columns: []string{department.EdgesOfDepartment2doctorinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctorinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := du.mutation.RemovedEdgesOfHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EdgesOfHistorytakingTable,
			Columns: []string{department.EdgesOfHistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.EdgesOfHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EdgesOfHistorytakingTable,
			Columns: []string{department.EdgesOfHistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DepartmentUpdateOne is the builder for updating a single Department entity.
type DepartmentUpdateOne struct {
	config
	hooks    []Hook
	mutation *DepartmentMutation
}

// SetDepartment sets the department field.
func (duo *DepartmentUpdateOne) SetDepartment(s string) *DepartmentUpdateOne {
	duo.mutation.SetDepartment(s)
	return duo
}

// AddEdgesOfDepartment2doctorinfoIDs adds the EdgesOfDepartment2doctorinfo edge to Doctorinfo by ids.
func (duo *DepartmentUpdateOne) AddEdgesOfDepartment2doctorinfoIDs(ids ...int) *DepartmentUpdateOne {
	duo.mutation.AddEdgesOfDepartment2doctorinfoIDs(ids...)
	return duo
}

// AddEdgesOfDepartment2doctorinfo adds the EdgesOfDepartment2doctorinfo edges to Doctorinfo.
func (duo *DepartmentUpdateOne) AddEdgesOfDepartment2doctorinfo(d ...*Doctorinfo) *DepartmentUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddEdgesOfDepartment2doctorinfoIDs(ids...)
}

// AddEdgesOfHistorytakingIDs adds the EdgesOfHistorytaking edge to Historytaking by ids.
func (duo *DepartmentUpdateOne) AddEdgesOfHistorytakingIDs(ids ...int) *DepartmentUpdateOne {
	duo.mutation.AddEdgesOfHistorytakingIDs(ids...)
	return duo
}

// AddEdgesOfHistorytaking adds the EdgesOfHistorytaking edges to Historytaking.
func (duo *DepartmentUpdateOne) AddEdgesOfHistorytaking(h ...*Historytaking) *DepartmentUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return duo.AddEdgesOfHistorytakingIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (duo *DepartmentUpdateOne) Mutation() *DepartmentMutation {
	return duo.mutation
}

// RemoveEdgesOfDepartment2doctorinfoIDs removes the EdgesOfDepartment2doctorinfo edge to Doctorinfo by ids.
func (duo *DepartmentUpdateOne) RemoveEdgesOfDepartment2doctorinfoIDs(ids ...int) *DepartmentUpdateOne {
	duo.mutation.RemoveEdgesOfDepartment2doctorinfoIDs(ids...)
	return duo
}

// RemoveEdgesOfDepartment2doctorinfo removes EdgesOfDepartment2doctorinfo edges to Doctorinfo.
func (duo *DepartmentUpdateOne) RemoveEdgesOfDepartment2doctorinfo(d ...*Doctorinfo) *DepartmentUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveEdgesOfDepartment2doctorinfoIDs(ids...)
}

// RemoveEdgesOfHistorytakingIDs removes the EdgesOfHistorytaking edge to Historytaking by ids.
func (duo *DepartmentUpdateOne) RemoveEdgesOfHistorytakingIDs(ids ...int) *DepartmentUpdateOne {
	duo.mutation.RemoveEdgesOfHistorytakingIDs(ids...)
	return duo
}

// RemoveEdgesOfHistorytaking removes EdgesOfHistorytaking edges to Historytaking.
func (duo *DepartmentUpdateOne) RemoveEdgesOfHistorytaking(h ...*Historytaking) *DepartmentUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return duo.RemoveEdgesOfHistorytakingIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DepartmentUpdateOne) Save(ctx context.Context) (*Department, error) {
	if v, ok := duo.mutation.Department(); ok {
		if err := department.DepartmentValidator(v); err != nil {
			return nil, &ValidationError{Name: "department", err: fmt.Errorf("ent: validator failed for field \"department\": %w", err)}
		}
	}

	var (
		err  error
		node *Department
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DepartmentUpdateOne) SaveX(ctx context.Context) *Department {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DepartmentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DepartmentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DepartmentUpdateOne) sqlSave(ctx context.Context) (d *Department, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Department.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Department(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldDepartment,
		})
	}
	if nodes := duo.mutation.RemovedEdgesOfDepartment2doctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EdgesOfDepartment2doctorinfoTable,
			Columns: []string{department.EdgesOfDepartment2doctorinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctorinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.EdgesOfDepartment2doctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EdgesOfDepartment2doctorinfoTable,
			Columns: []string{department.EdgesOfDepartment2doctorinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctorinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := duo.mutation.RemovedEdgesOfHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EdgesOfHistorytakingTable,
			Columns: []string{department.EdgesOfHistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.EdgesOfHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EdgesOfHistorytakingTable,
			Columns: []string{department.EdgesOfHistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Department{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
