// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/abilitypatientrights"
	"github.com/team10/app/ent/patientrightstype"
)

// AbilitypatientrightsCreate is the builder for creating a Abilitypatientrights entity.
type AbilitypatientrightsCreate struct {
	config
	mutation *AbilitypatientrightsMutation
	hooks    []Hook
}

// SetOperative sets the Operative field.
func (ac *AbilitypatientrightsCreate) SetOperative(i int) *AbilitypatientrightsCreate {
	ac.mutation.SetOperative(i)
	return ac
}

// SetMedicalSupplies sets the MedicalSupplies field.
func (ac *AbilitypatientrightsCreate) SetMedicalSupplies(i int) *AbilitypatientrightsCreate {
	ac.mutation.SetMedicalSupplies(i)
	return ac
}

// SetExamine sets the Examine field.
func (ac *AbilitypatientrightsCreate) SetExamine(i int) *AbilitypatientrightsCreate {
	ac.mutation.SetExamine(i)
	return ac
}

// AddEdgesOfAbilitypatientrightsPatientrightstypeIDs adds the EdgesOfAbilitypatientrightsPatientrightstype edge to Patientrightstype by ids.
func (ac *AbilitypatientrightsCreate) AddEdgesOfAbilitypatientrightsPatientrightstypeIDs(ids ...int) *AbilitypatientrightsCreate {
	ac.mutation.AddEdgesOfAbilitypatientrightsPatientrightstypeIDs(ids...)
	return ac
}

// AddEdgesOfAbilitypatientrightsPatientrightstype adds the EdgesOfAbilitypatientrightsPatientrightstype edges to Patientrightstype.
func (ac *AbilitypatientrightsCreate) AddEdgesOfAbilitypatientrightsPatientrightstype(p ...*Patientrightstype) *AbilitypatientrightsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddEdgesOfAbilitypatientrightsPatientrightstypeIDs(ids...)
}

// Mutation returns the AbilitypatientrightsMutation object of the builder.
func (ac *AbilitypatientrightsCreate) Mutation() *AbilitypatientrightsMutation {
	return ac.mutation
}

// Save creates the Abilitypatientrights in the database.
func (ac *AbilitypatientrightsCreate) Save(ctx context.Context) (*Abilitypatientrights, error) {
	if _, ok := ac.mutation.Operative(); !ok {
		return nil, &ValidationError{Name: "Operative", err: errors.New("ent: missing required field \"Operative\"")}
	}
	if _, ok := ac.mutation.MedicalSupplies(); !ok {
		return nil, &ValidationError{Name: "MedicalSupplies", err: errors.New("ent: missing required field \"MedicalSupplies\"")}
	}
	if _, ok := ac.mutation.Examine(); !ok {
		return nil, &ValidationError{Name: "Examine", err: errors.New("ent: missing required field \"Examine\"")}
	}
	var (
		err  error
		node *Abilitypatientrights
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AbilitypatientrightsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AbilitypatientrightsCreate) SaveX(ctx context.Context) *Abilitypatientrights {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AbilitypatientrightsCreate) sqlSave(ctx context.Context) (*Abilitypatientrights, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}

func (ac *AbilitypatientrightsCreate) createSpec() (*Abilitypatientrights, *sqlgraph.CreateSpec) {
	var (
		a     = &Abilitypatientrights{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: abilitypatientrights.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: abilitypatientrights.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Operative(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldOperative,
		})
		a.Operative = value
	}
	if value, ok := ac.mutation.MedicalSupplies(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldMedicalSupplies,
		})
		a.MedicalSupplies = value
	}
	if value, ok := ac.mutation.Examine(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldExamine,
		})
		a.Examine = value
	}
	if nodes := ac.mutation.EdgesOfAbilitypatientrightsPatientrightstypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilitypatientrights.EdgesOfAbilitypatientrightsPatientrightstypeTable,
			Columns: []string{abilitypatientrights.EdgesOfAbilitypatientrightsPatientrightstypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrightstype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return a, _spec
}
