// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/doctorinfo"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/prename"
)

// PrenameUpdate is the builder for updating Prename entities.
type PrenameUpdate struct {
	config
	hooks      []Hook
	mutation   *PrenameMutation
	predicates []predicate.Prename
}

// Where adds a new predicate for the builder.
func (pu *PrenameUpdate) Where(ps ...predicate.Prename) *PrenameUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPrefix sets the prefix field.
func (pu *PrenameUpdate) SetPrefix(s string) *PrenameUpdate {
	pu.mutation.SetPrefix(s)
	return pu
}

// AddEdgesOfPrename2doctorinfoIDs adds the EdgesOfPrename2doctorinfo edge to Doctorinfo by ids.
func (pu *PrenameUpdate) AddEdgesOfPrename2doctorinfoIDs(ids ...int) *PrenameUpdate {
	pu.mutation.AddEdgesOfPrename2doctorinfoIDs(ids...)
	return pu
}

// AddEdgesOfPrename2doctorinfo adds the EdgesOfPrename2doctorinfo edges to Doctorinfo.
func (pu *PrenameUpdate) AddEdgesOfPrename2doctorinfo(d ...*Doctorinfo) *PrenameUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.AddEdgesOfPrename2doctorinfoIDs(ids...)
}

// AddEdgesOfPatientrecordIDs adds the EdgesOfPatientrecord edge to Patientrecord by ids.
func (pu *PrenameUpdate) AddEdgesOfPatientrecordIDs(ids ...int) *PrenameUpdate {
	pu.mutation.AddEdgesOfPatientrecordIDs(ids...)
	return pu
}

// AddEdgesOfPatientrecord adds the EdgesOfPatientrecord edges to Patientrecord.
func (pu *PrenameUpdate) AddEdgesOfPatientrecord(p ...*Patientrecord) *PrenameUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddEdgesOfPatientrecordIDs(ids...)
}

// Mutation returns the PrenameMutation object of the builder.
func (pu *PrenameUpdate) Mutation() *PrenameMutation {
	return pu.mutation
}

// RemoveEdgesOfPrename2doctorinfoIDs removes the EdgesOfPrename2doctorinfo edge to Doctorinfo by ids.
func (pu *PrenameUpdate) RemoveEdgesOfPrename2doctorinfoIDs(ids ...int) *PrenameUpdate {
	pu.mutation.RemoveEdgesOfPrename2doctorinfoIDs(ids...)
	return pu
}

// RemoveEdgesOfPrename2doctorinfo removes EdgesOfPrename2doctorinfo edges to Doctorinfo.
func (pu *PrenameUpdate) RemoveEdgesOfPrename2doctorinfo(d ...*Doctorinfo) *PrenameUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.RemoveEdgesOfPrename2doctorinfoIDs(ids...)
}

// RemoveEdgesOfPatientrecordIDs removes the EdgesOfPatientrecord edge to Patientrecord by ids.
func (pu *PrenameUpdate) RemoveEdgesOfPatientrecordIDs(ids ...int) *PrenameUpdate {
	pu.mutation.RemoveEdgesOfPatientrecordIDs(ids...)
	return pu
}

// RemoveEdgesOfPatientrecord removes EdgesOfPatientrecord edges to Patientrecord.
func (pu *PrenameUpdate) RemoveEdgesOfPatientrecord(p ...*Patientrecord) *PrenameUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveEdgesOfPatientrecordIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PrenameUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.Prefix(); ok {
		if err := prename.PrefixValidator(v); err != nil {
			return 0, &ValidationError{Name: "prefix", err: fmt.Errorf("ent: validator failed for field \"prefix\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrenameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PrenameUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PrenameUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PrenameUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PrenameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prename.Table,
			Columns: prename.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prename.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Prefix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prename.FieldPrefix,
		})
	}
	if nodes := pu.mutation.RemovedEdgesOfPrename2doctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prename.EdgesOfPrename2doctorinfoTable,
			Columns: []string{prename.EdgesOfPrename2doctorinfoColumn},
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
	if nodes := pu.mutation.EdgesOfPrename2doctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prename.EdgesOfPrename2doctorinfoTable,
			Columns: []string{prename.EdgesOfPrename2doctorinfoColumn},
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
	if nodes := pu.mutation.RemovedEdgesOfPatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prename.EdgesOfPatientrecordTable,
			Columns: []string{prename.EdgesOfPatientrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.EdgesOfPatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prename.EdgesOfPatientrecordTable,
			Columns: []string{prename.EdgesOfPatientrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prename.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PrenameUpdateOne is the builder for updating a single Prename entity.
type PrenameUpdateOne struct {
	config
	hooks    []Hook
	mutation *PrenameMutation
}

// SetPrefix sets the prefix field.
func (puo *PrenameUpdateOne) SetPrefix(s string) *PrenameUpdateOne {
	puo.mutation.SetPrefix(s)
	return puo
}

// AddEdgesOfPrename2doctorinfoIDs adds the EdgesOfPrename2doctorinfo edge to Doctorinfo by ids.
func (puo *PrenameUpdateOne) AddEdgesOfPrename2doctorinfoIDs(ids ...int) *PrenameUpdateOne {
	puo.mutation.AddEdgesOfPrename2doctorinfoIDs(ids...)
	return puo
}

// AddEdgesOfPrename2doctorinfo adds the EdgesOfPrename2doctorinfo edges to Doctorinfo.
func (puo *PrenameUpdateOne) AddEdgesOfPrename2doctorinfo(d ...*Doctorinfo) *PrenameUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.AddEdgesOfPrename2doctorinfoIDs(ids...)
}

// AddEdgesOfPatientrecordIDs adds the EdgesOfPatientrecord edge to Patientrecord by ids.
func (puo *PrenameUpdateOne) AddEdgesOfPatientrecordIDs(ids ...int) *PrenameUpdateOne {
	puo.mutation.AddEdgesOfPatientrecordIDs(ids...)
	return puo
}

// AddEdgesOfPatientrecord adds the EdgesOfPatientrecord edges to Patientrecord.
func (puo *PrenameUpdateOne) AddEdgesOfPatientrecord(p ...*Patientrecord) *PrenameUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddEdgesOfPatientrecordIDs(ids...)
}

// Mutation returns the PrenameMutation object of the builder.
func (puo *PrenameUpdateOne) Mutation() *PrenameMutation {
	return puo.mutation
}

// RemoveEdgesOfPrename2doctorinfoIDs removes the EdgesOfPrename2doctorinfo edge to Doctorinfo by ids.
func (puo *PrenameUpdateOne) RemoveEdgesOfPrename2doctorinfoIDs(ids ...int) *PrenameUpdateOne {
	puo.mutation.RemoveEdgesOfPrename2doctorinfoIDs(ids...)
	return puo
}

// RemoveEdgesOfPrename2doctorinfo removes EdgesOfPrename2doctorinfo edges to Doctorinfo.
func (puo *PrenameUpdateOne) RemoveEdgesOfPrename2doctorinfo(d ...*Doctorinfo) *PrenameUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.RemoveEdgesOfPrename2doctorinfoIDs(ids...)
}

// RemoveEdgesOfPatientrecordIDs removes the EdgesOfPatientrecord edge to Patientrecord by ids.
func (puo *PrenameUpdateOne) RemoveEdgesOfPatientrecordIDs(ids ...int) *PrenameUpdateOne {
	puo.mutation.RemoveEdgesOfPatientrecordIDs(ids...)
	return puo
}

// RemoveEdgesOfPatientrecord removes EdgesOfPatientrecord edges to Patientrecord.
func (puo *PrenameUpdateOne) RemoveEdgesOfPatientrecord(p ...*Patientrecord) *PrenameUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveEdgesOfPatientrecordIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PrenameUpdateOne) Save(ctx context.Context) (*Prename, error) {
	if v, ok := puo.mutation.Prefix(); ok {
		if err := prename.PrefixValidator(v); err != nil {
			return nil, &ValidationError{Name: "prefix", err: fmt.Errorf("ent: validator failed for field \"prefix\": %w", err)}
		}
	}

	var (
		err  error
		node *Prename
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrenameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PrenameUpdateOne) SaveX(ctx context.Context) *Prename {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *PrenameUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PrenameUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PrenameUpdateOne) sqlSave(ctx context.Context) (pr *Prename, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prename.Table,
			Columns: prename.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prename.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Prename.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Prefix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prename.FieldPrefix,
		})
	}
	if nodes := puo.mutation.RemovedEdgesOfPrename2doctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prename.EdgesOfPrename2doctorinfoTable,
			Columns: []string{prename.EdgesOfPrename2doctorinfoColumn},
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
	if nodes := puo.mutation.EdgesOfPrename2doctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prename.EdgesOfPrename2doctorinfoTable,
			Columns: []string{prename.EdgesOfPrename2doctorinfoColumn},
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
	if nodes := puo.mutation.RemovedEdgesOfPatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prename.EdgesOfPatientrecordTable,
			Columns: []string{prename.EdgesOfPatientrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.EdgesOfPatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prename.EdgesOfPatientrecordTable,
			Columns: []string{prename.EdgesOfPatientrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Prename{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prename.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
