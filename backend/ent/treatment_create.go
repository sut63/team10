// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/theuo/app/ent/doctorinfo"
	"github.com/theuo/app/ent/patientrecord"
	"github.com/theuo/app/ent/treatment"
	"github.com/theuo/app/ent/typetreatment"
	"github.com/theuo/app/ent/unpaybill"
)

// TreatmentCreate is the builder for creating a Treatment entity.
type TreatmentCreate struct {
	config
	mutation *TreatmentMutation
	hooks    []Hook
}

// SetTreatment sets the treatment field.
func (tc *TreatmentCreate) SetTreatment(s string) *TreatmentCreate {
	tc.mutation.SetTreatment(s)
	return tc
}

// SetDatetime sets the datetime field.
func (tc *TreatmentCreate) SetDatetime(t time.Time) *TreatmentCreate {
	tc.mutation.SetDatetime(t)
	return tc
}

// SetTypetreatmentID sets the typetreatment edge to Typetreatment by id.
func (tc *TreatmentCreate) SetTypetreatmentID(id int) *TreatmentCreate {
	tc.mutation.SetTypetreatmentID(id)
	return tc
}

// SetNillableTypetreatmentID sets the typetreatment edge to Typetreatment by id if the given value is not nil.
func (tc *TreatmentCreate) SetNillableTypetreatmentID(id *int) *TreatmentCreate {
	if id != nil {
		tc = tc.SetTypetreatmentID(*id)
	}
	return tc
}

// SetTypetreatment sets the typetreatment edge to Typetreatment.
func (tc *TreatmentCreate) SetTypetreatment(t *Typetreatment) *TreatmentCreate {
	return tc.SetTypetreatmentID(t.ID)
}

// SetPatientrecordID sets the patientrecord edge to Patientrecord by id.
func (tc *TreatmentCreate) SetPatientrecordID(id int) *TreatmentCreate {
	tc.mutation.SetPatientrecordID(id)
	return tc
}

// SetNillablePatientrecordID sets the patientrecord edge to Patientrecord by id if the given value is not nil.
func (tc *TreatmentCreate) SetNillablePatientrecordID(id *int) *TreatmentCreate {
	if id != nil {
		tc = tc.SetPatientrecordID(*id)
	}
	return tc
}

// SetPatientrecord sets the patientrecord edge to Patientrecord.
func (tc *TreatmentCreate) SetPatientrecord(p *Patientrecord) *TreatmentCreate {
	return tc.SetPatientrecordID(p.ID)
}

// SetDoctorinfoID sets the doctorinfo edge to Doctorinfo by id.
func (tc *TreatmentCreate) SetDoctorinfoID(id int) *TreatmentCreate {
	tc.mutation.SetDoctorinfoID(id)
	return tc
}

// SetNillableDoctorinfoID sets the doctorinfo edge to Doctorinfo by id if the given value is not nil.
func (tc *TreatmentCreate) SetNillableDoctorinfoID(id *int) *TreatmentCreate {
	if id != nil {
		tc = tc.SetDoctorinfoID(*id)
	}
	return tc
}

// SetDoctorinfo sets the doctorinfo edge to Doctorinfo.
func (tc *TreatmentCreate) SetDoctorinfo(d *Doctorinfo) *TreatmentCreate {
	return tc.SetDoctorinfoID(d.ID)
}

// SetUnpaybillsID sets the unpaybills edge to Unpaybill by id.
func (tc *TreatmentCreate) SetUnpaybillsID(id int) *TreatmentCreate {
	tc.mutation.SetUnpaybillsID(id)
	return tc
}

// SetNillableUnpaybillsID sets the unpaybills edge to Unpaybill by id if the given value is not nil.
func (tc *TreatmentCreate) SetNillableUnpaybillsID(id *int) *TreatmentCreate {
	if id != nil {
		tc = tc.SetUnpaybillsID(*id)
	}
	return tc
}

// SetUnpaybills sets the unpaybills edge to Unpaybill.
func (tc *TreatmentCreate) SetUnpaybills(u *Unpaybill) *TreatmentCreate {
	return tc.SetUnpaybillsID(u.ID)
}

// Mutation returns the TreatmentMutation object of the builder.
func (tc *TreatmentCreate) Mutation() *TreatmentMutation {
	return tc.mutation
}

// Save creates the Treatment in the database.
func (tc *TreatmentCreate) Save(ctx context.Context) (*Treatment, error) {
	if _, ok := tc.mutation.Treatment(); !ok {
		return nil, &ValidationError{Name: "treatment", err: errors.New("ent: missing required field \"treatment\"")}
	}
	if v, ok := tc.mutation.Treatment(); ok {
		if err := treatment.TreatmentValidator(v); err != nil {
			return nil, &ValidationError{Name: "treatment", err: fmt.Errorf("ent: validator failed for field \"treatment\": %w", err)}
		}
	}
	if _, ok := tc.mutation.Datetime(); !ok {
		return nil, &ValidationError{Name: "datetime", err: errors.New("ent: missing required field \"datetime\"")}
	}
	var (
		err  error
		node *Treatment
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TreatmentCreate) SaveX(ctx context.Context) *Treatment {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TreatmentCreate) sqlSave(ctx context.Context) (*Treatment, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TreatmentCreate) createSpec() (*Treatment, *sqlgraph.CreateSpec) {
	var (
		t     = &Treatment{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: treatment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: treatment.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Treatment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatment.FieldTreatment,
		})
		t.Treatment = value
	}
	if value, ok := tc.mutation.Datetime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: treatment.FieldDatetime,
		})
		t.Datetime = value
	}
	if nodes := tc.mutation.TypetreatmentIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.PatientrecordIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.DoctorinfoIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UnpaybillsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return t, _spec
}
