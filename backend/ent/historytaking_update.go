// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/symptomseverity"
)

// HistorytakingUpdate is the builder for updating Historytaking entities.
type HistorytakingUpdate struct {
	config
	hooks      []Hook
	mutation   *HistorytakingMutation
	predicates []predicate.Historytaking
}

// Where adds a new predicate for the builder.
func (hu *HistorytakingUpdate) Where(ps ...predicate.Historytaking) *HistorytakingUpdate {
	hu.predicates = append(hu.predicates, ps...)
	return hu
}

// SetHight sets the hight field.
func (hu *HistorytakingUpdate) SetHight(f float32) *HistorytakingUpdate {
	hu.mutation.ResetHight()
	hu.mutation.SetHight(f)
	return hu
}

// AddHight adds f to hight.
func (hu *HistorytakingUpdate) AddHight(f float32) *HistorytakingUpdate {
	hu.mutation.AddHight(f)
	return hu
}

// SetWeight sets the weight field.
func (hu *HistorytakingUpdate) SetWeight(f float32) *HistorytakingUpdate {
	hu.mutation.ResetWeight()
	hu.mutation.SetWeight(f)
	return hu
}

// AddWeight adds f to weight.
func (hu *HistorytakingUpdate) AddWeight(f float32) *HistorytakingUpdate {
	hu.mutation.AddWeight(f)
	return hu
}

// SetTemp sets the temp field.
func (hu *HistorytakingUpdate) SetTemp(f float32) *HistorytakingUpdate {
	hu.mutation.ResetTemp()
	hu.mutation.SetTemp(f)
	return hu
}

// AddTemp adds f to temp.
func (hu *HistorytakingUpdate) AddTemp(f float32) *HistorytakingUpdate {
	hu.mutation.AddTemp(f)
	return hu
}

// SetPulse sets the pulse field.
func (hu *HistorytakingUpdate) SetPulse(i int) *HistorytakingUpdate {
	hu.mutation.ResetPulse()
	hu.mutation.SetPulse(i)
	return hu
}

// AddPulse adds i to pulse.
func (hu *HistorytakingUpdate) AddPulse(i int) *HistorytakingUpdate {
	hu.mutation.AddPulse(i)
	return hu
}

// SetRespiration sets the respiration field.
func (hu *HistorytakingUpdate) SetRespiration(i int) *HistorytakingUpdate {
	hu.mutation.ResetRespiration()
	hu.mutation.SetRespiration(i)
	return hu
}

// AddRespiration adds i to respiration.
func (hu *HistorytakingUpdate) AddRespiration(i int) *HistorytakingUpdate {
	hu.mutation.AddRespiration(i)
	return hu
}

// SetBp sets the bp field.
func (hu *HistorytakingUpdate) SetBp(i int) *HistorytakingUpdate {
	hu.mutation.ResetBp()
	hu.mutation.SetBp(i)
	return hu
}

// AddBp adds i to bp.
func (hu *HistorytakingUpdate) AddBp(i int) *HistorytakingUpdate {
	hu.mutation.AddBp(i)
	return hu
}

// SetOxygen sets the oxygen field.
func (hu *HistorytakingUpdate) SetOxygen(s string) *HistorytakingUpdate {
	hu.mutation.SetOxygen(s)
	return hu
}

// SetSymptom sets the symptom field.
func (hu *HistorytakingUpdate) SetSymptom(s string) *HistorytakingUpdate {
	hu.mutation.SetSymptom(s)
	return hu
}

// SetDatetime sets the datetime field.
func (hu *HistorytakingUpdate) SetDatetime(t time.Time) *HistorytakingUpdate {
	hu.mutation.SetDatetime(t)
	return hu
}

// SetEdgesOfNurseID sets the EdgesOfNurse edge to Nurse by id.
func (hu *HistorytakingUpdate) SetEdgesOfNurseID(id int) *HistorytakingUpdate {
	hu.mutation.SetEdgesOfNurseID(id)
	return hu
}

// SetNillableEdgesOfNurseID sets the EdgesOfNurse edge to Nurse by id if the given value is not nil.
func (hu *HistorytakingUpdate) SetNillableEdgesOfNurseID(id *int) *HistorytakingUpdate {
	if id != nil {
		hu = hu.SetEdgesOfNurseID(*id)
	}
	return hu
}

// SetEdgesOfNurse sets the EdgesOfNurse edge to Nurse.
func (hu *HistorytakingUpdate) SetEdgesOfNurse(n *Nurse) *HistorytakingUpdate {
	return hu.SetEdgesOfNurseID(n.ID)
}

// SetEdgesOfDepartmentID sets the EdgesOfDepartment edge to Department by id.
func (hu *HistorytakingUpdate) SetEdgesOfDepartmentID(id int) *HistorytakingUpdate {
	hu.mutation.SetEdgesOfDepartmentID(id)
	return hu
}

// SetNillableEdgesOfDepartmentID sets the EdgesOfDepartment edge to Department by id if the given value is not nil.
func (hu *HistorytakingUpdate) SetNillableEdgesOfDepartmentID(id *int) *HistorytakingUpdate {
	if id != nil {
		hu = hu.SetEdgesOfDepartmentID(*id)
	}
	return hu
}

// SetEdgesOfDepartment sets the EdgesOfDepartment edge to Department.
func (hu *HistorytakingUpdate) SetEdgesOfDepartment(d *Department) *HistorytakingUpdate {
	return hu.SetEdgesOfDepartmentID(d.ID)
}

// SetEdgesOfSymptomseverityID sets the EdgesOfSymptomseverity edge to Symptomseverity by id.
func (hu *HistorytakingUpdate) SetEdgesOfSymptomseverityID(id int) *HistorytakingUpdate {
	hu.mutation.SetEdgesOfSymptomseverityID(id)
	return hu
}

// SetNillableEdgesOfSymptomseverityID sets the EdgesOfSymptomseverity edge to Symptomseverity by id if the given value is not nil.
func (hu *HistorytakingUpdate) SetNillableEdgesOfSymptomseverityID(id *int) *HistorytakingUpdate {
	if id != nil {
		hu = hu.SetEdgesOfSymptomseverityID(*id)
	}
	return hu
}

// SetEdgesOfSymptomseverity sets the EdgesOfSymptomseverity edge to Symptomseverity.
func (hu *HistorytakingUpdate) SetEdgesOfSymptomseverity(s *Symptomseverity) *HistorytakingUpdate {
	return hu.SetEdgesOfSymptomseverityID(s.ID)
}

// SetEdgesOfPatientrecordID sets the EdgesOfPatientrecord edge to Patientrecord by id.
func (hu *HistorytakingUpdate) SetEdgesOfPatientrecordID(id int) *HistorytakingUpdate {
	hu.mutation.SetEdgesOfPatientrecordID(id)
	return hu
}

// SetNillableEdgesOfPatientrecordID sets the EdgesOfPatientrecord edge to Patientrecord by id if the given value is not nil.
func (hu *HistorytakingUpdate) SetNillableEdgesOfPatientrecordID(id *int) *HistorytakingUpdate {
	if id != nil {
		hu = hu.SetEdgesOfPatientrecordID(*id)
	}
	return hu
}

// SetEdgesOfPatientrecord sets the EdgesOfPatientrecord edge to Patientrecord.
func (hu *HistorytakingUpdate) SetEdgesOfPatientrecord(p *Patientrecord) *HistorytakingUpdate {
	return hu.SetEdgesOfPatientrecordID(p.ID)
}

// Mutation returns the HistorytakingMutation object of the builder.
func (hu *HistorytakingUpdate) Mutation() *HistorytakingMutation {
	return hu.mutation
}

// ClearEdgesOfNurse clears the EdgesOfNurse edge to Nurse.
func (hu *HistorytakingUpdate) ClearEdgesOfNurse() *HistorytakingUpdate {
	hu.mutation.ClearEdgesOfNurse()
	return hu
}

// ClearEdgesOfDepartment clears the EdgesOfDepartment edge to Department.
func (hu *HistorytakingUpdate) ClearEdgesOfDepartment() *HistorytakingUpdate {
	hu.mutation.ClearEdgesOfDepartment()
	return hu
}

// ClearEdgesOfSymptomseverity clears the EdgesOfSymptomseverity edge to Symptomseverity.
func (hu *HistorytakingUpdate) ClearEdgesOfSymptomseverity() *HistorytakingUpdate {
	hu.mutation.ClearEdgesOfSymptomseverity()
	return hu
}

// ClearEdgesOfPatientrecord clears the EdgesOfPatientrecord edge to Patientrecord.
func (hu *HistorytakingUpdate) ClearEdgesOfPatientrecord() *HistorytakingUpdate {
	hu.mutation.ClearEdgesOfPatientrecord()
	return hu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (hu *HistorytakingUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(hu.hooks) == 0 {
		affected, err = hu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HistorytakingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hu.mutation = mutation
			affected, err = hu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hu.hooks) - 1; i >= 0; i-- {
			mut = hu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HistorytakingUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HistorytakingUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HistorytakingUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hu *HistorytakingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   historytaking.Table,
			Columns: historytaking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: historytaking.FieldID,
			},
		},
	}
	if ps := hu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Hight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldHight,
		})
	}
	if value, ok := hu.mutation.AddedHight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldHight,
		})
	}
	if value, ok := hu.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldWeight,
		})
	}
	if value, ok := hu.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldWeight,
		})
	}
	if value, ok := hu.mutation.Temp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldTemp,
		})
	}
	if value, ok := hu.mutation.AddedTemp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldTemp,
		})
	}
	if value, ok := hu.mutation.Pulse(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldPulse,
		})
	}
	if value, ok := hu.mutation.AddedPulse(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldPulse,
		})
	}
	if value, ok := hu.mutation.Respiration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldRespiration,
		})
	}
	if value, ok := hu.mutation.AddedRespiration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldRespiration,
		})
	}
	if value, ok := hu.mutation.Bp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldBp,
		})
	}
	if value, ok := hu.mutation.AddedBp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldBp,
		})
	}
	if value, ok := hu.mutation.Oxygen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: historytaking.FieldOxygen,
		})
	}
	if value, ok := hu.mutation.Symptom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: historytaking.FieldSymptom,
		})
	}
	if value, ok := hu.mutation.Datetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: historytaking.FieldDatetime,
		})
	}
	if hu.mutation.EdgesOfNurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfNurseTable,
			Columns: []string{historytaking.EdgesOfNurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.EdgesOfNurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfNurseTable,
			Columns: []string{historytaking.EdgesOfNurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.EdgesOfDepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfDepartmentTable,
			Columns: []string{historytaking.EdgesOfDepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.EdgesOfDepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfDepartmentTable,
			Columns: []string{historytaking.EdgesOfDepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.EdgesOfSymptomseverityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfSymptomseverityTable,
			Columns: []string{historytaking.EdgesOfSymptomseverityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: symptomseverity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.EdgesOfSymptomseverityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfSymptomseverityTable,
			Columns: []string{historytaking.EdgesOfSymptomseverityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: symptomseverity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.EdgesOfPatientrecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfPatientrecordTable,
			Columns: []string{historytaking.EdgesOfPatientrecordColumn},
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
	if nodes := hu.mutation.EdgesOfPatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfPatientrecordTable,
			Columns: []string{historytaking.EdgesOfPatientrecordColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{historytaking.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// HistorytakingUpdateOne is the builder for updating a single Historytaking entity.
type HistorytakingUpdateOne struct {
	config
	hooks    []Hook
	mutation *HistorytakingMutation
}

// SetHight sets the hight field.
func (huo *HistorytakingUpdateOne) SetHight(f float32) *HistorytakingUpdateOne {
	huo.mutation.ResetHight()
	huo.mutation.SetHight(f)
	return huo
}

// AddHight adds f to hight.
func (huo *HistorytakingUpdateOne) AddHight(f float32) *HistorytakingUpdateOne {
	huo.mutation.AddHight(f)
	return huo
}

// SetWeight sets the weight field.
func (huo *HistorytakingUpdateOne) SetWeight(f float32) *HistorytakingUpdateOne {
	huo.mutation.ResetWeight()
	huo.mutation.SetWeight(f)
	return huo
}

// AddWeight adds f to weight.
func (huo *HistorytakingUpdateOne) AddWeight(f float32) *HistorytakingUpdateOne {
	huo.mutation.AddWeight(f)
	return huo
}

// SetTemp sets the temp field.
func (huo *HistorytakingUpdateOne) SetTemp(f float32) *HistorytakingUpdateOne {
	huo.mutation.ResetTemp()
	huo.mutation.SetTemp(f)
	return huo
}

// AddTemp adds f to temp.
func (huo *HistorytakingUpdateOne) AddTemp(f float32) *HistorytakingUpdateOne {
	huo.mutation.AddTemp(f)
	return huo
}

// SetPulse sets the pulse field.
func (huo *HistorytakingUpdateOne) SetPulse(i int) *HistorytakingUpdateOne {
	huo.mutation.ResetPulse()
	huo.mutation.SetPulse(i)
	return huo
}

// AddPulse adds i to pulse.
func (huo *HistorytakingUpdateOne) AddPulse(i int) *HistorytakingUpdateOne {
	huo.mutation.AddPulse(i)
	return huo
}

// SetRespiration sets the respiration field.
func (huo *HistorytakingUpdateOne) SetRespiration(i int) *HistorytakingUpdateOne {
	huo.mutation.ResetRespiration()
	huo.mutation.SetRespiration(i)
	return huo
}

// AddRespiration adds i to respiration.
func (huo *HistorytakingUpdateOne) AddRespiration(i int) *HistorytakingUpdateOne {
	huo.mutation.AddRespiration(i)
	return huo
}

// SetBp sets the bp field.
func (huo *HistorytakingUpdateOne) SetBp(i int) *HistorytakingUpdateOne {
	huo.mutation.ResetBp()
	huo.mutation.SetBp(i)
	return huo
}

// AddBp adds i to bp.
func (huo *HistorytakingUpdateOne) AddBp(i int) *HistorytakingUpdateOne {
	huo.mutation.AddBp(i)
	return huo
}

// SetOxygen sets the oxygen field.
func (huo *HistorytakingUpdateOne) SetOxygen(s string) *HistorytakingUpdateOne {
	huo.mutation.SetOxygen(s)
	return huo
}

// SetSymptom sets the symptom field.
func (huo *HistorytakingUpdateOne) SetSymptom(s string) *HistorytakingUpdateOne {
	huo.mutation.SetSymptom(s)
	return huo
}

// SetDatetime sets the datetime field.
func (huo *HistorytakingUpdateOne) SetDatetime(t time.Time) *HistorytakingUpdateOne {
	huo.mutation.SetDatetime(t)
	return huo
}

// SetEdgesOfNurseID sets the EdgesOfNurse edge to Nurse by id.
func (huo *HistorytakingUpdateOne) SetEdgesOfNurseID(id int) *HistorytakingUpdateOne {
	huo.mutation.SetEdgesOfNurseID(id)
	return huo
}

// SetNillableEdgesOfNurseID sets the EdgesOfNurse edge to Nurse by id if the given value is not nil.
func (huo *HistorytakingUpdateOne) SetNillableEdgesOfNurseID(id *int) *HistorytakingUpdateOne {
	if id != nil {
		huo = huo.SetEdgesOfNurseID(*id)
	}
	return huo
}

// SetEdgesOfNurse sets the EdgesOfNurse edge to Nurse.
func (huo *HistorytakingUpdateOne) SetEdgesOfNurse(n *Nurse) *HistorytakingUpdateOne {
	return huo.SetEdgesOfNurseID(n.ID)
}

// SetEdgesOfDepartmentID sets the EdgesOfDepartment edge to Department by id.
func (huo *HistorytakingUpdateOne) SetEdgesOfDepartmentID(id int) *HistorytakingUpdateOne {
	huo.mutation.SetEdgesOfDepartmentID(id)
	return huo
}

// SetNillableEdgesOfDepartmentID sets the EdgesOfDepartment edge to Department by id if the given value is not nil.
func (huo *HistorytakingUpdateOne) SetNillableEdgesOfDepartmentID(id *int) *HistorytakingUpdateOne {
	if id != nil {
		huo = huo.SetEdgesOfDepartmentID(*id)
	}
	return huo
}

// SetEdgesOfDepartment sets the EdgesOfDepartment edge to Department.
func (huo *HistorytakingUpdateOne) SetEdgesOfDepartment(d *Department) *HistorytakingUpdateOne {
	return huo.SetEdgesOfDepartmentID(d.ID)
}

// SetEdgesOfSymptomseverityID sets the EdgesOfSymptomseverity edge to Symptomseverity by id.
func (huo *HistorytakingUpdateOne) SetEdgesOfSymptomseverityID(id int) *HistorytakingUpdateOne {
	huo.mutation.SetEdgesOfSymptomseverityID(id)
	return huo
}

// SetNillableEdgesOfSymptomseverityID sets the EdgesOfSymptomseverity edge to Symptomseverity by id if the given value is not nil.
func (huo *HistorytakingUpdateOne) SetNillableEdgesOfSymptomseverityID(id *int) *HistorytakingUpdateOne {
	if id != nil {
		huo = huo.SetEdgesOfSymptomseverityID(*id)
	}
	return huo
}

// SetEdgesOfSymptomseverity sets the EdgesOfSymptomseverity edge to Symptomseverity.
func (huo *HistorytakingUpdateOne) SetEdgesOfSymptomseverity(s *Symptomseverity) *HistorytakingUpdateOne {
	return huo.SetEdgesOfSymptomseverityID(s.ID)
}

// SetEdgesOfPatientrecordID sets the EdgesOfPatientrecord edge to Patientrecord by id.
func (huo *HistorytakingUpdateOne) SetEdgesOfPatientrecordID(id int) *HistorytakingUpdateOne {
	huo.mutation.SetEdgesOfPatientrecordID(id)
	return huo
}

// SetNillableEdgesOfPatientrecordID sets the EdgesOfPatientrecord edge to Patientrecord by id if the given value is not nil.
func (huo *HistorytakingUpdateOne) SetNillableEdgesOfPatientrecordID(id *int) *HistorytakingUpdateOne {
	if id != nil {
		huo = huo.SetEdgesOfPatientrecordID(*id)
	}
	return huo
}

// SetEdgesOfPatientrecord sets the EdgesOfPatientrecord edge to Patientrecord.
func (huo *HistorytakingUpdateOne) SetEdgesOfPatientrecord(p *Patientrecord) *HistorytakingUpdateOne {
	return huo.SetEdgesOfPatientrecordID(p.ID)
}

// Mutation returns the HistorytakingMutation object of the builder.
func (huo *HistorytakingUpdateOne) Mutation() *HistorytakingMutation {
	return huo.mutation
}

// ClearEdgesOfNurse clears the EdgesOfNurse edge to Nurse.
func (huo *HistorytakingUpdateOne) ClearEdgesOfNurse() *HistorytakingUpdateOne {
	huo.mutation.ClearEdgesOfNurse()
	return huo
}

// ClearEdgesOfDepartment clears the EdgesOfDepartment edge to Department.
func (huo *HistorytakingUpdateOne) ClearEdgesOfDepartment() *HistorytakingUpdateOne {
	huo.mutation.ClearEdgesOfDepartment()
	return huo
}

// ClearEdgesOfSymptomseverity clears the EdgesOfSymptomseverity edge to Symptomseverity.
func (huo *HistorytakingUpdateOne) ClearEdgesOfSymptomseverity() *HistorytakingUpdateOne {
	huo.mutation.ClearEdgesOfSymptomseverity()
	return huo
}

// ClearEdgesOfPatientrecord clears the EdgesOfPatientrecord edge to Patientrecord.
func (huo *HistorytakingUpdateOne) ClearEdgesOfPatientrecord() *HistorytakingUpdateOne {
	huo.mutation.ClearEdgesOfPatientrecord()
	return huo
}

// Save executes the query and returns the updated entity.
func (huo *HistorytakingUpdateOne) Save(ctx context.Context) (*Historytaking, error) {

	var (
		err  error
		node *Historytaking
	)
	if len(huo.hooks) == 0 {
		node, err = huo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HistorytakingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			huo.mutation = mutation
			node, err = huo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(huo.hooks) - 1; i >= 0; i-- {
			mut = huo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, huo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HistorytakingUpdateOne) SaveX(ctx context.Context) *Historytaking {
	h, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return h
}

// Exec executes the query on the entity.
func (huo *HistorytakingUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HistorytakingUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (huo *HistorytakingUpdateOne) sqlSave(ctx context.Context) (h *Historytaking, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   historytaking.Table,
			Columns: historytaking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: historytaking.FieldID,
			},
		},
	}
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Historytaking.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := huo.mutation.Hight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldHight,
		})
	}
	if value, ok := huo.mutation.AddedHight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldHight,
		})
	}
	if value, ok := huo.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldWeight,
		})
	}
	if value, ok := huo.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldWeight,
		})
	}
	if value, ok := huo.mutation.Temp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldTemp,
		})
	}
	if value, ok := huo.mutation.AddedTemp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldTemp,
		})
	}
	if value, ok := huo.mutation.Pulse(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldPulse,
		})
	}
	if value, ok := huo.mutation.AddedPulse(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldPulse,
		})
	}
	if value, ok := huo.mutation.Respiration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldRespiration,
		})
	}
	if value, ok := huo.mutation.AddedRespiration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldRespiration,
		})
	}
	if value, ok := huo.mutation.Bp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldBp,
		})
	}
	if value, ok := huo.mutation.AddedBp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldBp,
		})
	}
	if value, ok := huo.mutation.Oxygen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: historytaking.FieldOxygen,
		})
	}
	if value, ok := huo.mutation.Symptom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: historytaking.FieldSymptom,
		})
	}
	if value, ok := huo.mutation.Datetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: historytaking.FieldDatetime,
		})
	}
	if huo.mutation.EdgesOfNurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfNurseTable,
			Columns: []string{historytaking.EdgesOfNurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.EdgesOfNurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfNurseTable,
			Columns: []string{historytaking.EdgesOfNurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.EdgesOfDepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfDepartmentTable,
			Columns: []string{historytaking.EdgesOfDepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.EdgesOfDepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfDepartmentTable,
			Columns: []string{historytaking.EdgesOfDepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.EdgesOfSymptomseverityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfSymptomseverityTable,
			Columns: []string{historytaking.EdgesOfSymptomseverityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: symptomseverity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.EdgesOfSymptomseverityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfSymptomseverityTable,
			Columns: []string{historytaking.EdgesOfSymptomseverityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: symptomseverity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.EdgesOfPatientrecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfPatientrecordTable,
			Columns: []string{historytaking.EdgesOfPatientrecordColumn},
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
	if nodes := huo.mutation.EdgesOfPatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.EdgesOfPatientrecordTable,
			Columns: []string{historytaking.EdgesOfPatientrecordColumn},
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
	h = &Historytaking{config: huo.config}
	_spec.Assign = h.assignValues
	_spec.ScanValues = h.scanValues()
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{historytaking.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return h, nil
}
