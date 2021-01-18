// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/user"
)

// NurseUpdate is the builder for updating Nurse entities.
type NurseUpdate struct {
	config
	hooks      []Hook
	mutation   *NurseMutation
	predicates []predicate.Nurse
}

// Where adds a new predicate for the builder.
func (nu *NurseUpdate) Where(ps ...predicate.Nurse) *NurseUpdate {
	nu.predicates = append(nu.predicates, ps...)
	return nu
}

// SetName sets the name field.
func (nu *NurseUpdate) SetName(s string) *NurseUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetNursinglicense sets the nursinglicense field.
func (nu *NurseUpdate) SetNursinglicense(s string) *NurseUpdate {
	nu.mutation.SetNursinglicense(s)
	return nu
}

// SetPosition sets the position field.
func (nu *NurseUpdate) SetPosition(s string) *NurseUpdate {
	nu.mutation.SetPosition(s)
	return nu
}

// AddEdgesOfHistorytakingIDs adds the EdgesOfHistorytaking edge to Historytaking by ids.
func (nu *NurseUpdate) AddEdgesOfHistorytakingIDs(ids ...int) *NurseUpdate {
	nu.mutation.AddEdgesOfHistorytakingIDs(ids...)
	return nu
}

// AddEdgesOfHistorytaking adds the EdgesOfHistorytaking edges to Historytaking.
func (nu *NurseUpdate) AddEdgesOfHistorytaking(h ...*Historytaking) *NurseUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return nu.AddEdgesOfHistorytakingIDs(ids...)
}

// SetEdgesOfUserID sets the EdgesOfUser edge to User by id.
func (nu *NurseUpdate) SetEdgesOfUserID(id int) *NurseUpdate {
	nu.mutation.SetEdgesOfUserID(id)
	return nu
}

// SetNillableEdgesOfUserID sets the EdgesOfUser edge to User by id if the given value is not nil.
func (nu *NurseUpdate) SetNillableEdgesOfUserID(id *int) *NurseUpdate {
	if id != nil {
		nu = nu.SetEdgesOfUserID(*id)
	}
	return nu
}

// SetEdgesOfUser sets the EdgesOfUser edge to User.
func (nu *NurseUpdate) SetEdgesOfUser(u *User) *NurseUpdate {
	return nu.SetEdgesOfUserID(u.ID)
}

// Mutation returns the NurseMutation object of the builder.
func (nu *NurseUpdate) Mutation() *NurseMutation {
	return nu.mutation
}

// RemoveEdgesOfHistorytakingIDs removes the EdgesOfHistorytaking edge to Historytaking by ids.
func (nu *NurseUpdate) RemoveEdgesOfHistorytakingIDs(ids ...int) *NurseUpdate {
	nu.mutation.RemoveEdgesOfHistorytakingIDs(ids...)
	return nu
}

// RemoveEdgesOfHistorytaking removes EdgesOfHistorytaking edges to Historytaking.
func (nu *NurseUpdate) RemoveEdgesOfHistorytaking(h ...*Historytaking) *NurseUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return nu.RemoveEdgesOfHistorytakingIDs(ids...)
}

// ClearEdgesOfUser clears the EdgesOfUser edge to User.
func (nu *NurseUpdate) ClearEdgesOfUser() *NurseUpdate {
	nu.mutation.ClearEdgesOfUser()
	return nu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (nu *NurseUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NurseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NurseUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NurseUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NurseUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NurseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nurse.Table,
			Columns: nurse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nurse.FieldID,
			},
		},
	}
	if ps := nu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldName,
		})
	}
	if value, ok := nu.mutation.Nursinglicense(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNursinglicense,
		})
	}
	if value, ok := nu.mutation.Position(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldPosition,
		})
	}
	if nodes := nu.mutation.RemovedEdgesOfHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.EdgesOfHistorytakingTable,
			Columns: []string{nurse.EdgesOfHistorytakingColumn},
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
	if nodes := nu.mutation.EdgesOfHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.EdgesOfHistorytakingTable,
			Columns: []string{nurse.EdgesOfHistorytakingColumn},
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
	if nu.mutation.EdgesOfUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nurse.EdgesOfUserTable,
			Columns: []string{nurse.EdgesOfUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.EdgesOfUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nurse.EdgesOfUserTable,
			Columns: []string{nurse.EdgesOfUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nurse.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// NurseUpdateOne is the builder for updating a single Nurse entity.
type NurseUpdateOne struct {
	config
	hooks    []Hook
	mutation *NurseMutation
}

// SetName sets the name field.
func (nuo *NurseUpdateOne) SetName(s string) *NurseUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetNursinglicense sets the nursinglicense field.
func (nuo *NurseUpdateOne) SetNursinglicense(s string) *NurseUpdateOne {
	nuo.mutation.SetNursinglicense(s)
	return nuo
}

// SetPosition sets the position field.
func (nuo *NurseUpdateOne) SetPosition(s string) *NurseUpdateOne {
	nuo.mutation.SetPosition(s)
	return nuo
}

// AddEdgesOfHistorytakingIDs adds the EdgesOfHistorytaking edge to Historytaking by ids.
func (nuo *NurseUpdateOne) AddEdgesOfHistorytakingIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.AddEdgesOfHistorytakingIDs(ids...)
	return nuo
}

// AddEdgesOfHistorytaking adds the EdgesOfHistorytaking edges to Historytaking.
func (nuo *NurseUpdateOne) AddEdgesOfHistorytaking(h ...*Historytaking) *NurseUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return nuo.AddEdgesOfHistorytakingIDs(ids...)
}

// SetEdgesOfUserID sets the EdgesOfUser edge to User by id.
func (nuo *NurseUpdateOne) SetEdgesOfUserID(id int) *NurseUpdateOne {
	nuo.mutation.SetEdgesOfUserID(id)
	return nuo
}

// SetNillableEdgesOfUserID sets the EdgesOfUser edge to User by id if the given value is not nil.
func (nuo *NurseUpdateOne) SetNillableEdgesOfUserID(id *int) *NurseUpdateOne {
	if id != nil {
		nuo = nuo.SetEdgesOfUserID(*id)
	}
	return nuo
}

// SetEdgesOfUser sets the EdgesOfUser edge to User.
func (nuo *NurseUpdateOne) SetEdgesOfUser(u *User) *NurseUpdateOne {
	return nuo.SetEdgesOfUserID(u.ID)
}

// Mutation returns the NurseMutation object of the builder.
func (nuo *NurseUpdateOne) Mutation() *NurseMutation {
	return nuo.mutation
}

// RemoveEdgesOfHistorytakingIDs removes the EdgesOfHistorytaking edge to Historytaking by ids.
func (nuo *NurseUpdateOne) RemoveEdgesOfHistorytakingIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.RemoveEdgesOfHistorytakingIDs(ids...)
	return nuo
}

// RemoveEdgesOfHistorytaking removes EdgesOfHistorytaking edges to Historytaking.
func (nuo *NurseUpdateOne) RemoveEdgesOfHistorytaking(h ...*Historytaking) *NurseUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return nuo.RemoveEdgesOfHistorytakingIDs(ids...)
}

// ClearEdgesOfUser clears the EdgesOfUser edge to User.
func (nuo *NurseUpdateOne) ClearEdgesOfUser() *NurseUpdateOne {
	nuo.mutation.ClearEdgesOfUser()
	return nuo
}

// Save executes the query and returns the updated entity.
func (nuo *NurseUpdateOne) Save(ctx context.Context) (*Nurse, error) {

	var (
		err  error
		node *Nurse
	)
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NurseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			mut = nuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NurseUpdateOne) SaveX(ctx context.Context) *Nurse {
	n, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Exec executes the query on the entity.
func (nuo *NurseUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NurseUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NurseUpdateOne) sqlSave(ctx context.Context) (n *Nurse, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nurse.Table,
			Columns: nurse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nurse.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Nurse.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := nuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldName,
		})
	}
	if value, ok := nuo.mutation.Nursinglicense(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNursinglicense,
		})
	}
	if value, ok := nuo.mutation.Position(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldPosition,
		})
	}
	if nodes := nuo.mutation.RemovedEdgesOfHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.EdgesOfHistorytakingTable,
			Columns: []string{nurse.EdgesOfHistorytakingColumn},
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
	if nodes := nuo.mutation.EdgesOfHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.EdgesOfHistorytakingTable,
			Columns: []string{nurse.EdgesOfHistorytakingColumn},
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
	if nuo.mutation.EdgesOfUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nurse.EdgesOfUserTable,
			Columns: []string{nurse.EdgesOfUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.EdgesOfUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nurse.EdgesOfUserTable,
			Columns: []string{nurse.EdgesOfUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	n = &Nurse{config: nuo.config}
	_spec.Assign = n.assignValues
	_spec.ScanValues = n.scanValues()
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nurse.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return n, nil
}
