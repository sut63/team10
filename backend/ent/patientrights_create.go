// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/abilitypatientrights"
	"github.com/team10/app/ent/insurance"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/patientrights"
)

// PatientrightsCreate is the builder for creating a Patientrights entity.
type PatientrightsCreate struct {
	config
	mutation *PatientrightsMutation
	hooks    []Hook
}

// SetPermissionDate sets the PermissionDate field.
func (pc *PatientrightsCreate) SetPermissionDate(t time.Time) *PatientrightsCreate {
	pc.mutation.SetPermissionDate(t)
	return pc
}

// SetPermission sets the Permission field.
func (pc *PatientrightsCreate) SetPermission(s string) *PatientrightsCreate {
	pc.mutation.SetPermission(s)
	return pc
}

// SetPermissionArea sets the PermissionArea field.
func (pc *PatientrightsCreate) SetPermissionArea(s string) *PatientrightsCreate {
	pc.mutation.SetPermissionArea(s)
	return pc
}

// SetResponsible sets the Responsible field.
func (pc *PatientrightsCreate) SetResponsible(s string) *PatientrightsCreate {
	pc.mutation.SetResponsible(s)
	return pc
}

// SetEdgesOfPatientrightsAbilitypatientrightsID sets the EdgesOfPatientrightsAbilitypatientrights edge to Abilitypatientrights by id.
func (pc *PatientrightsCreate) SetEdgesOfPatientrightsAbilitypatientrightsID(id int) *PatientrightsCreate {
	pc.mutation.SetEdgesOfPatientrightsAbilitypatientrightsID(id)
	return pc
}

// SetNillableEdgesOfPatientrightsAbilitypatientrightsID sets the EdgesOfPatientrightsAbilitypatientrights edge to Abilitypatientrights by id if the given value is not nil.
func (pc *PatientrightsCreate) SetNillableEdgesOfPatientrightsAbilitypatientrightsID(id *int) *PatientrightsCreate {
	if id != nil {
		pc = pc.SetEdgesOfPatientrightsAbilitypatientrightsID(*id)
	}
	return pc
}

// SetEdgesOfPatientrightsAbilitypatientrights sets the EdgesOfPatientrightsAbilitypatientrights edge to Abilitypatientrights.
func (pc *PatientrightsCreate) SetEdgesOfPatientrightsAbilitypatientrights(a *Abilitypatientrights) *PatientrightsCreate {
	return pc.SetEdgesOfPatientrightsAbilitypatientrightsID(a.ID)
}

// SetEdgesOfPatientrightsInsuranceID sets the EdgesOfPatientrightsInsurance edge to Insurance by id.
func (pc *PatientrightsCreate) SetEdgesOfPatientrightsInsuranceID(id int) *PatientrightsCreate {
	pc.mutation.SetEdgesOfPatientrightsInsuranceID(id)
	return pc
}

// SetNillableEdgesOfPatientrightsInsuranceID sets the EdgesOfPatientrightsInsurance edge to Insurance by id if the given value is not nil.
func (pc *PatientrightsCreate) SetNillableEdgesOfPatientrightsInsuranceID(id *int) *PatientrightsCreate {
	if id != nil {
		pc = pc.SetEdgesOfPatientrightsInsuranceID(*id)
	}
	return pc
}

// SetEdgesOfPatientrightsInsurance sets the EdgesOfPatientrightsInsurance edge to Insurance.
func (pc *PatientrightsCreate) SetEdgesOfPatientrightsInsurance(i *Insurance) *PatientrightsCreate {
	return pc.SetEdgesOfPatientrightsInsuranceID(i.ID)
}

// SetEdgesOfPatientrightsPatientrecordID sets the EdgesOfPatientrightsPatientrecord edge to Patientrecord by id.
func (pc *PatientrightsCreate) SetEdgesOfPatientrightsPatientrecordID(id int) *PatientrightsCreate {
	pc.mutation.SetEdgesOfPatientrightsPatientrecordID(id)
	return pc
}

// SetNillableEdgesOfPatientrightsPatientrecordID sets the EdgesOfPatientrightsPatientrecord edge to Patientrecord by id if the given value is not nil.
func (pc *PatientrightsCreate) SetNillableEdgesOfPatientrightsPatientrecordID(id *int) *PatientrightsCreate {
	if id != nil {
		pc = pc.SetEdgesOfPatientrightsPatientrecordID(*id)
	}
	return pc
}

// SetEdgesOfPatientrightsPatientrecord sets the EdgesOfPatientrightsPatientrecord edge to Patientrecord.
func (pc *PatientrightsCreate) SetEdgesOfPatientrightsPatientrecord(p *Patientrecord) *PatientrightsCreate {
	return pc.SetEdgesOfPatientrightsPatientrecordID(p.ID)
}

// SetEdgesOfPatientrightsMedicalrecordstaffID sets the EdgesOfPatientrightsMedicalrecordstaff edge to Medicalrecordstaff by id.
func (pc *PatientrightsCreate) SetEdgesOfPatientrightsMedicalrecordstaffID(id int) *PatientrightsCreate {
	pc.mutation.SetEdgesOfPatientrightsMedicalrecordstaffID(id)
	return pc
}

// SetNillableEdgesOfPatientrightsMedicalrecordstaffID sets the EdgesOfPatientrightsMedicalrecordstaff edge to Medicalrecordstaff by id if the given value is not nil.
func (pc *PatientrightsCreate) SetNillableEdgesOfPatientrightsMedicalrecordstaffID(id *int) *PatientrightsCreate {
	if id != nil {
		pc = pc.SetEdgesOfPatientrightsMedicalrecordstaffID(*id)
	}
	return pc
}

// SetEdgesOfPatientrightsMedicalrecordstaff sets the EdgesOfPatientrightsMedicalrecordstaff edge to Medicalrecordstaff.
func (pc *PatientrightsCreate) SetEdgesOfPatientrightsMedicalrecordstaff(m *Medicalrecordstaff) *PatientrightsCreate {
	return pc.SetEdgesOfPatientrightsMedicalrecordstaffID(m.ID)
}

// Mutation returns the PatientrightsMutation object of the builder.
func (pc *PatientrightsCreate) Mutation() *PatientrightsMutation {
	return pc.mutation
}

// Save creates the Patientrights in the database.
func (pc *PatientrightsCreate) Save(ctx context.Context) (*Patientrights, error) {
	if _, ok := pc.mutation.PermissionDate(); !ok {
		return nil, &ValidationError{Name: "PermissionDate", err: errors.New("ent: missing required field \"PermissionDate\"")}
	}
	if _, ok := pc.mutation.Permission(); !ok {
		return nil, &ValidationError{Name: "Permission", err: errors.New("ent: missing required field \"Permission\"")}
	}
	if v, ok := pc.mutation.Permission(); ok {
		if err := patientrights.PermissionValidator(v); err != nil {
			return nil, &ValidationError{Name: "Permission", err: fmt.Errorf("ent: validator failed for field \"Permission\": %w", err)}
		}
	}
	if _, ok := pc.mutation.PermissionArea(); !ok {
		return nil, &ValidationError{Name: "PermissionArea", err: errors.New("ent: missing required field \"PermissionArea\"")}
	}
	if v, ok := pc.mutation.PermissionArea(); ok {
		if err := patientrights.PermissionAreaValidator(v); err != nil {
			return nil, &ValidationError{Name: "PermissionArea", err: fmt.Errorf("ent: validator failed for field \"PermissionArea\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Responsible(); !ok {
		return nil, &ValidationError{Name: "Responsible", err: errors.New("ent: missing required field \"Responsible\"")}
	}
	if v, ok := pc.mutation.Responsible(); ok {
		if err := patientrights.ResponsibleValidator(v); err != nil {
			return nil, &ValidationError{Name: "Responsible", err: fmt.Errorf("ent: validator failed for field \"Responsible\": %w", err)}
		}
	}
	var (
		err  error
		node *Patientrights
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientrightsMutation)
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
func (pc *PatientrightsCreate) SaveX(ctx context.Context) *Patientrights {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PatientrightsCreate) sqlSave(ctx context.Context) (*Patientrights, error) {
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

func (pc *PatientrightsCreate) createSpec() (*Patientrights, *sqlgraph.CreateSpec) {
	var (
		pa    = &Patientrights{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patientrights.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrights.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PermissionDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patientrights.FieldPermissionDate,
		})
		pa.PermissionDate = value
	}
	if value, ok := pc.mutation.Permission(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrights.FieldPermission,
		})
		pa.Permission = value
	}
	if value, ok := pc.mutation.PermissionArea(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrights.FieldPermissionArea,
		})
		pa.PermissionArea = value
	}
	if value, ok := pc.mutation.Responsible(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrights.FieldResponsible,
		})
		pa.Responsible = value
	}
	if nodes := pc.mutation.EdgesOfPatientrightsAbilitypatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrights.EdgesOfPatientrightsAbilitypatientrightsTable,
			Columns: []string{patientrights.EdgesOfPatientrightsAbilitypatientrightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: abilitypatientrights.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.EdgesOfPatientrightsInsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrights.EdgesOfPatientrightsInsuranceTable,
			Columns: []string{patientrights.EdgesOfPatientrightsInsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: insurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.EdgesOfPatientrightsPatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrights.EdgesOfPatientrightsPatientrecordTable,
			Columns: []string{patientrights.EdgesOfPatientrightsPatientrecordColumn},
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
	if nodes := pc.mutation.EdgesOfPatientrightsMedicalrecordstaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrights.EdgesOfPatientrightsMedicalrecordstaffTable,
			Columns: []string{patientrights.EdgesOfPatientrightsMedicalrecordstaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalrecordstaff.FieldID,
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
