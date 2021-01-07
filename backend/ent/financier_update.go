// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/bill"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/user"
)

// FinancierUpdate is the builder for updating Financier entities.
type FinancierUpdate struct {
	config
	hooks      []Hook
	mutation   *FinancierMutation
	predicates []predicate.Financier
}

// Where adds a new predicate for the builder.
func (fu *FinancierUpdate) Where(ps ...predicate.Financier) *FinancierUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetName sets the name field.
func (fu *FinancierUpdate) SetName(s string) *FinancierUpdate {
	fu.mutation.SetName(s)
	return fu
}

// AddBillIDs adds the bills edge to Bill by ids.
func (fu *FinancierUpdate) AddBillIDs(ids ...int) *FinancierUpdate {
	fu.mutation.AddBillIDs(ids...)
	return fu
}

// AddBills adds the bills edges to Bill.
func (fu *FinancierUpdate) AddBills(b ...*Bill) *FinancierUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return fu.AddBillIDs(ids...)
}

// SetUserID sets the user edge to User by id.
func (fu *FinancierUpdate) SetUserID(id int) *FinancierUpdate {
	fu.mutation.SetUserID(id)
	return fu
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (fu *FinancierUpdate) SetNillableUserID(id *int) *FinancierUpdate {
	if id != nil {
		fu = fu.SetUserID(*id)
	}
	return fu
}

// SetUser sets the user edge to User.
func (fu *FinancierUpdate) SetUser(u *User) *FinancierUpdate {
	return fu.SetUserID(u.ID)
}

// Mutation returns the FinancierMutation object of the builder.
func (fu *FinancierUpdate) Mutation() *FinancierMutation {
	return fu.mutation
}

// RemoveBillIDs removes the bills edge to Bill by ids.
func (fu *FinancierUpdate) RemoveBillIDs(ids ...int) *FinancierUpdate {
	fu.mutation.RemoveBillIDs(ids...)
	return fu
}

// RemoveBills removes bills edges to Bill.
func (fu *FinancierUpdate) RemoveBills(b ...*Bill) *FinancierUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return fu.RemoveBillIDs(ids...)
}

// ClearUser clears the user edge to User.
func (fu *FinancierUpdate) ClearUser() *FinancierUpdate {
	fu.mutation.ClearUser()
	return fu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FinancierUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := fu.mutation.Name(); ok {
		if err := financier.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FinancierMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FinancierUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FinancierUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FinancierUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FinancierUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   financier.Table,
			Columns: financier.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: financier.FieldID,
			},
		},
	}
	if ps := fu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: financier.FieldName,
		})
	}
	if nodes := fu.mutation.RemovedBillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   financier.BillsTable,
			Columns: []string{financier.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   financier.BillsTable,
			Columns: []string{financier.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   financier.UserTable,
			Columns: []string{financier.UserColumn},
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
	if nodes := fu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   financier.UserTable,
			Columns: []string{financier.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{financier.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FinancierUpdateOne is the builder for updating a single Financier entity.
type FinancierUpdateOne struct {
	config
	hooks    []Hook
	mutation *FinancierMutation
}

// SetName sets the name field.
func (fuo *FinancierUpdateOne) SetName(s string) *FinancierUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// AddBillIDs adds the bills edge to Bill by ids.
func (fuo *FinancierUpdateOne) AddBillIDs(ids ...int) *FinancierUpdateOne {
	fuo.mutation.AddBillIDs(ids...)
	return fuo
}

// AddBills adds the bills edges to Bill.
func (fuo *FinancierUpdateOne) AddBills(b ...*Bill) *FinancierUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return fuo.AddBillIDs(ids...)
}

// SetUserID sets the user edge to User by id.
func (fuo *FinancierUpdateOne) SetUserID(id int) *FinancierUpdateOne {
	fuo.mutation.SetUserID(id)
	return fuo
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (fuo *FinancierUpdateOne) SetNillableUserID(id *int) *FinancierUpdateOne {
	if id != nil {
		fuo = fuo.SetUserID(*id)
	}
	return fuo
}

// SetUser sets the user edge to User.
func (fuo *FinancierUpdateOne) SetUser(u *User) *FinancierUpdateOne {
	return fuo.SetUserID(u.ID)
}

// Mutation returns the FinancierMutation object of the builder.
func (fuo *FinancierUpdateOne) Mutation() *FinancierMutation {
	return fuo.mutation
}

// RemoveBillIDs removes the bills edge to Bill by ids.
func (fuo *FinancierUpdateOne) RemoveBillIDs(ids ...int) *FinancierUpdateOne {
	fuo.mutation.RemoveBillIDs(ids...)
	return fuo
}

// RemoveBills removes bills edges to Bill.
func (fuo *FinancierUpdateOne) RemoveBills(b ...*Bill) *FinancierUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return fuo.RemoveBillIDs(ids...)
}

// ClearUser clears the user edge to User.
func (fuo *FinancierUpdateOne) ClearUser() *FinancierUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// Save executes the query and returns the updated entity.
func (fuo *FinancierUpdateOne) Save(ctx context.Context) (*Financier, error) {
	if v, ok := fuo.mutation.Name(); ok {
		if err := financier.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err  error
		node *Financier
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FinancierMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FinancierUpdateOne) SaveX(ctx context.Context) *Financier {
	f, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// Exec executes the query on the entity.
func (fuo *FinancierUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FinancierUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FinancierUpdateOne) sqlSave(ctx context.Context) (f *Financier, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   financier.Table,
			Columns: financier.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: financier.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Financier.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := fuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: financier.FieldName,
		})
	}
	if nodes := fuo.mutation.RemovedBillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   financier.BillsTable,
			Columns: []string{financier.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   financier.BillsTable,
			Columns: []string{financier.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   financier.UserTable,
			Columns: []string{financier.UserColumn},
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
	if nodes := fuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   financier.UserTable,
			Columns: []string{financier.UserColumn},
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
	f = &Financier{config: fuo.config}
	_spec.Assign = f.assignValues
	_spec.ScanValues = f.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{financier.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}
