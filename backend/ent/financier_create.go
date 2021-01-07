// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/bill"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/user"
)

// FinancierCreate is the builder for creating a Financier entity.
type FinancierCreate struct {
	config
	mutation *FinancierMutation
	hooks    []Hook
}

// SetName sets the name field.
func (fc *FinancierCreate) SetName(s string) *FinancierCreate {
	fc.mutation.SetName(s)
	return fc
}

// AddBillIDs adds the bills edge to Bill by ids.
func (fc *FinancierCreate) AddBillIDs(ids ...int) *FinancierCreate {
	fc.mutation.AddBillIDs(ids...)
	return fc
}

// AddBills adds the bills edges to Bill.
func (fc *FinancierCreate) AddBills(b ...*Bill) *FinancierCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return fc.AddBillIDs(ids...)
}

// SetUserID sets the user edge to User by id.
func (fc *FinancierCreate) SetUserID(id int) *FinancierCreate {
	fc.mutation.SetUserID(id)
	return fc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (fc *FinancierCreate) SetNillableUserID(id *int) *FinancierCreate {
	if id != nil {
		fc = fc.SetUserID(*id)
	}
	return fc
}

// SetUser sets the user edge to User.
func (fc *FinancierCreate) SetUser(u *User) *FinancierCreate {
	return fc.SetUserID(u.ID)
}

// Mutation returns the FinancierMutation object of the builder.
func (fc *FinancierCreate) Mutation() *FinancierMutation {
	return fc.mutation
}

// Save creates the Financier in the database.
func (fc *FinancierCreate) Save(ctx context.Context) (*Financier, error) {
	if _, ok := fc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := fc.mutation.Name(); ok {
		if err := financier.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	var (
		err  error
		node *Financier
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FinancierMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FinancierCreate) SaveX(ctx context.Context) *Financier {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FinancierCreate) sqlSave(ctx context.Context) (*Financier, error) {
	f, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	f.ID = int(id)
	return f, nil
}

func (fc *FinancierCreate) createSpec() (*Financier, *sqlgraph.CreateSpec) {
	var (
		f     = &Financier{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: financier.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: financier.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: financier.FieldName,
		})
		f.Name = value
	}
	if nodes := fc.mutation.BillsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return f, _spec
}
