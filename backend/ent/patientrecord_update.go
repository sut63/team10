// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/gender"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/patientrights"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/prename"
	"github.com/team10/app/ent/treatment"
)

// PatientrecordUpdate is the builder for updating Patientrecord entities.
type PatientrecordUpdate struct {
	config
	hooks      []Hook
	mutation   *PatientrecordMutation
	predicates []predicate.Patientrecord
}

// Where adds a new predicate for the builder.
func (pu *PatientrecordUpdate) Where(ps ...predicate.Patientrecord) *PatientrecordUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetName sets the Name field.
func (pu *PatientrecordUpdate) SetName(s string) *PatientrecordUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetIdcardnumber sets the Idcardnumber field.
func (pu *PatientrecordUpdate) SetIdcardnumber(i int) *PatientrecordUpdate {
	pu.mutation.ResetIdcardnumber()
	pu.mutation.SetIdcardnumber(i)
	return pu
}

// AddIdcardnumber adds i to Idcardnumber.
func (pu *PatientrecordUpdate) AddIdcardnumber(i int) *PatientrecordUpdate {
	pu.mutation.AddIdcardnumber(i)
	return pu
}

// SetAge sets the Age field.
func (pu *PatientrecordUpdate) SetAge(i int) *PatientrecordUpdate {
	pu.mutation.ResetAge()
	pu.mutation.SetAge(i)
	return pu
}

// AddAge adds i to Age.
func (pu *PatientrecordUpdate) AddAge(i int) *PatientrecordUpdate {
	pu.mutation.AddAge(i)
	return pu
}

// SetBirthday sets the Birthday field.
func (pu *PatientrecordUpdate) SetBirthday(t time.Time) *PatientrecordUpdate {
	pu.mutation.SetBirthday(t)
	return pu
}

// SetBloodtype sets the Bloodtype field.
func (pu *PatientrecordUpdate) SetBloodtype(s string) *PatientrecordUpdate {
	pu.mutation.SetBloodtype(s)
	return pu
}

// SetDisease sets the Disease field.
func (pu *PatientrecordUpdate) SetDisease(s string) *PatientrecordUpdate {
	pu.mutation.SetDisease(s)
	return pu
}

// SetAllergic sets the Allergic field.
func (pu *PatientrecordUpdate) SetAllergic(s string) *PatientrecordUpdate {
	pu.mutation.SetAllergic(s)
	return pu
}

// SetPhonenumber sets the Phonenumber field.
func (pu *PatientrecordUpdate) SetPhonenumber(s string) *PatientrecordUpdate {
	pu.mutation.SetPhonenumber(s)
	return pu
}

// SetEmail sets the Email field.
func (pu *PatientrecordUpdate) SetEmail(s string) *PatientrecordUpdate {
	pu.mutation.SetEmail(s)
	return pu
}

// SetHome sets the Home field.
func (pu *PatientrecordUpdate) SetHome(s string) *PatientrecordUpdate {
	pu.mutation.SetHome(s)
	return pu
}

// SetDate sets the Date field.
func (pu *PatientrecordUpdate) SetDate(t time.Time) *PatientrecordUpdate {
	pu.mutation.SetDate(t)
	return pu
}

// SetGenderID sets the gender edge to Gender by id.
func (pu *PatientrecordUpdate) SetGenderID(id int) *PatientrecordUpdate {
	pu.mutation.SetGenderID(id)
	return pu
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (pu *PatientrecordUpdate) SetNillableGenderID(id *int) *PatientrecordUpdate {
	if id != nil {
		pu = pu.SetGenderID(*id)
	}
	return pu
}

// SetGender sets the gender edge to Gender.
func (pu *PatientrecordUpdate) SetGender(g *Gender) *PatientrecordUpdate {
	return pu.SetGenderID(g.ID)
}

// SetMedicalrecordstaffID sets the medicalrecordstaff edge to Medicalrecordstaff by id.
func (pu *PatientrecordUpdate) SetMedicalrecordstaffID(id int) *PatientrecordUpdate {
	pu.mutation.SetMedicalrecordstaffID(id)
	return pu
}

// SetNillableMedicalrecordstaffID sets the medicalrecordstaff edge to Medicalrecordstaff by id if the given value is not nil.
func (pu *PatientrecordUpdate) SetNillableMedicalrecordstaffID(id *int) *PatientrecordUpdate {
	if id != nil {
		pu = pu.SetMedicalrecordstaffID(*id)
	}
	return pu
}

// SetMedicalrecordstaff sets the medicalrecordstaff edge to Medicalrecordstaff.
func (pu *PatientrecordUpdate) SetMedicalrecordstaff(m *Medicalrecordstaff) *PatientrecordUpdate {
	return pu.SetMedicalrecordstaffID(m.ID)
}

// SetPrenameID sets the prename edge to Prename by id.
func (pu *PatientrecordUpdate) SetPrenameID(id int) *PatientrecordUpdate {
	pu.mutation.SetPrenameID(id)
	return pu
}

// SetNillablePrenameID sets the prename edge to Prename by id if the given value is not nil.
func (pu *PatientrecordUpdate) SetNillablePrenameID(id *int) *PatientrecordUpdate {
	if id != nil {
		pu = pu.SetPrenameID(*id)
	}
	return pu
}

// SetPrename sets the prename edge to Prename.
func (pu *PatientrecordUpdate) SetPrename(p *Prename) *PatientrecordUpdate {
	return pu.SetPrenameID(p.ID)
}

// AddHistorytakingIDs adds the historytaking edge to Historytaking by ids.
func (pu *PatientrecordUpdate) AddHistorytakingIDs(ids ...int) *PatientrecordUpdate {
	pu.mutation.AddHistorytakingIDs(ids...)
	return pu
}

// AddHistorytaking adds the historytaking edges to Historytaking.
func (pu *PatientrecordUpdate) AddHistorytaking(h ...*Historytaking) *PatientrecordUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pu.AddHistorytakingIDs(ids...)
}

// AddTreatmentIDs adds the treatment edge to Treatment by ids.
func (pu *PatientrecordUpdate) AddTreatmentIDs(ids ...int) *PatientrecordUpdate {
	pu.mutation.AddTreatmentIDs(ids...)
	return pu
}

// AddTreatment adds the treatment edges to Treatment.
func (pu *PatientrecordUpdate) AddTreatment(t ...*Treatment) *PatientrecordUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTreatmentIDs(ids...)
}

// AddPatientrecordPatientrightIDs adds the PatientrecordPatientrights edge to Patientrights by ids.
func (pu *PatientrecordUpdate) AddPatientrecordPatientrightIDs(ids ...int) *PatientrecordUpdate {
	pu.mutation.AddPatientrecordPatientrightIDs(ids...)
	return pu
}

// AddPatientrecordPatientrights adds the PatientrecordPatientrights edges to Patientrights.
func (pu *PatientrecordUpdate) AddPatientrecordPatientrights(p ...*Patientrights) *PatientrecordUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPatientrecordPatientrightIDs(ids...)
}

// Mutation returns the PatientrecordMutation object of the builder.
func (pu *PatientrecordUpdate) Mutation() *PatientrecordMutation {
	return pu.mutation
}

// ClearGender clears the gender edge to Gender.
func (pu *PatientrecordUpdate) ClearGender() *PatientrecordUpdate {
	pu.mutation.ClearGender()
	return pu
}

// ClearMedicalrecordstaff clears the medicalrecordstaff edge to Medicalrecordstaff.
func (pu *PatientrecordUpdate) ClearMedicalrecordstaff() *PatientrecordUpdate {
	pu.mutation.ClearMedicalrecordstaff()
	return pu
}

// ClearPrename clears the prename edge to Prename.
func (pu *PatientrecordUpdate) ClearPrename() *PatientrecordUpdate {
	pu.mutation.ClearPrename()
	return pu
}

// RemoveHistorytakingIDs removes the historytaking edge to Historytaking by ids.
func (pu *PatientrecordUpdate) RemoveHistorytakingIDs(ids ...int) *PatientrecordUpdate {
	pu.mutation.RemoveHistorytakingIDs(ids...)
	return pu
}

// RemoveHistorytaking removes historytaking edges to Historytaking.
func (pu *PatientrecordUpdate) RemoveHistorytaking(h ...*Historytaking) *PatientrecordUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pu.RemoveHistorytakingIDs(ids...)
}

// RemoveTreatmentIDs removes the treatment edge to Treatment by ids.
func (pu *PatientrecordUpdate) RemoveTreatmentIDs(ids ...int) *PatientrecordUpdate {
	pu.mutation.RemoveTreatmentIDs(ids...)
	return pu
}

// RemoveTreatment removes treatment edges to Treatment.
func (pu *PatientrecordUpdate) RemoveTreatment(t ...*Treatment) *PatientrecordUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTreatmentIDs(ids...)
}

// RemovePatientrecordPatientrightIDs removes the PatientrecordPatientrights edge to Patientrights by ids.
func (pu *PatientrecordUpdate) RemovePatientrecordPatientrightIDs(ids ...int) *PatientrecordUpdate {
	pu.mutation.RemovePatientrecordPatientrightIDs(ids...)
	return pu
}

// RemovePatientrecordPatientrights removes PatientrecordPatientrights edges to Patientrights.
func (pu *PatientrecordUpdate) RemovePatientrecordPatientrights(p ...*Patientrights) *PatientrecordUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePatientrecordPatientrightIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PatientrecordUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientrecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PatientrecordUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientrecordUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientrecordUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PatientrecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patientrecord.Table,
			Columns: patientrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrecord.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldName,
		})
	}
	if value, ok := pu.mutation.Idcardnumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patientrecord.FieldIdcardnumber,
		})
	}
	if value, ok := pu.mutation.AddedIdcardnumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patientrecord.FieldIdcardnumber,
		})
	}
	if value, ok := pu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patientrecord.FieldAge,
		})
	}
	if value, ok := pu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patientrecord.FieldAge,
		})
	}
	if value, ok := pu.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patientrecord.FieldBirthday,
		})
	}
	if value, ok := pu.mutation.Bloodtype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldBloodtype,
		})
	}
	if value, ok := pu.mutation.Disease(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldDisease,
		})
	}
	if value, ok := pu.mutation.Allergic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldAllergic,
		})
	}
	if value, ok := pu.mutation.Phonenumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldPhonenumber,
		})
	}
	if value, ok := pu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldEmail,
		})
	}
	if value, ok := pu.mutation.Home(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldHome,
		})
	}
	if value, ok := pu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patientrecord.FieldDate,
		})
	}
	if pu.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.GenderTable,
			Columns: []string{patientrecord.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.GenderTable,
			Columns: []string{patientrecord.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.MedicalrecordstaffCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.MedicalrecordstaffTable,
			Columns: []string{patientrecord.MedicalrecordstaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalrecordstaff.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MedicalrecordstaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.MedicalrecordstaffTable,
			Columns: []string{patientrecord.MedicalrecordstaffColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PrenameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.PrenameTable,
			Columns: []string{patientrecord.PrenameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prename.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PrenameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.PrenameTable,
			Columns: []string{patientrecord.PrenameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prename.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.HistorytakingTable,
			Columns: []string{patientrecord.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.HistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.HistorytakingTable,
			Columns: []string{patientrecord.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedTreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.TreatmentTable,
			Columns: []string{patientrecord.TreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: treatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.TreatmentTable,
			Columns: []string{patientrecord.TreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: treatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedPatientrecordPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.PatientrecordPatientrightsTable,
			Columns: []string{patientrecord.PatientrecordPatientrightsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PatientrecordPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.PatientrecordPatientrightsTable,
			Columns: []string{patientrecord.PatientrecordPatientrightsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patientrecord.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientrecordUpdateOne is the builder for updating a single Patientrecord entity.
type PatientrecordUpdateOne struct {
	config
	hooks    []Hook
	mutation *PatientrecordMutation
}

// SetName sets the Name field.
func (puo *PatientrecordUpdateOne) SetName(s string) *PatientrecordUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetIdcardnumber sets the Idcardnumber field.
func (puo *PatientrecordUpdateOne) SetIdcardnumber(i int) *PatientrecordUpdateOne {
	puo.mutation.ResetIdcardnumber()
	puo.mutation.SetIdcardnumber(i)
	return puo
}

// AddIdcardnumber adds i to Idcardnumber.
func (puo *PatientrecordUpdateOne) AddIdcardnumber(i int) *PatientrecordUpdateOne {
	puo.mutation.AddIdcardnumber(i)
	return puo
}

// SetAge sets the Age field.
func (puo *PatientrecordUpdateOne) SetAge(i int) *PatientrecordUpdateOne {
	puo.mutation.ResetAge()
	puo.mutation.SetAge(i)
	return puo
}

// AddAge adds i to Age.
func (puo *PatientrecordUpdateOne) AddAge(i int) *PatientrecordUpdateOne {
	puo.mutation.AddAge(i)
	return puo
}

// SetBirthday sets the Birthday field.
func (puo *PatientrecordUpdateOne) SetBirthday(t time.Time) *PatientrecordUpdateOne {
	puo.mutation.SetBirthday(t)
	return puo
}

// SetBloodtype sets the Bloodtype field.
func (puo *PatientrecordUpdateOne) SetBloodtype(s string) *PatientrecordUpdateOne {
	puo.mutation.SetBloodtype(s)
	return puo
}

// SetDisease sets the Disease field.
func (puo *PatientrecordUpdateOne) SetDisease(s string) *PatientrecordUpdateOne {
	puo.mutation.SetDisease(s)
	return puo
}

// SetAllergic sets the Allergic field.
func (puo *PatientrecordUpdateOne) SetAllergic(s string) *PatientrecordUpdateOne {
	puo.mutation.SetAllergic(s)
	return puo
}

// SetPhonenumber sets the Phonenumber field.
func (puo *PatientrecordUpdateOne) SetPhonenumber(s string) *PatientrecordUpdateOne {
	puo.mutation.SetPhonenumber(s)
	return puo
}

// SetEmail sets the Email field.
func (puo *PatientrecordUpdateOne) SetEmail(s string) *PatientrecordUpdateOne {
	puo.mutation.SetEmail(s)
	return puo
}

// SetHome sets the Home field.
func (puo *PatientrecordUpdateOne) SetHome(s string) *PatientrecordUpdateOne {
	puo.mutation.SetHome(s)
	return puo
}

// SetDate sets the Date field.
func (puo *PatientrecordUpdateOne) SetDate(t time.Time) *PatientrecordUpdateOne {
	puo.mutation.SetDate(t)
	return puo
}

// SetGenderID sets the gender edge to Gender by id.
func (puo *PatientrecordUpdateOne) SetGenderID(id int) *PatientrecordUpdateOne {
	puo.mutation.SetGenderID(id)
	return puo
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (puo *PatientrecordUpdateOne) SetNillableGenderID(id *int) *PatientrecordUpdateOne {
	if id != nil {
		puo = puo.SetGenderID(*id)
	}
	return puo
}

// SetGender sets the gender edge to Gender.
func (puo *PatientrecordUpdateOne) SetGender(g *Gender) *PatientrecordUpdateOne {
	return puo.SetGenderID(g.ID)
}

// SetMedicalrecordstaffID sets the medicalrecordstaff edge to Medicalrecordstaff by id.
func (puo *PatientrecordUpdateOne) SetMedicalrecordstaffID(id int) *PatientrecordUpdateOne {
	puo.mutation.SetMedicalrecordstaffID(id)
	return puo
}

// SetNillableMedicalrecordstaffID sets the medicalrecordstaff edge to Medicalrecordstaff by id if the given value is not nil.
func (puo *PatientrecordUpdateOne) SetNillableMedicalrecordstaffID(id *int) *PatientrecordUpdateOne {
	if id != nil {
		puo = puo.SetMedicalrecordstaffID(*id)
	}
	return puo
}

// SetMedicalrecordstaff sets the medicalrecordstaff edge to Medicalrecordstaff.
func (puo *PatientrecordUpdateOne) SetMedicalrecordstaff(m *Medicalrecordstaff) *PatientrecordUpdateOne {
	return puo.SetMedicalrecordstaffID(m.ID)
}

// SetPrenameID sets the prename edge to Prename by id.
func (puo *PatientrecordUpdateOne) SetPrenameID(id int) *PatientrecordUpdateOne {
	puo.mutation.SetPrenameID(id)
	return puo
}

// SetNillablePrenameID sets the prename edge to Prename by id if the given value is not nil.
func (puo *PatientrecordUpdateOne) SetNillablePrenameID(id *int) *PatientrecordUpdateOne {
	if id != nil {
		puo = puo.SetPrenameID(*id)
	}
	return puo
}

// SetPrename sets the prename edge to Prename.
func (puo *PatientrecordUpdateOne) SetPrename(p *Prename) *PatientrecordUpdateOne {
	return puo.SetPrenameID(p.ID)
}

// AddHistorytakingIDs adds the historytaking edge to Historytaking by ids.
func (puo *PatientrecordUpdateOne) AddHistorytakingIDs(ids ...int) *PatientrecordUpdateOne {
	puo.mutation.AddHistorytakingIDs(ids...)
	return puo
}

// AddHistorytaking adds the historytaking edges to Historytaking.
func (puo *PatientrecordUpdateOne) AddHistorytaking(h ...*Historytaking) *PatientrecordUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return puo.AddHistorytakingIDs(ids...)
}

// AddTreatmentIDs adds the treatment edge to Treatment by ids.
func (puo *PatientrecordUpdateOne) AddTreatmentIDs(ids ...int) *PatientrecordUpdateOne {
	puo.mutation.AddTreatmentIDs(ids...)
	return puo
}

// AddTreatment adds the treatment edges to Treatment.
func (puo *PatientrecordUpdateOne) AddTreatment(t ...*Treatment) *PatientrecordUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTreatmentIDs(ids...)
}

// AddPatientrecordPatientrightIDs adds the PatientrecordPatientrights edge to Patientrights by ids.
func (puo *PatientrecordUpdateOne) AddPatientrecordPatientrightIDs(ids ...int) *PatientrecordUpdateOne {
	puo.mutation.AddPatientrecordPatientrightIDs(ids...)
	return puo
}

// AddPatientrecordPatientrights adds the PatientrecordPatientrights edges to Patientrights.
func (puo *PatientrecordUpdateOne) AddPatientrecordPatientrights(p ...*Patientrights) *PatientrecordUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPatientrecordPatientrightIDs(ids...)
}

// Mutation returns the PatientrecordMutation object of the builder.
func (puo *PatientrecordUpdateOne) Mutation() *PatientrecordMutation {
	return puo.mutation
}

// ClearGender clears the gender edge to Gender.
func (puo *PatientrecordUpdateOne) ClearGender() *PatientrecordUpdateOne {
	puo.mutation.ClearGender()
	return puo
}

// ClearMedicalrecordstaff clears the medicalrecordstaff edge to Medicalrecordstaff.
func (puo *PatientrecordUpdateOne) ClearMedicalrecordstaff() *PatientrecordUpdateOne {
	puo.mutation.ClearMedicalrecordstaff()
	return puo
}

// ClearPrename clears the prename edge to Prename.
func (puo *PatientrecordUpdateOne) ClearPrename() *PatientrecordUpdateOne {
	puo.mutation.ClearPrename()
	return puo
}

// RemoveHistorytakingIDs removes the historytaking edge to Historytaking by ids.
func (puo *PatientrecordUpdateOne) RemoveHistorytakingIDs(ids ...int) *PatientrecordUpdateOne {
	puo.mutation.RemoveHistorytakingIDs(ids...)
	return puo
}

// RemoveHistorytaking removes historytaking edges to Historytaking.
func (puo *PatientrecordUpdateOne) RemoveHistorytaking(h ...*Historytaking) *PatientrecordUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return puo.RemoveHistorytakingIDs(ids...)
}

// RemoveTreatmentIDs removes the treatment edge to Treatment by ids.
func (puo *PatientrecordUpdateOne) RemoveTreatmentIDs(ids ...int) *PatientrecordUpdateOne {
	puo.mutation.RemoveTreatmentIDs(ids...)
	return puo
}

// RemoveTreatment removes treatment edges to Treatment.
func (puo *PatientrecordUpdateOne) RemoveTreatment(t ...*Treatment) *PatientrecordUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTreatmentIDs(ids...)
}

// RemovePatientrecordPatientrightIDs removes the PatientrecordPatientrights edge to Patientrights by ids.
func (puo *PatientrecordUpdateOne) RemovePatientrecordPatientrightIDs(ids ...int) *PatientrecordUpdateOne {
	puo.mutation.RemovePatientrecordPatientrightIDs(ids...)
	return puo
}

// RemovePatientrecordPatientrights removes PatientrecordPatientrights edges to Patientrights.
func (puo *PatientrecordUpdateOne) RemovePatientrecordPatientrights(p ...*Patientrights) *PatientrecordUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePatientrecordPatientrightIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PatientrecordUpdateOne) Save(ctx context.Context) (*Patientrecord, error) {

	var (
		err  error
		node *Patientrecord
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientrecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PatientrecordUpdateOne) SaveX(ctx context.Context) *Patientrecord {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PatientrecordUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientrecordUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PatientrecordUpdateOne) sqlSave(ctx context.Context) (pa *Patientrecord, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patientrecord.Table,
			Columns: patientrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrecord.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Patientrecord.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldName,
		})
	}
	if value, ok := puo.mutation.Idcardnumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patientrecord.FieldIdcardnumber,
		})
	}
	if value, ok := puo.mutation.AddedIdcardnumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patientrecord.FieldIdcardnumber,
		})
	}
	if value, ok := puo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patientrecord.FieldAge,
		})
	}
	if value, ok := puo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patientrecord.FieldAge,
		})
	}
	if value, ok := puo.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patientrecord.FieldBirthday,
		})
	}
	if value, ok := puo.mutation.Bloodtype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldBloodtype,
		})
	}
	if value, ok := puo.mutation.Disease(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldDisease,
		})
	}
	if value, ok := puo.mutation.Allergic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldAllergic,
		})
	}
	if value, ok := puo.mutation.Phonenumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldPhonenumber,
		})
	}
	if value, ok := puo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldEmail,
		})
	}
	if value, ok := puo.mutation.Home(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldHome,
		})
	}
	if value, ok := puo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patientrecord.FieldDate,
		})
	}
	if puo.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.GenderTable,
			Columns: []string{patientrecord.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.GenderTable,
			Columns: []string{patientrecord.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.MedicalrecordstaffCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.MedicalrecordstaffTable,
			Columns: []string{patientrecord.MedicalrecordstaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalrecordstaff.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MedicalrecordstaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.MedicalrecordstaffTable,
			Columns: []string{patientrecord.MedicalrecordstaffColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PrenameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.PrenameTable,
			Columns: []string{patientrecord.PrenameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prename.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PrenameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrecord.PrenameTable,
			Columns: []string{patientrecord.PrenameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prename.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.HistorytakingTable,
			Columns: []string{patientrecord.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.HistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.HistorytakingTable,
			Columns: []string{patientrecord.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedTreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.TreatmentTable,
			Columns: []string{patientrecord.TreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: treatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.TreatmentTable,
			Columns: []string{patientrecord.TreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: treatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedPatientrecordPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.PatientrecordPatientrightsTable,
			Columns: []string{patientrecord.PatientrecordPatientrightsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PatientrecordPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.PatientrecordPatientrightsTable,
			Columns: []string{patientrecord.PatientrecordPatientrightsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Patientrecord{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patientrecord.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}
