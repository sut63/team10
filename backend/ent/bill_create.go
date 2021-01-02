// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/theuo/app/ent/bill"
	"github.com/theuo/app/ent/financier"
	"github.com/theuo/app/ent/paytype"
	"github.com/theuo/app/ent/unpaybill"
)

// BillCreate is the builder for creating a Bill entity.
type BillCreate struct {
	config
	mutation *BillMutation
	hooks    []Hook
}

// SetAmount sets the Amount field.
func (bc *BillCreate) SetAmount(s string) *BillCreate {
	bc.mutation.SetAmount(s)
	return bc
}

// SetDate sets the Date field.
func (bc *BillCreate) SetDate(t time.Time) *BillCreate {
	bc.mutation.SetDate(t)
	return bc
}

// SetPaytypeID sets the paytype edge to Paytype by id.
func (bc *BillCreate) SetPaytypeID(id int) *BillCreate {
	bc.mutation.SetPaytypeID(id)
	return bc
}

// SetNillablePaytypeID sets the paytype edge to Paytype by id if the given value is not nil.
func (bc *BillCreate) SetNillablePaytypeID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetPaytypeID(*id)
	}
	return bc
}

// SetPaytype sets the paytype edge to Paytype.
func (bc *BillCreate) SetPaytype(p *Paytype) *BillCreate {
	return bc.SetPaytypeID(p.ID)
}

// SetOfficerID sets the officer edge to Financier by id.
func (bc *BillCreate) SetOfficerID(id int) *BillCreate {
	bc.mutation.SetOfficerID(id)
	return bc
}

// SetNillableOfficerID sets the officer edge to Financier by id if the given value is not nil.
func (bc *BillCreate) SetNillableOfficerID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetOfficerID(*id)
	}
	return bc
}

// SetOfficer sets the officer edge to Financier.
func (bc *BillCreate) SetOfficer(f *Financier) *BillCreate {
	return bc.SetOfficerID(f.ID)
}

// SetTreatmentID sets the treatment edge to Unpaybill by id.
func (bc *BillCreate) SetTreatmentID(id int) *BillCreate {
	bc.mutation.SetTreatmentID(id)
	return bc
}

// SetNillableTreatmentID sets the treatment edge to Unpaybill by id if the given value is not nil.
func (bc *BillCreate) SetNillableTreatmentID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetTreatmentID(*id)
	}
	return bc
}

// SetTreatment sets the treatment edge to Unpaybill.
func (bc *BillCreate) SetTreatment(u *Unpaybill) *BillCreate {
	return bc.SetTreatmentID(u.ID)
}

// Mutation returns the BillMutation object of the builder.
func (bc *BillCreate) Mutation() *BillMutation {
	return bc.mutation
}

// Save creates the Bill in the database.
func (bc *BillCreate) Save(ctx context.Context) (*Bill, error) {
	if _, ok := bc.mutation.Amount(); !ok {
		return nil, &ValidationError{Name: "Amount", err: errors.New("ent: missing required field \"Amount\"")}
	}
	if v, ok := bc.mutation.Amount(); ok {
		if err := bill.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "Amount", err: fmt.Errorf("ent: validator failed for field \"Amount\": %w", err)}
		}
	}
	if _, ok := bc.mutation.Date(); !ok {
		return nil, &ValidationError{Name: "Date", err: errors.New("ent: missing required field \"Date\"")}
	}
	var (
		err  error
		node *Bill
	)
	if len(bc.hooks) == 0 {
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BillCreate) SaveX(ctx context.Context) *Bill {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BillCreate) sqlSave(ctx context.Context) (*Bill, error) {
	b, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	b.ID = int(id)
	return b, nil
}

func (bc *BillCreate) createSpec() (*Bill, *sqlgraph.CreateSpec) {
	var (
		b     = &Bill{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bill.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bill.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldAmount,
		})
		b.Amount = value
	}
	if value, ok := bc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bill.FieldDate,
		})
		b.Date = value
	}
	if nodes := bc.mutation.PaytypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.PaytypeTable,
			Columns: []string{bill.PaytypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: paytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.OfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.OfficerTable,
			Columns: []string{bill.OfficerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: financier.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.TreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.TreatmentTable,
			Columns: []string{bill.TreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unpaybill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return b, _spec
}
