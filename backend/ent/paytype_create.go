// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/b6109868/app/ent/bill"
	"github.com/b6109868/app/ent/paytype"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PaytypeCreate is the builder for creating a Paytype entity.
type PaytypeCreate struct {
	config
	mutation *PaytypeMutation
	hooks    []Hook
}

// SetPaytype sets the paytype field.
func (pc *PaytypeCreate) SetPaytype(s string) *PaytypeCreate {
	pc.mutation.SetPaytype(s)
	return pc
}

// AddBillIDs adds the bills edge to Bill by ids.
func (pc *PaytypeCreate) AddBillIDs(ids ...int) *PaytypeCreate {
	pc.mutation.AddBillIDs(ids...)
	return pc
}

// AddBills adds the bills edges to Bill.
func (pc *PaytypeCreate) AddBills(b ...*Bill) *PaytypeCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddBillIDs(ids...)
}

// Mutation returns the PaytypeMutation object of the builder.
func (pc *PaytypeCreate) Mutation() *PaytypeMutation {
	return pc.mutation
}

// Save creates the Paytype in the database.
func (pc *PaytypeCreate) Save(ctx context.Context) (*Paytype, error) {
	if _, ok := pc.mutation.Paytype(); !ok {
		return nil, &ValidationError{Name: "paytype", err: errors.New("ent: missing required field \"paytype\"")}
	}
	if v, ok := pc.mutation.Paytype(); ok {
		if err := paytype.PaytypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "paytype", err: fmt.Errorf("ent: validator failed for field \"paytype\": %w", err)}
		}
	}
	var (
		err  error
		node *Paytype
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaytypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaytypeCreate) SaveX(ctx context.Context) *Paytype {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PaytypeCreate) sqlSave(ctx context.Context) (*Paytype, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PaytypeCreate) createSpec() (*Paytype, *sqlgraph.CreateSpec) {
	var (
		pa    = &Paytype{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: paytype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paytype.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Paytype(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: paytype.FieldPaytype,
		})
		pa.Paytype = value
	}
	if nodes := pc.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paytype.BillsTable,
			Columns: []string{paytype.BillsColumn},
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
	return pa, _spec
}
