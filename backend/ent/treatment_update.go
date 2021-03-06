// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/doctor"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/treatment"
	"github.com/team10/app/ent/typetreatment"
	"github.com/team10/app/ent/unpaybill"
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

// SetSymptom sets the Symptom field.
func (tu *TreatmentUpdate) SetSymptom(s string) *TreatmentUpdate {
	tu.mutation.SetSymptom(s)
	return tu
}

// SetTreat sets the Treat field.
func (tu *TreatmentUpdate) SetTreat(s string) *TreatmentUpdate {
	tu.mutation.SetTreat(s)
	return tu
}

// SetMedicine sets the Medicine field.
func (tu *TreatmentUpdate) SetMedicine(s string) *TreatmentUpdate {
	tu.mutation.SetMedicine(s)
	return tu
}

// SetDatetreat sets the Datetreat field.
func (tu *TreatmentUpdate) SetDatetreat(t time.Time) *TreatmentUpdate {
	tu.mutation.SetDatetreat(t)
	return tu
}

// SetEdgesOfTypetreatmentID sets the EdgesOfTypetreatment edge to Typetreatment by id.
func (tu *TreatmentUpdate) SetEdgesOfTypetreatmentID(id int) *TreatmentUpdate {
	tu.mutation.SetEdgesOfTypetreatmentID(id)
	return tu
}

// SetNillableEdgesOfTypetreatmentID sets the EdgesOfTypetreatment edge to Typetreatment by id if the given value is not nil.
func (tu *TreatmentUpdate) SetNillableEdgesOfTypetreatmentID(id *int) *TreatmentUpdate {
	if id != nil {
		tu = tu.SetEdgesOfTypetreatmentID(*id)
	}
	return tu
}

// SetEdgesOfTypetreatment sets the EdgesOfTypetreatment edge to Typetreatment.
func (tu *TreatmentUpdate) SetEdgesOfTypetreatment(t *Typetreatment) *TreatmentUpdate {
	return tu.SetEdgesOfTypetreatmentID(t.ID)
}

// SetEdgesOfPatientrecordID sets the EdgesOfPatientrecord edge to Patientrecord by id.
func (tu *TreatmentUpdate) SetEdgesOfPatientrecordID(id int) *TreatmentUpdate {
	tu.mutation.SetEdgesOfPatientrecordID(id)
	return tu
}

// SetNillableEdgesOfPatientrecordID sets the EdgesOfPatientrecord edge to Patientrecord by id if the given value is not nil.
func (tu *TreatmentUpdate) SetNillableEdgesOfPatientrecordID(id *int) *TreatmentUpdate {
	if id != nil {
		tu = tu.SetEdgesOfPatientrecordID(*id)
	}
	return tu
}

// SetEdgesOfPatientrecord sets the EdgesOfPatientrecord edge to Patientrecord.
func (tu *TreatmentUpdate) SetEdgesOfPatientrecord(p *Patientrecord) *TreatmentUpdate {
	return tu.SetEdgesOfPatientrecordID(p.ID)
}

// SetEdgesOfDoctorID sets the EdgesOfDoctor edge to Doctor by id.
func (tu *TreatmentUpdate) SetEdgesOfDoctorID(id int) *TreatmentUpdate {
	tu.mutation.SetEdgesOfDoctorID(id)
	return tu
}

// SetNillableEdgesOfDoctorID sets the EdgesOfDoctor edge to Doctor by id if the given value is not nil.
func (tu *TreatmentUpdate) SetNillableEdgesOfDoctorID(id *int) *TreatmentUpdate {
	if id != nil {
		tu = tu.SetEdgesOfDoctorID(*id)
	}
	return tu
}

// SetEdgesOfDoctor sets the EdgesOfDoctor edge to Doctor.
func (tu *TreatmentUpdate) SetEdgesOfDoctor(d *Doctor) *TreatmentUpdate {
	return tu.SetEdgesOfDoctorID(d.ID)
}

// SetEdgesOfUnpaybillsID sets the EdgesOfUnpaybills edge to Unpaybill by id.
func (tu *TreatmentUpdate) SetEdgesOfUnpaybillsID(id int) *TreatmentUpdate {
	tu.mutation.SetEdgesOfUnpaybillsID(id)
	return tu
}

// SetNillableEdgesOfUnpaybillsID sets the EdgesOfUnpaybills edge to Unpaybill by id if the given value is not nil.
func (tu *TreatmentUpdate) SetNillableEdgesOfUnpaybillsID(id *int) *TreatmentUpdate {
	if id != nil {
		tu = tu.SetEdgesOfUnpaybillsID(*id)
	}
	return tu
}

// SetEdgesOfUnpaybills sets the EdgesOfUnpaybills edge to Unpaybill.
func (tu *TreatmentUpdate) SetEdgesOfUnpaybills(u *Unpaybill) *TreatmentUpdate {
	return tu.SetEdgesOfUnpaybillsID(u.ID)
}

// Mutation returns the TreatmentMutation object of the builder.
func (tu *TreatmentUpdate) Mutation() *TreatmentMutation {
	return tu.mutation
}

// ClearEdgesOfTypetreatment clears the EdgesOfTypetreatment edge to Typetreatment.
func (tu *TreatmentUpdate) ClearEdgesOfTypetreatment() *TreatmentUpdate {
	tu.mutation.ClearEdgesOfTypetreatment()
	return tu
}

// ClearEdgesOfPatientrecord clears the EdgesOfPatientrecord edge to Patientrecord.
func (tu *TreatmentUpdate) ClearEdgesOfPatientrecord() *TreatmentUpdate {
	tu.mutation.ClearEdgesOfPatientrecord()
	return tu
}

// ClearEdgesOfDoctor clears the EdgesOfDoctor edge to Doctor.
func (tu *TreatmentUpdate) ClearEdgesOfDoctor() *TreatmentUpdate {
	tu.mutation.ClearEdgesOfDoctor()
	return tu
}

// ClearEdgesOfUnpaybills clears the EdgesOfUnpaybills edge to Unpaybill.
func (tu *TreatmentUpdate) ClearEdgesOfUnpaybills() *TreatmentUpdate {
	tu.mutation.ClearEdgesOfUnpaybills()
	return tu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TreatmentUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := tu.mutation.Symptom(); ok {
		if err := treatment.SymptomValidator(v); err != nil {
			return 0, &ValidationError{Name: "Symptom", err: fmt.Errorf("ent: validator failed for field \"Symptom\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Treat(); ok {
		if err := treatment.TreatValidator(v); err != nil {
			return 0, &ValidationError{Name: "Treat", err: fmt.Errorf("ent: validator failed for field \"Treat\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Medicine(); ok {
		if err := treatment.MedicineValidator(v); err != nil {
			return 0, &ValidationError{Name: "Medicine", err: fmt.Errorf("ent: validator failed for field \"Medicine\": %w", err)}
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
	if value, ok := tu.mutation.Symptom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatment.FieldSymptom,
		})
	}
	if value, ok := tu.mutation.Treat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatment.FieldTreat,
		})
	}
	if value, ok := tu.mutation.Medicine(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatment.FieldMedicine,
		})
	}
	if value, ok := tu.mutation.Datetreat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: treatment.FieldDatetreat,
		})
	}
	if tu.mutation.EdgesOfTypetreatmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfTypetreatmentTable,
			Columns: []string{treatment.EdgesOfTypetreatmentColumn},
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
	if nodes := tu.mutation.EdgesOfTypetreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfTypetreatmentTable,
			Columns: []string{treatment.EdgesOfTypetreatmentColumn},
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
	if tu.mutation.EdgesOfPatientrecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfPatientrecordTable,
			Columns: []string{treatment.EdgesOfPatientrecordColumn},
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
	if nodes := tu.mutation.EdgesOfPatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfPatientrecordTable,
			Columns: []string{treatment.EdgesOfPatientrecordColumn},
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
	if tu.mutation.EdgesOfDoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfDoctorTable,
			Columns: []string{treatment.EdgesOfDoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.EdgesOfDoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfDoctorTable,
			Columns: []string{treatment.EdgesOfDoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.EdgesOfUnpaybillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   treatment.EdgesOfUnpaybillsTable,
			Columns: []string{treatment.EdgesOfUnpaybillsColumn},
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
	if nodes := tu.mutation.EdgesOfUnpaybillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   treatment.EdgesOfUnpaybillsTable,
			Columns: []string{treatment.EdgesOfUnpaybillsColumn},
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

// SetSymptom sets the Symptom field.
func (tuo *TreatmentUpdateOne) SetSymptom(s string) *TreatmentUpdateOne {
	tuo.mutation.SetSymptom(s)
	return tuo
}

// SetTreat sets the Treat field.
func (tuo *TreatmentUpdateOne) SetTreat(s string) *TreatmentUpdateOne {
	tuo.mutation.SetTreat(s)
	return tuo
}

// SetMedicine sets the Medicine field.
func (tuo *TreatmentUpdateOne) SetMedicine(s string) *TreatmentUpdateOne {
	tuo.mutation.SetMedicine(s)
	return tuo
}

// SetDatetreat sets the Datetreat field.
func (tuo *TreatmentUpdateOne) SetDatetreat(t time.Time) *TreatmentUpdateOne {
	tuo.mutation.SetDatetreat(t)
	return tuo
}

// SetEdgesOfTypetreatmentID sets the EdgesOfTypetreatment edge to Typetreatment by id.
func (tuo *TreatmentUpdateOne) SetEdgesOfTypetreatmentID(id int) *TreatmentUpdateOne {
	tuo.mutation.SetEdgesOfTypetreatmentID(id)
	return tuo
}

// SetNillableEdgesOfTypetreatmentID sets the EdgesOfTypetreatment edge to Typetreatment by id if the given value is not nil.
func (tuo *TreatmentUpdateOne) SetNillableEdgesOfTypetreatmentID(id *int) *TreatmentUpdateOne {
	if id != nil {
		tuo = tuo.SetEdgesOfTypetreatmentID(*id)
	}
	return tuo
}

// SetEdgesOfTypetreatment sets the EdgesOfTypetreatment edge to Typetreatment.
func (tuo *TreatmentUpdateOne) SetEdgesOfTypetreatment(t *Typetreatment) *TreatmentUpdateOne {
	return tuo.SetEdgesOfTypetreatmentID(t.ID)
}

// SetEdgesOfPatientrecordID sets the EdgesOfPatientrecord edge to Patientrecord by id.
func (tuo *TreatmentUpdateOne) SetEdgesOfPatientrecordID(id int) *TreatmentUpdateOne {
	tuo.mutation.SetEdgesOfPatientrecordID(id)
	return tuo
}

// SetNillableEdgesOfPatientrecordID sets the EdgesOfPatientrecord edge to Patientrecord by id if the given value is not nil.
func (tuo *TreatmentUpdateOne) SetNillableEdgesOfPatientrecordID(id *int) *TreatmentUpdateOne {
	if id != nil {
		tuo = tuo.SetEdgesOfPatientrecordID(*id)
	}
	return tuo
}

// SetEdgesOfPatientrecord sets the EdgesOfPatientrecord edge to Patientrecord.
func (tuo *TreatmentUpdateOne) SetEdgesOfPatientrecord(p *Patientrecord) *TreatmentUpdateOne {
	return tuo.SetEdgesOfPatientrecordID(p.ID)
}

// SetEdgesOfDoctorID sets the EdgesOfDoctor edge to Doctor by id.
func (tuo *TreatmentUpdateOne) SetEdgesOfDoctorID(id int) *TreatmentUpdateOne {
	tuo.mutation.SetEdgesOfDoctorID(id)
	return tuo
}

// SetNillableEdgesOfDoctorID sets the EdgesOfDoctor edge to Doctor by id if the given value is not nil.
func (tuo *TreatmentUpdateOne) SetNillableEdgesOfDoctorID(id *int) *TreatmentUpdateOne {
	if id != nil {
		tuo = tuo.SetEdgesOfDoctorID(*id)
	}
	return tuo
}

// SetEdgesOfDoctor sets the EdgesOfDoctor edge to Doctor.
func (tuo *TreatmentUpdateOne) SetEdgesOfDoctor(d *Doctor) *TreatmentUpdateOne {
	return tuo.SetEdgesOfDoctorID(d.ID)
}

// SetEdgesOfUnpaybillsID sets the EdgesOfUnpaybills edge to Unpaybill by id.
func (tuo *TreatmentUpdateOne) SetEdgesOfUnpaybillsID(id int) *TreatmentUpdateOne {
	tuo.mutation.SetEdgesOfUnpaybillsID(id)
	return tuo
}

// SetNillableEdgesOfUnpaybillsID sets the EdgesOfUnpaybills edge to Unpaybill by id if the given value is not nil.
func (tuo *TreatmentUpdateOne) SetNillableEdgesOfUnpaybillsID(id *int) *TreatmentUpdateOne {
	if id != nil {
		tuo = tuo.SetEdgesOfUnpaybillsID(*id)
	}
	return tuo
}

// SetEdgesOfUnpaybills sets the EdgesOfUnpaybills edge to Unpaybill.
func (tuo *TreatmentUpdateOne) SetEdgesOfUnpaybills(u *Unpaybill) *TreatmentUpdateOne {
	return tuo.SetEdgesOfUnpaybillsID(u.ID)
}

// Mutation returns the TreatmentMutation object of the builder.
func (tuo *TreatmentUpdateOne) Mutation() *TreatmentMutation {
	return tuo.mutation
}

// ClearEdgesOfTypetreatment clears the EdgesOfTypetreatment edge to Typetreatment.
func (tuo *TreatmentUpdateOne) ClearEdgesOfTypetreatment() *TreatmentUpdateOne {
	tuo.mutation.ClearEdgesOfTypetreatment()
	return tuo
}

// ClearEdgesOfPatientrecord clears the EdgesOfPatientrecord edge to Patientrecord.
func (tuo *TreatmentUpdateOne) ClearEdgesOfPatientrecord() *TreatmentUpdateOne {
	tuo.mutation.ClearEdgesOfPatientrecord()
	return tuo
}

// ClearEdgesOfDoctor clears the EdgesOfDoctor edge to Doctor.
func (tuo *TreatmentUpdateOne) ClearEdgesOfDoctor() *TreatmentUpdateOne {
	tuo.mutation.ClearEdgesOfDoctor()
	return tuo
}

// ClearEdgesOfUnpaybills clears the EdgesOfUnpaybills edge to Unpaybill.
func (tuo *TreatmentUpdateOne) ClearEdgesOfUnpaybills() *TreatmentUpdateOne {
	tuo.mutation.ClearEdgesOfUnpaybills()
	return tuo
}

// Save executes the query and returns the updated entity.
func (tuo *TreatmentUpdateOne) Save(ctx context.Context) (*Treatment, error) {
	if v, ok := tuo.mutation.Symptom(); ok {
		if err := treatment.SymptomValidator(v); err != nil {
			return nil, &ValidationError{Name: "Symptom", err: fmt.Errorf("ent: validator failed for field \"Symptom\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Treat(); ok {
		if err := treatment.TreatValidator(v); err != nil {
			return nil, &ValidationError{Name: "Treat", err: fmt.Errorf("ent: validator failed for field \"Treat\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Medicine(); ok {
		if err := treatment.MedicineValidator(v); err != nil {
			return nil, &ValidationError{Name: "Medicine", err: fmt.Errorf("ent: validator failed for field \"Medicine\": %w", err)}
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
	if value, ok := tuo.mutation.Symptom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatment.FieldSymptom,
		})
	}
	if value, ok := tuo.mutation.Treat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatment.FieldTreat,
		})
	}
	if value, ok := tuo.mutation.Medicine(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatment.FieldMedicine,
		})
	}
	if value, ok := tuo.mutation.Datetreat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: treatment.FieldDatetreat,
		})
	}
	if tuo.mutation.EdgesOfTypetreatmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfTypetreatmentTable,
			Columns: []string{treatment.EdgesOfTypetreatmentColumn},
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
	if nodes := tuo.mutation.EdgesOfTypetreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfTypetreatmentTable,
			Columns: []string{treatment.EdgesOfTypetreatmentColumn},
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
	if tuo.mutation.EdgesOfPatientrecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfPatientrecordTable,
			Columns: []string{treatment.EdgesOfPatientrecordColumn},
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
	if nodes := tuo.mutation.EdgesOfPatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfPatientrecordTable,
			Columns: []string{treatment.EdgesOfPatientrecordColumn},
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
	if tuo.mutation.EdgesOfDoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfDoctorTable,
			Columns: []string{treatment.EdgesOfDoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.EdgesOfDoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.EdgesOfDoctorTable,
			Columns: []string{treatment.EdgesOfDoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.EdgesOfUnpaybillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   treatment.EdgesOfUnpaybillsTable,
			Columns: []string{treatment.EdgesOfUnpaybillsColumn},
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
	if nodes := tuo.mutation.EdgesOfUnpaybillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   treatment.EdgesOfUnpaybillsTable,
			Columns: []string{treatment.EdgesOfUnpaybillsColumn},
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
