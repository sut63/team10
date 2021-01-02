// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/theuo/app/ent/bill"
	"github.com/theuo/app/ent/financier"
	"github.com/theuo/app/ent/paytype"
	"github.com/theuo/app/ent/predicate"
	"github.com/theuo/app/ent/unpaybill"
)

// BillUpdate is the builder for updating Bill entities.
type BillUpdate struct {
	config
	hooks      []Hook
	mutation   *BillMutation
	predicates []predicate.Bill
}

// Where adds a new predicate for the builder.
func (bu *BillUpdate) Where(ps ...predicate.Bill) *BillUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetAmount sets the Amount field.
func (bu *BillUpdate) SetAmount(s string) *BillUpdate {
	bu.mutation.SetAmount(s)
	return bu
}

// SetDate sets the Date field.
func (bu *BillUpdate) SetDate(t time.Time) *BillUpdate {
	bu.mutation.SetDate(t)
	return bu
}

// SetPaytypeID sets the paytype edge to Paytype by id.
func (bu *BillUpdate) SetPaytypeID(id int) *BillUpdate {
	bu.mutation.SetPaytypeID(id)
	return bu
}

// SetNillablePaytypeID sets the paytype edge to Paytype by id if the given value is not nil.
func (bu *BillUpdate) SetNillablePaytypeID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetPaytypeID(*id)
	}
	return bu
}

// SetPaytype sets the paytype edge to Paytype.
func (bu *BillUpdate) SetPaytype(p *Paytype) *BillUpdate {
	return bu.SetPaytypeID(p.ID)
}

// SetOfficerID sets the officer edge to Financier by id.
func (bu *BillUpdate) SetOfficerID(id int) *BillUpdate {
	bu.mutation.SetOfficerID(id)
	return bu
}

// SetNillableOfficerID sets the officer edge to Financier by id if the given value is not nil.
func (bu *BillUpdate) SetNillableOfficerID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetOfficerID(*id)
	}
	return bu
}

// SetOfficer sets the officer edge to Financier.
func (bu *BillUpdate) SetOfficer(f *Financier) *BillUpdate {
	return bu.SetOfficerID(f.ID)
}

// SetTreatmentID sets the treatment edge to Unpaybill by id.
func (bu *BillUpdate) SetTreatmentID(id int) *BillUpdate {
	bu.mutation.SetTreatmentID(id)
	return bu
}

// SetNillableTreatmentID sets the treatment edge to Unpaybill by id if the given value is not nil.
func (bu *BillUpdate) SetNillableTreatmentID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetTreatmentID(*id)
	}
	return bu
}

// SetTreatment sets the treatment edge to Unpaybill.
func (bu *BillUpdate) SetTreatment(u *Unpaybill) *BillUpdate {
	return bu.SetTreatmentID(u.ID)
}

// Mutation returns the BillMutation object of the builder.
func (bu *BillUpdate) Mutation() *BillMutation {
	return bu.mutation
}

// ClearPaytype clears the paytype edge to Paytype.
func (bu *BillUpdate) ClearPaytype() *BillUpdate {
	bu.mutation.ClearPaytype()
	return bu
}

// ClearOfficer clears the officer edge to Financier.
func (bu *BillUpdate) ClearOfficer() *BillUpdate {
	bu.mutation.ClearOfficer()
	return bu
}

// ClearTreatment clears the treatment edge to Unpaybill.
func (bu *BillUpdate) ClearTreatment() *BillUpdate {
	bu.mutation.ClearTreatment()
	return bu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BillUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := bu.mutation.Amount(); ok {
		if err := bill.AmountValidator(v); err != nil {
			return 0, &ValidationError{Name: "Amount", err: fmt.Errorf("ent: validator failed for field \"Amount\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BillUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BillUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BillUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BillUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bill.Table,
			Columns: bill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bill.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldAmount,
		})
	}
	if value, ok := bu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bill.FieldDate,
		})
	}
	if bu.mutation.PaytypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.PaytypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.OfficerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.OfficerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.TreatmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TreatmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bill.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BillUpdateOne is the builder for updating a single Bill entity.
type BillUpdateOne struct {
	config
	hooks    []Hook
	mutation *BillMutation
}

// SetAmount sets the Amount field.
func (buo *BillUpdateOne) SetAmount(s string) *BillUpdateOne {
	buo.mutation.SetAmount(s)
	return buo
}

// SetDate sets the Date field.
func (buo *BillUpdateOne) SetDate(t time.Time) *BillUpdateOne {
	buo.mutation.SetDate(t)
	return buo
}

// SetPaytypeID sets the paytype edge to Paytype by id.
func (buo *BillUpdateOne) SetPaytypeID(id int) *BillUpdateOne {
	buo.mutation.SetPaytypeID(id)
	return buo
}

// SetNillablePaytypeID sets the paytype edge to Paytype by id if the given value is not nil.
func (buo *BillUpdateOne) SetNillablePaytypeID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetPaytypeID(*id)
	}
	return buo
}

// SetPaytype sets the paytype edge to Paytype.
func (buo *BillUpdateOne) SetPaytype(p *Paytype) *BillUpdateOne {
	return buo.SetPaytypeID(p.ID)
}

// SetOfficerID sets the officer edge to Financier by id.
func (buo *BillUpdateOne) SetOfficerID(id int) *BillUpdateOne {
	buo.mutation.SetOfficerID(id)
	return buo
}

// SetNillableOfficerID sets the officer edge to Financier by id if the given value is not nil.
func (buo *BillUpdateOne) SetNillableOfficerID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetOfficerID(*id)
	}
	return buo
}

// SetOfficer sets the officer edge to Financier.
func (buo *BillUpdateOne) SetOfficer(f *Financier) *BillUpdateOne {
	return buo.SetOfficerID(f.ID)
}

// SetTreatmentID sets the treatment edge to Unpaybill by id.
func (buo *BillUpdateOne) SetTreatmentID(id int) *BillUpdateOne {
	buo.mutation.SetTreatmentID(id)
	return buo
}

// SetNillableTreatmentID sets the treatment edge to Unpaybill by id if the given value is not nil.
func (buo *BillUpdateOne) SetNillableTreatmentID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetTreatmentID(*id)
	}
	return buo
}

// SetTreatment sets the treatment edge to Unpaybill.
func (buo *BillUpdateOne) SetTreatment(u *Unpaybill) *BillUpdateOne {
	return buo.SetTreatmentID(u.ID)
}

// Mutation returns the BillMutation object of the builder.
func (buo *BillUpdateOne) Mutation() *BillMutation {
	return buo.mutation
}

// ClearPaytype clears the paytype edge to Paytype.
func (buo *BillUpdateOne) ClearPaytype() *BillUpdateOne {
	buo.mutation.ClearPaytype()
	return buo
}

// ClearOfficer clears the officer edge to Financier.
func (buo *BillUpdateOne) ClearOfficer() *BillUpdateOne {
	buo.mutation.ClearOfficer()
	return buo
}

// ClearTreatment clears the treatment edge to Unpaybill.
func (buo *BillUpdateOne) ClearTreatment() *BillUpdateOne {
	buo.mutation.ClearTreatment()
	return buo
}

// Save executes the query and returns the updated entity.
func (buo *BillUpdateOne) Save(ctx context.Context) (*Bill, error) {
	if v, ok := buo.mutation.Amount(); ok {
		if err := bill.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "Amount", err: fmt.Errorf("ent: validator failed for field \"Amount\": %w", err)}
		}
	}

	var (
		err  error
		node *Bill
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BillUpdateOne) SaveX(ctx context.Context) *Bill {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BillUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BillUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BillUpdateOne) sqlSave(ctx context.Context) (b *Bill, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bill.Table,
			Columns: bill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bill.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Bill.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldAmount,
		})
	}
	if value, ok := buo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bill.FieldDate,
		})
	}
	if buo.mutation.PaytypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.PaytypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.OfficerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.OfficerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.TreatmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TreatmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	b = &Bill{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bill.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}