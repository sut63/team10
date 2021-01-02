// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/patientrights"
	"github.com/team10/app/ent/user"
)

// MedicalrecordstaffCreate is the builder for creating a Medicalrecordstaff entity.
type MedicalrecordstaffCreate struct {
	config
	mutation *MedicalrecordstaffMutation
	hooks    []Hook
}

// SetName sets the Name field.
func (mc *MedicalrecordstaffCreate) SetName(s string) *MedicalrecordstaffCreate {
	mc.mutation.SetName(s)
	return mc
}

// AddPatientrecordIDs adds the patientrecord edge to Patientrecord by ids.
func (mc *MedicalrecordstaffCreate) AddPatientrecordIDs(ids ...int) *MedicalrecordstaffCreate {
	mc.mutation.AddPatientrecordIDs(ids...)
	return mc
}

// AddPatientrecord adds the patientrecord edges to Patientrecord.
func (mc *MedicalrecordstaffCreate) AddPatientrecord(p ...*Patientrecord) *MedicalrecordstaffCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddPatientrecordIDs(ids...)
}

// AddMedicalrecordstaffPatientrightIDs adds the MedicalrecordstaffPatientrights edge to Patientrights by ids.
func (mc *MedicalrecordstaffCreate) AddMedicalrecordstaffPatientrightIDs(ids ...int) *MedicalrecordstaffCreate {
	mc.mutation.AddMedicalrecordstaffPatientrightIDs(ids...)
	return mc
}

// AddMedicalrecordstaffPatientrights adds the MedicalrecordstaffPatientrights edges to Patientrights.
func (mc *MedicalrecordstaffCreate) AddMedicalrecordstaffPatientrights(p ...*Patientrights) *MedicalrecordstaffCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddMedicalrecordstaffPatientrightIDs(ids...)
}

// SetUserID sets the user edge to User by id.
func (mc *MedicalrecordstaffCreate) SetUserID(id int) *MedicalrecordstaffCreate {
	mc.mutation.SetUserID(id)
	return mc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (mc *MedicalrecordstaffCreate) SetNillableUserID(id *int) *MedicalrecordstaffCreate {
	if id != nil {
		mc = mc.SetUserID(*id)
	}
	return mc
}

// SetUser sets the user edge to User.
func (mc *MedicalrecordstaffCreate) SetUser(u *User) *MedicalrecordstaffCreate {
	return mc.SetUserID(u.ID)
}

// Mutation returns the MedicalrecordstaffMutation object of the builder.
func (mc *MedicalrecordstaffCreate) Mutation() *MedicalrecordstaffMutation {
	return mc.mutation
}

// Save creates the Medicalrecordstaff in the database.
func (mc *MedicalrecordstaffCreate) Save(ctx context.Context) (*Medicalrecordstaff, error) {
	if _, ok := mc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "Name", err: errors.New("ent: missing required field \"Name\"")}
	}
	var (
		err  error
		node *Medicalrecordstaff
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicalrecordstaffMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MedicalrecordstaffCreate) SaveX(ctx context.Context) *Medicalrecordstaff {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MedicalrecordstaffCreate) sqlSave(ctx context.Context) (*Medicalrecordstaff, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MedicalrecordstaffCreate) createSpec() (*Medicalrecordstaff, *sqlgraph.CreateSpec) {
	var (
		m     = &Medicalrecordstaff{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: medicalrecordstaff.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalrecordstaff.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicalrecordstaff.FieldName,
		})
		m.Name = value
	}
	if nodes := mc.mutation.PatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalrecordstaff.PatientrecordTable,
			Columns: []string{medicalrecordstaff.PatientrecordColumn},
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
	if nodes := mc.mutation.MedicalrecordstaffPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalrecordstaff.MedicalrecordstaffPatientrightsTable,
			Columns: []string{medicalrecordstaff.MedicalrecordstaffPatientrightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrights.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   medicalrecordstaff.UserTable,
			Columns: []string{medicalrecordstaff.UserColumn},
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
	return m, _spec
}
