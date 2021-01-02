// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/b6109868/app/ent/doctorinfo"
	"github.com/b6109868/app/ent/patientrecord"
	"github.com/b6109868/app/ent/predicate"
	"github.com/b6109868/app/ent/treatment"
	"github.com/b6109868/app/ent/typetreatment"
	"github.com/b6109868/app/ent/unpaybill"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// TreatmentUpdate is the builder for updating Treatment entities.
type TreatmentUpdate struct {
	config
	hooks      []Hook
	mutation   *TreatmentMutation
	predicates []predicate.Treatment
}

// Where adds a new predicate for the builder.
func (tu *TreatmentUpdate) Where(ps ...predicate.Treatment) *TreatmentUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetTreatment sets the treatment field.
func (tu *TreatmentUpdate) SetTreatment(s string) *TreatmentUpdate {
	tu.mutation.SetTreatment(s)
	return tu
}

// SetDatetime sets the datetime field.
func (tu *TreatmentUpdate) SetDatetime(t time.Time) *TreatmentUpdate {
	tu.mutation.SetDatetime(t)
	return tu
}

// SetTypetreatmentID sets the typetreatment edge to Typetreatment by id.
func (tu *TreatmentUpdate) SetTypetreatmentID(id int) *TreatmentUpdate {
	tu.mutation.SetTypetreatmentID(id)
	return tu
}

// SetNillableTypetreatmentID sets the typetreatment edge to Typetreatment by id if the given value is not nil.
func (tu *TreatmentUpdate) SetNillableTypetreatmentID(id *int) *TreatmentUpdate {
	if id != nil {
		tu = tu.SetTypetreatmentID(*id)
	}
	return tu
}

// SetTypetreatment sets the typetreatment edge to Typetreatment.
func (tu *TreatmentUpdate) SetTypetreatment(t *Typetreatment) *TreatmentUpdate {
	return tu.SetTypetreatmentID(t.ID)
}

// SetPatientrecordID sets the patientrecord edge to Patientrecord by id.
func (tu *TreatmentUpdate) SetPatientrecordID(id int) *TreatmentUpdate {
	tu.mutation.SetPatientrecordID(id)
	return tu
}

// SetNillablePatientrecordID sets the patientrecord edge to Patientrecord by id if the given value is not nil.
func (tu *TreatmentUpdate) SetNillablePatientrecordID(id *int) *TreatmentUpdate {
	if id != nil {
		tu = tu.SetPatientrecordID(*id)
	}
	return tu
}

// SetPatientrecord sets the patientrecord edge to Patientrecord.
func (tu *TreatmentUpdate) SetPatientrecord(p *Patientrecord) *TreatmentUpdate {
	return tu.SetPatientrecordID(p.ID)
}

// SetDoctorinfoID sets the doctorinfo edge to Doctorinfo by id.
func (tu *TreatmentUpdate) SetDoctorinfoID(id int) *TreatmentUpdate {
	tu.mutation.SetDoctorinfoID(id)
	return tu
}

// SetNillableDoctorinfoID sets the doctorinfo edge to Doctorinfo by id if the given value is not nil.
func (tu *TreatmentUpdate) SetNillableDoctorinfoID(id *int) *TreatmentUpdate {
	if id != nil {
		tu = tu.SetDoctorinfoID(*id)
	}
	return tu
}

// SetDoctorinfo sets the doctorinfo edge to Doctorinfo.
func (tu *TreatmentUpdate) SetDoctorinfo(d *Doctorinfo) *TreatmentUpdate {
	return tu.SetDoctorinfoID(d.ID)
}

// SetUnpaybillsID sets the unpaybills edge to Unpaybill by id.
func (tu *TreatmentUpdate) SetUnpaybillsID(id int) *TreatmentUpdate {
	tu.mutation.SetUnpaybillsID(id)
	return tu
}

// SetNillableUnpaybillsID sets the unpaybills edge to Unpaybill by id if the given value is not nil.
func (tu *TreatmentUpdate) SetNillableUnpaybillsID(id *int) *TreatmentUpdate {
	if id != nil {
		tu = tu.SetUnpaybillsID(*id)
	}
	return tu
}

// SetUnpaybills sets the unpaybills edge to Unpaybill.
func (tu *TreatmentUpdate) SetUnpaybills(u *Unpaybill) *TreatmentUpdate {
	return tu.SetUnpaybillsID(u.ID)
}

// Mutation returns the TreatmentMutation object of the builder.
func (tu *TreatmentUpdate) Mutation() *TreatmentMutation {
	return tu.mutation
}

// ClearTypetreatment clears the typetreatment edge to Typetreatment.
func (tu *TreatmentUpdate) ClearTypetreatment() *TreatmentUpdate {
	tu.mutation.ClearTypetreatment()
	return tu
}

// ClearPatientrecord clears the patientrecord edge to Patientrecord.
func (tu *TreatmentUpdate) ClearPatientrecord() *TreatmentUpdate {
	tu.mutation.ClearPatientrecord()
	return tu
}

// ClearDoctorinfo clears the doctorinfo edge to Doctorinfo.
func (tu *TreatmentUpdate) ClearDoctorinfo() *TreatmentUpdate {
	tu.mutation.ClearDoctorinfo()
	return tu
}

// ClearUnpaybills clears the unpaybills edge to Unpaybill.
func (tu *TreatmentUpdate) ClearUnpaybills() *TreatmentUpdate {
	tu.mutation.ClearUnpaybills()
	return tu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TreatmentUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := tu.mutation.Treatment(); ok {
		if err := treatment.TreatmentValidator(v); err != nil {
			return 0, &ValidationError{Name: "treatment", err: fmt.Errorf("ent: validator failed for field \"treatment\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TreatmentUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TreatmentUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TreatmentUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TreatmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   treatment.Table,
			Columns: treatment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: treatment.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Treatment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatment.FieldTreatment,
		})
	}
	if value, ok := tu.mutation.Datetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: treatment.FieldDatetime,
		})
	}
	if tu.mutation.TypetreatmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.TypetreatmentTable,
			Columns: []string{treatment.TypetreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typetreatment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TypetreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.TypetreatmentTable,
			Columns: []string{treatment.TypetreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typetreatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.PatientrecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.PatientrecordTable,
			Columns: []string{treatment.PatientrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrecord.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.PatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.PatientrecordTable,
			Columns: []string{treatment.PatientrecordColumn},
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
	if tu.mutation.DoctorinfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.DoctorinfoTable,
			Columns: []string{treatment.DoctorinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctorinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.DoctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.DoctorinfoTable,
			Columns: []string{treatment.DoctorinfoColumn},
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
	if tu.mutation.UnpaybillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   treatment.UnpaybillsTable,
			Columns: []string{treatment.UnpaybillsColumn},
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
	if nodes := tu.mutation.UnpaybillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   treatment.UnpaybillsTable,
			Columns: []string{treatment.UnpaybillsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{treatment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TreatmentUpdateOne is the builder for updating a single Treatment entity.
type TreatmentUpdateOne struct {
	config
	hooks    []Hook
	mutation *TreatmentMutation
}

// SetTreatment sets the treatment field.
func (tuo *TreatmentUpdateOne) SetTreatment(s string) *TreatmentUpdateOne {
	tuo.mutation.SetTreatment(s)
	return tuo
}

// SetDatetime sets the datetime field.
func (tuo *TreatmentUpdateOne) SetDatetime(t time.Time) *TreatmentUpdateOne {
	tuo.mutation.SetDatetime(t)
	return tuo
}

// SetTypetreatmentID sets the typetreatment edge to Typetreatment by id.
func (tuo *TreatmentUpdateOne) SetTypetreatmentID(id int) *TreatmentUpdateOne {
	tuo.mutation.SetTypetreatmentID(id)
	return tuo
}

// SetNillableTypetreatmentID sets the typetreatment edge to Typetreatment by id if the given value is not nil.
func (tuo *TreatmentUpdateOne) SetNillableTypetreatmentID(id *int) *TreatmentUpdateOne {
	if id != nil {
		tuo = tuo.SetTypetreatmentID(*id)
	}
	return tuo
}

// SetTypetreatment sets the typetreatment edge to Typetreatment.
func (tuo *TreatmentUpdateOne) SetTypetreatment(t *Typetreatment) *TreatmentUpdateOne {
	return tuo.SetTypetreatmentID(t.ID)
}

// SetPatientrecordID sets the patientrecord edge to Patientrecord by id.
func (tuo *TreatmentUpdateOne) SetPatientrecordID(id int) *TreatmentUpdateOne {
	tuo.mutation.SetPatientrecordID(id)
	return tuo
}

// SetNillablePatientrecordID sets the patientrecord edge to Patientrecord by id if the given value is not nil.
func (tuo *TreatmentUpdateOne) SetNillablePatientrecordID(id *int) *TreatmentUpdateOne {
	if id != nil {
		tuo = tuo.SetPatientrecordID(*id)
	}
	return tuo
}

// SetPatientrecord sets the patientrecord edge to Patientrecord.
func (tuo *TreatmentUpdateOne) SetPatientrecord(p *Patientrecord) *TreatmentUpdateOne {
	return tuo.SetPatientrecordID(p.ID)
}

// SetDoctorinfoID sets the doctorinfo edge to Doctorinfo by id.
func (tuo *TreatmentUpdateOne) SetDoctorinfoID(id int) *TreatmentUpdateOne {
	tuo.mutation.SetDoctorinfoID(id)
	return tuo
}

// SetNillableDoctorinfoID sets the doctorinfo edge to Doctorinfo by id if the given value is not nil.
func (tuo *TreatmentUpdateOne) SetNillableDoctorinfoID(id *int) *TreatmentUpdateOne {
	if id != nil {
		tuo = tuo.SetDoctorinfoID(*id)
	}
	return tuo
}

// SetDoctorinfo sets the doctorinfo edge to Doctorinfo.
func (tuo *TreatmentUpdateOne) SetDoctorinfo(d *Doctorinfo) *TreatmentUpdateOne {
	return tuo.SetDoctorinfoID(d.ID)
}

// SetUnpaybillsID sets the unpaybills edge to Unpaybill by id.
func (tuo *TreatmentUpdateOne) SetUnpaybillsID(id int) *TreatmentUpdateOne {
	tuo.mutation.SetUnpaybillsID(id)
	return tuo
}

// SetNillableUnpaybillsID sets the unpaybills edge to Unpaybill by id if the given value is not nil.
func (tuo *TreatmentUpdateOne) SetNillableUnpaybillsID(id *int) *TreatmentUpdateOne {
	if id != nil {
		tuo = tuo.SetUnpaybillsID(*id)
	}
	return tuo
}

// SetUnpaybills sets the unpaybills edge to Unpaybill.
func (tuo *TreatmentUpdateOne) SetUnpaybills(u *Unpaybill) *TreatmentUpdateOne {
	return tuo.SetUnpaybillsID(u.ID)
}

// Mutation returns the TreatmentMutation object of the builder.
func (tuo *TreatmentUpdateOne) Mutation() *TreatmentMutation {
	return tuo.mutation
}

// ClearTypetreatment clears the typetreatment edge to Typetreatment.
func (tuo *TreatmentUpdateOne) ClearTypetreatment() *TreatmentUpdateOne {
	tuo.mutation.ClearTypetreatment()
	return tuo
}

// ClearPatientrecord clears the patientrecord edge to Patientrecord.
func (tuo *TreatmentUpdateOne) ClearPatientrecord() *TreatmentUpdateOne {
	tuo.mutation.ClearPatientrecord()
	return tuo
}

// ClearDoctorinfo clears the doctorinfo edge to Doctorinfo.
func (tuo *TreatmentUpdateOne) ClearDoctorinfo() *TreatmentUpdateOne {
	tuo.mutation.ClearDoctorinfo()
	return tuo
}

// ClearUnpaybills clears the unpaybills edge to Unpaybill.
func (tuo *TreatmentUpdateOne) ClearUnpaybills() *TreatmentUpdateOne {
	tuo.mutation.ClearUnpaybills()
	return tuo
}

// Save executes the query and returns the updated entity.
func (tuo *TreatmentUpdateOne) Save(ctx context.Context) (*Treatment, error) {
	if v, ok := tuo.mutation.Treatment(); ok {
		if err := treatment.TreatmentValidator(v); err != nil {
			return nil, &ValidationError{Name: "treatment", err: fmt.Errorf("ent: validator failed for field \"treatment\": %w", err)}
		}
	}

	var (
		err  error
		node *Treatment
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TreatmentUpdateOne) SaveX(ctx context.Context) *Treatment {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TreatmentUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TreatmentUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TreatmentUpdateOne) sqlSave(ctx context.Context) (t *Treatment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   treatment.Table,
			Columns: treatment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: treatment.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Treatment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := tuo.mutation.Treatment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatment.FieldTreatment,
		})
	}
	if value, ok := tuo.mutation.Datetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: treatment.FieldDatetime,
		})
	}
	if tuo.mutation.TypetreatmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.TypetreatmentTable,
			Columns: []string{treatment.TypetreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typetreatment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TypetreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.TypetreatmentTable,
			Columns: []string{treatment.TypetreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typetreatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.PatientrecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.PatientrecordTable,
			Columns: []string{treatment.PatientrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrecord.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.PatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.PatientrecordTable,
			Columns: []string{treatment.PatientrecordColumn},
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
	if tuo.mutation.DoctorinfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.DoctorinfoTable,
			Columns: []string{treatment.DoctorinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctorinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.DoctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.DoctorinfoTable,
			Columns: []string{treatment.DoctorinfoColumn},
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
	if tuo.mutation.UnpaybillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   treatment.UnpaybillsTable,
			Columns: []string{treatment.UnpaybillsColumn},
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
	if nodes := tuo.mutation.UnpaybillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   treatment.UnpaybillsTable,
			Columns: []string{treatment.UnpaybillsColumn},
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
	t = &Treatment{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{treatment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
