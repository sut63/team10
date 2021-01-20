// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/bill"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/paytype"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/unpaybill"
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

// SetPayer sets the Payer field.
func (bu *BillUpdate) SetPayer(s string) *BillUpdate {
	bu.mutation.SetPayer(s)
	return bu
}

// SetPayercontact sets the Payercontact field.
func (bu *BillUpdate) SetPayercontact(s string) *BillUpdate {
	bu.mutation.SetPayercontact(s)
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

// SetEdgesOfPaytypeID sets the EdgesOfPaytype edge to Paytype by id.
func (bu *BillUpdate) SetEdgesOfPaytypeID(id int) *BillUpdate {
	bu.mutation.SetEdgesOfPaytypeID(id)
	return bu
}

// SetNillableEdgesOfPaytypeID sets the EdgesOfPaytype edge to Paytype by id if the given value is not nil.
func (bu *BillUpdate) SetNillableEdgesOfPaytypeID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetEdgesOfPaytypeID(*id)
	}
	return bu
}

// SetEdgesOfPaytype sets the EdgesOfPaytype edge to Paytype.
func (bu *BillUpdate) SetEdgesOfPaytype(p *Paytype) *BillUpdate {
	return bu.SetEdgesOfPaytypeID(p.ID)
}

// SetEdgesOfOfficerID sets the EdgesOfOfficer edge to Financier by id.
func (bu *BillUpdate) SetEdgesOfOfficerID(id int) *BillUpdate {
	bu.mutation.SetEdgesOfOfficerID(id)
	return bu
}

// SetNillableEdgesOfOfficerID sets the EdgesOfOfficer edge to Financier by id if the given value is not nil.
func (bu *BillUpdate) SetNillableEdgesOfOfficerID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetEdgesOfOfficerID(*id)
	}
	return bu
}

// SetEdgesOfOfficer sets the EdgesOfOfficer edge to Financier.
func (bu *BillUpdate) SetEdgesOfOfficer(f *Financier) *BillUpdate {
	return bu.SetEdgesOfOfficerID(f.ID)
}

// SetEdgesOfTreatmentID sets the EdgesOfTreatment edge to Unpaybill by id.
func (bu *BillUpdate) SetEdgesOfTreatmentID(id int) *BillUpdate {
	bu.mutation.SetEdgesOfTreatmentID(id)
	return bu
}

// SetNillableEdgesOfTreatmentID sets the EdgesOfTreatment edge to Unpaybill by id if the given value is not nil.
func (bu *BillUpdate) SetNillableEdgesOfTreatmentID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetEdgesOfTreatmentID(*id)
	}
	return bu
}

// SetEdgesOfTreatment sets the EdgesOfTreatment edge to Unpaybill.
func (bu *BillUpdate) SetEdgesOfTreatment(u *Unpaybill) *BillUpdate {
	return bu.SetEdgesOfTreatmentID(u.ID)
}

// Mutation returns the BillMutation object of the builder.
func (bu *BillUpdate) Mutation() *BillMutation {
	return bu.mutation
}

// ClearEdgesOfPaytype clears the EdgesOfPaytype edge to Paytype.
func (bu *BillUpdate) ClearEdgesOfPaytype() *BillUpdate {
	bu.mutation.ClearEdgesOfPaytype()
	return bu
}

// ClearEdgesOfOfficer clears the EdgesOfOfficer edge to Financier.
func (bu *BillUpdate) ClearEdgesOfOfficer() *BillUpdate {
	bu.mutation.ClearEdgesOfOfficer()
	return bu
}

// ClearEdgesOfTreatment clears the EdgesOfTreatment edge to Unpaybill.
func (bu *BillUpdate) ClearEdgesOfTreatment() *BillUpdate {
	bu.mutation.ClearEdgesOfTreatment()
	return bu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BillUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := bu.mutation.Payer(); ok {
		if err := bill.PayerValidator(v); err != nil {
			return 0, &ValidationError{Name: "Payer", err: fmt.Errorf("ent: validator failed for field \"Payer\": %w", err)}
		}
	}
	if v, ok := bu.mutation.Payercontact(); ok {
		if err := bill.PayercontactValidator(v); err != nil {
			return 0, &ValidationError{Name: "Payercontact", err: fmt.Errorf("ent: validator failed for field \"Payercontact\": %w", err)}
		}
	}
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
	if value, ok := bu.mutation.Payer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldPayer,
		})
	}
	if value, ok := bu.mutation.Payercontact(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldPayercontact,
		})
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
	if bu.mutation.EdgesOfPaytypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.EdgesOfPaytypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.EdgesOfOfficerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.EdgesOfOfficerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.EdgesOfTreatmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.EdgesOfTreatmentTable,
			Columns: []string{bill.EdgesOfTreatmentColumn},
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
	if nodes := bu.mutation.EdgesOfTreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.EdgesOfTreatmentTable,
			Columns: []string{bill.EdgesOfTreatmentColumn},
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

// SetPayer sets the Payer field.
func (buo *BillUpdateOne) SetPayer(s string) *BillUpdateOne {
	buo.mutation.SetPayer(s)
	return buo
}

// SetPayercontact sets the Payercontact field.
func (buo *BillUpdateOne) SetPayercontact(s string) *BillUpdateOne {
	buo.mutation.SetPayercontact(s)
	return buo
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

// SetEdgesOfPaytypeID sets the EdgesOfPaytype edge to Paytype by id.
func (buo *BillUpdateOne) SetEdgesOfPaytypeID(id int) *BillUpdateOne {
	buo.mutation.SetEdgesOfPaytypeID(id)
	return buo
}

// SetNillableEdgesOfPaytypeID sets the EdgesOfPaytype edge to Paytype by id if the given value is not nil.
func (buo *BillUpdateOne) SetNillableEdgesOfPaytypeID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetEdgesOfPaytypeID(*id)
	}
	return buo
}

// SetEdgesOfPaytype sets the EdgesOfPaytype edge to Paytype.
func (buo *BillUpdateOne) SetEdgesOfPaytype(p *Paytype) *BillUpdateOne {
	return buo.SetEdgesOfPaytypeID(p.ID)
}

// SetEdgesOfOfficerID sets the EdgesOfOfficer edge to Financier by id.
func (buo *BillUpdateOne) SetEdgesOfOfficerID(id int) *BillUpdateOne {
	buo.mutation.SetEdgesOfOfficerID(id)
	return buo
}

// SetNillableEdgesOfOfficerID sets the EdgesOfOfficer edge to Financier by id if the given value is not nil.
func (buo *BillUpdateOne) SetNillableEdgesOfOfficerID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetEdgesOfOfficerID(*id)
	}
	return buo
}

// SetEdgesOfOfficer sets the EdgesOfOfficer edge to Financier.
func (buo *BillUpdateOne) SetEdgesOfOfficer(f *Financier) *BillUpdateOne {
	return buo.SetEdgesOfOfficerID(f.ID)
}

// SetEdgesOfTreatmentID sets the EdgesOfTreatment edge to Unpaybill by id.
func (buo *BillUpdateOne) SetEdgesOfTreatmentID(id int) *BillUpdateOne {
	buo.mutation.SetEdgesOfTreatmentID(id)
	return buo
}

// SetNillableEdgesOfTreatmentID sets the EdgesOfTreatment edge to Unpaybill by id if the given value is not nil.
func (buo *BillUpdateOne) SetNillableEdgesOfTreatmentID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetEdgesOfTreatmentID(*id)
	}
	return buo
}

// SetEdgesOfTreatment sets the EdgesOfTreatment edge to Unpaybill.
func (buo *BillUpdateOne) SetEdgesOfTreatment(u *Unpaybill) *BillUpdateOne {
	return buo.SetEdgesOfTreatmentID(u.ID)
}

// Mutation returns the BillMutation object of the builder.
func (buo *BillUpdateOne) Mutation() *BillMutation {
	return buo.mutation
}

// ClearEdgesOfPaytype clears the EdgesOfPaytype edge to Paytype.
func (buo *BillUpdateOne) ClearEdgesOfPaytype() *BillUpdateOne {
	buo.mutation.ClearEdgesOfPaytype()
	return buo
}

// ClearEdgesOfOfficer clears the EdgesOfOfficer edge to Financier.
func (buo *BillUpdateOne) ClearEdgesOfOfficer() *BillUpdateOne {
	buo.mutation.ClearEdgesOfOfficer()
	return buo
}

// ClearEdgesOfTreatment clears the EdgesOfTreatment edge to Unpaybill.
func (buo *BillUpdateOne) ClearEdgesOfTreatment() *BillUpdateOne {
	buo.mutation.ClearEdgesOfTreatment()
	return buo
}

// Save executes the query and returns the updated entity.
func (buo *BillUpdateOne) Save(ctx context.Context) (*Bill, error) {
	if v, ok := buo.mutation.Payer(); ok {
		if err := bill.PayerValidator(v); err != nil {
			return nil, &ValidationError{Name: "Payer", err: fmt.Errorf("ent: validator failed for field \"Payer\": %w", err)}
		}
	}
	if v, ok := buo.mutation.Payercontact(); ok {
		if err := bill.PayercontactValidator(v); err != nil {
			return nil, &ValidationError{Name: "Payercontact", err: fmt.Errorf("ent: validator failed for field \"Payercontact\": %w", err)}
		}
	}
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
	if value, ok := buo.mutation.Payer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldPayer,
		})
	}
	if value, ok := buo.mutation.Payercontact(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldPayercontact,
		})
	}
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
	if buo.mutation.EdgesOfPaytypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.EdgesOfPaytypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.EdgesOfOfficerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.EdgesOfOfficerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.EdgesOfTreatmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.EdgesOfTreatmentTable,
			Columns: []string{bill.EdgesOfTreatmentColumn},
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
	if nodes := buo.mutation.EdgesOfTreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.EdgesOfTreatmentTable,
			Columns: []string{bill.EdgesOfTreatmentColumn},
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
