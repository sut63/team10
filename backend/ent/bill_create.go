// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/bill"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/paytype"
	"github.com/team10/app/ent/unpaybill"
)

// BillCreate is the builder for creating a Bill entity.
type BillCreate struct {
	config
	mutation *BillMutation
	hooks    []Hook
}

// SetAmount sets the Amount field.
func (bc *BillCreate) SetAmount(i int) *BillCreate {
	bc.mutation.SetAmount(i)
	return bc
}

// SetPayer sets the Payer field.
func (bc *BillCreate) SetPayer(s string) *BillCreate {
	bc.mutation.SetPayer(s)
	return bc
}

// SetPayercontact sets the Payercontact field.
func (bc *BillCreate) SetPayercontact(s string) *BillCreate {
	bc.mutation.SetPayercontact(s)
	return bc
}

// SetDate sets the Date field.
func (bc *BillCreate) SetDate(t time.Time) *BillCreate {
	bc.mutation.SetDate(t)
	return bc
}

// SetEdgesOfPaytypeID sets the EdgesOfPaytype edge to Paytype by id.
func (bc *BillCreate) SetEdgesOfPaytypeID(id int) *BillCreate {
	bc.mutation.SetEdgesOfPaytypeID(id)
	return bc
}

// SetNillableEdgesOfPaytypeID sets the EdgesOfPaytype edge to Paytype by id if the given value is not nil.
func (bc *BillCreate) SetNillableEdgesOfPaytypeID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetEdgesOfPaytypeID(*id)
	}
	return bc
}

// SetEdgesOfPaytype sets the EdgesOfPaytype edge to Paytype.
func (bc *BillCreate) SetEdgesOfPaytype(p *Paytype) *BillCreate {
	return bc.SetEdgesOfPaytypeID(p.ID)
}

// SetEdgesOfOfficerID sets the EdgesOfOfficer edge to Financier by id.
func (bc *BillCreate) SetEdgesOfOfficerID(id int) *BillCreate {
	bc.mutation.SetEdgesOfOfficerID(id)
	return bc
}

// SetNillableEdgesOfOfficerID sets the EdgesOfOfficer edge to Financier by id if the given value is not nil.
func (bc *BillCreate) SetNillableEdgesOfOfficerID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetEdgesOfOfficerID(*id)
	}
	return bc
}

// SetEdgesOfOfficer sets the EdgesOfOfficer edge to Financier.
func (bc *BillCreate) SetEdgesOfOfficer(f *Financier) *BillCreate {
	return bc.SetEdgesOfOfficerID(f.ID)
}

// SetEdgesOfUnpaybillID sets the EdgesOfUnpaybill edge to Unpaybill by id.
func (bc *BillCreate) SetEdgesOfUnpaybillID(id int) *BillCreate {
	bc.mutation.SetEdgesOfUnpaybillID(id)
	return bc
}

// SetNillableEdgesOfUnpaybillID sets the EdgesOfUnpaybill edge to Unpaybill by id if the given value is not nil.
func (bc *BillCreate) SetNillableEdgesOfUnpaybillID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetEdgesOfUnpaybillID(*id)
	}
	return bc
}

// SetEdgesOfUnpaybill sets the EdgesOfUnpaybill edge to Unpaybill.
func (bc *BillCreate) SetEdgesOfUnpaybill(u *Unpaybill) *BillCreate {
	return bc.SetEdgesOfUnpaybillID(u.ID)
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
	if _, ok := bc.mutation.Payer(); !ok {
		return nil, &ValidationError{Name: "Payer", err: errors.New("ent: missing required field \"Payer\"")}
	}
	if v, ok := bc.mutation.Payer(); ok {
		if err := bill.PayerValidator(v); err != nil {
			return nil, &ValidationError{Name: "Payer", err: fmt.Errorf("ent: validator failed for field \"Payer\": %w", err)}
		}
	}
	if _, ok := bc.mutation.Payercontact(); !ok {
		return nil, &ValidationError{Name: "Payercontact", err: errors.New("ent: missing required field \"Payercontact\"")}
	}
	if v, ok := bc.mutation.Payercontact(); ok {
		if err := bill.PayercontactValidator(v); err != nil {
			return nil, &ValidationError{Name: "Payercontact", err: fmt.Errorf("ent: validator failed for field \"Payercontact\": %w", err)}
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
			Type:   field.TypeInt,
			Value:  value,
			Column: bill.FieldAmount,
		})
		b.Amount = value
	}
	if value, ok := bc.mutation.Payer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldPayer,
		})
		b.Payer = value
	}
	if value, ok := bc.mutation.Payercontact(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldPayercontact,
		})
		b.Payercontact = value
	}
	if value, ok := bc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bill.FieldDate,
		})
		b.Date = value
	}
	if nodes := bc.mutation.EdgesOfPaytypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.EdgesOfPaytypeTable,
			Columns: []string{bill.EdgesOfPaytypeColumn},
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
	if nodes := bc.mutation.EdgesOfOfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.EdgesOfOfficerTable,
			Columns: []string{bill.EdgesOfOfficerColumn},
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
	if nodes := bc.mutation.EdgesOfUnpaybillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.EdgesOfUnpaybillTable,
			Columns: []string{bill.EdgesOfUnpaybillColumn},
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
