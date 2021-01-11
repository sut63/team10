// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/doctor"
	"github.com/team10/app/ent/doctorinfo"
	"github.com/team10/app/ent/educationlevel"
	"github.com/team10/app/ent/officeroom"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/prename"
)

// DoctorinfoUpdate is the builder for updating Doctorinfo entities.
type DoctorinfoUpdate struct {
	config
	hooks      []Hook
	mutation   *DoctorinfoMutation
	predicates []predicate.Doctorinfo
}

// Where adds a new predicate for the builder.
func (du *DoctorinfoUpdate) Where(ps ...predicate.Doctorinfo) *DoctorinfoUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDoctorname sets the doctorname field.
func (du *DoctorinfoUpdate) SetDoctorname(s string) *DoctorinfoUpdate {
	du.mutation.SetDoctorname(s)
	return du
}

// SetDoctorsurname sets the doctorsurname field.
func (du *DoctorinfoUpdate) SetDoctorsurname(s string) *DoctorinfoUpdate {
	du.mutation.SetDoctorsurname(s)
	return du
}

// SetTelephonenumber sets the telephonenumber field.
func (du *DoctorinfoUpdate) SetTelephonenumber(s string) *DoctorinfoUpdate {
	du.mutation.SetTelephonenumber(s)
	return du
}

// SetLicensenumber sets the licensenumber field.
func (du *DoctorinfoUpdate) SetLicensenumber(s string) *DoctorinfoUpdate {
	du.mutation.SetLicensenumber(s)
	return du
}

// SetDepartmentID sets the department edge to Department by id.
func (du *DoctorinfoUpdate) SetDepartmentID(id int) *DoctorinfoUpdate {
	du.mutation.SetDepartmentID(id)
	return du
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (du *DoctorinfoUpdate) SetNillableDepartmentID(id *int) *DoctorinfoUpdate {
	if id != nil {
		du = du.SetDepartmentID(*id)
	}
	return du
}

// SetDepartment sets the department edge to Department.
func (du *DoctorinfoUpdate) SetDepartment(d *Department) *DoctorinfoUpdate {
	return du.SetDepartmentID(d.ID)
}

// SetEducationlevelID sets the educationlevel edge to Educationlevel by id.
func (du *DoctorinfoUpdate) SetEducationlevelID(id int) *DoctorinfoUpdate {
	du.mutation.SetEducationlevelID(id)
	return du
}

// SetNillableEducationlevelID sets the educationlevel edge to Educationlevel by id if the given value is not nil.
func (du *DoctorinfoUpdate) SetNillableEducationlevelID(id *int) *DoctorinfoUpdate {
	if id != nil {
		du = du.SetEducationlevelID(*id)
	}
	return du
}

// SetEducationlevel sets the educationlevel edge to Educationlevel.
func (du *DoctorinfoUpdate) SetEducationlevel(e *Educationlevel) *DoctorinfoUpdate {
	return du.SetEducationlevelID(e.ID)
}

// SetOfficeroomID sets the officeroom edge to Officeroom by id.
func (du *DoctorinfoUpdate) SetOfficeroomID(id int) *DoctorinfoUpdate {
	du.mutation.SetOfficeroomID(id)
	return du
}

// SetNillableOfficeroomID sets the officeroom edge to Officeroom by id if the given value is not nil.
func (du *DoctorinfoUpdate) SetNillableOfficeroomID(id *int) *DoctorinfoUpdate {
	if id != nil {
		du = du.SetOfficeroomID(*id)
	}
	return du
}

// SetOfficeroom sets the officeroom edge to Officeroom.
func (du *DoctorinfoUpdate) SetOfficeroom(o *Officeroom) *DoctorinfoUpdate {
	return du.SetOfficeroomID(o.ID)
}

// SetPrenameID sets the prename edge to Prename by id.
func (du *DoctorinfoUpdate) SetPrenameID(id int) *DoctorinfoUpdate {
	du.mutation.SetPrenameID(id)
	return du
}

// SetNillablePrenameID sets the prename edge to Prename by id if the given value is not nil.
func (du *DoctorinfoUpdate) SetNillablePrenameID(id *int) *DoctorinfoUpdate {
	if id != nil {
		du = du.SetPrenameID(*id)
	}
	return du
}

// SetPrename sets the prename edge to Prename.
func (du *DoctorinfoUpdate) SetPrename(p *Prename) *DoctorinfoUpdate {
	return du.SetPrenameID(p.ID)
}

// AddDoctorIDs adds the doctor edge to Doctor by ids.
func (du *DoctorinfoUpdate) AddDoctorIDs(ids ...int) *DoctorinfoUpdate {
	du.mutation.AddDoctorIDs(ids...)
	return du
}

// AddDoctor adds the doctor edges to Doctor.
func (du *DoctorinfoUpdate) AddDoctor(d ...*Doctor) *DoctorinfoUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDoctorIDs(ids...)
}

// Mutation returns the DoctorinfoMutation object of the builder.
func (du *DoctorinfoUpdate) Mutation() *DoctorinfoMutation {
	return du.mutation
}

// ClearDepartment clears the department edge to Department.
func (du *DoctorinfoUpdate) ClearDepartment() *DoctorinfoUpdate {
	du.mutation.ClearDepartment()
	return du
}

// ClearEducationlevel clears the educationlevel edge to Educationlevel.
func (du *DoctorinfoUpdate) ClearEducationlevel() *DoctorinfoUpdate {
	du.mutation.ClearEducationlevel()
	return du
}

// ClearOfficeroom clears the officeroom edge to Officeroom.
func (du *DoctorinfoUpdate) ClearOfficeroom() *DoctorinfoUpdate {
	du.mutation.ClearOfficeroom()
	return du
}

// ClearPrename clears the prename edge to Prename.
func (du *DoctorinfoUpdate) ClearPrename() *DoctorinfoUpdate {
	du.mutation.ClearPrename()
	return du
}

// RemoveDoctorIDs removes the doctor edge to Doctor by ids.
func (du *DoctorinfoUpdate) RemoveDoctorIDs(ids ...int) *DoctorinfoUpdate {
	du.mutation.RemoveDoctorIDs(ids...)
	return du
}

// RemoveDoctor removes doctor edges to Doctor.
func (du *DoctorinfoUpdate) RemoveDoctor(d ...*Doctor) *DoctorinfoUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDoctorIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DoctorinfoUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Doctorname(); ok {
		if err := doctorinfo.DoctornameValidator(v); err != nil {
			return 0, &ValidationError{Name: "doctorname", err: fmt.Errorf("ent: validator failed for field \"doctorname\": %w", err)}
		}
	}
	if v, ok := du.mutation.Doctorsurname(); ok {
		if err := doctorinfo.DoctorsurnameValidator(v); err != nil {
			return 0, &ValidationError{Name: "doctorsurname", err: fmt.Errorf("ent: validator failed for field \"doctorsurname\": %w", err)}
		}
	}
	if v, ok := du.mutation.Telephonenumber(); ok {
		if err := doctorinfo.TelephonenumberValidator(v); err != nil {
			return 0, &ValidationError{Name: "telephonenumber", err: fmt.Errorf("ent: validator failed for field \"telephonenumber\": %w", err)}
		}
	}
	if v, ok := du.mutation.Licensenumber(); ok {
		if err := doctorinfo.LicensenumberValidator(v); err != nil {
			return 0, &ValidationError{Name: "licensenumber", err: fmt.Errorf("ent: validator failed for field \"licensenumber\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorinfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DoctorinfoUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DoctorinfoUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DoctorinfoUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DoctorinfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctorinfo.Table,
			Columns: doctorinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctorinfo.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Doctorname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctorinfo.FieldDoctorname,
		})
	}
	if value, ok := du.mutation.Doctorsurname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctorinfo.FieldDoctorsurname,
		})
	}
	if value, ok := du.mutation.Telephonenumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctorinfo.FieldTelephonenumber,
		})
	}
	if value, ok := du.mutation.Licensenumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctorinfo.FieldLicensenumber,
		})
	}
	if du.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.DepartmentTable,
			Columns: []string{doctorinfo.DepartmentColumn},
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
	if nodes := du.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.DepartmentTable,
			Columns: []string{doctorinfo.DepartmentColumn},
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
	if du.mutation.EducationlevelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.EducationlevelTable,
			Columns: []string{doctorinfo.EducationlevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: educationlevel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.EducationlevelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.EducationlevelTable,
			Columns: []string{doctorinfo.EducationlevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: educationlevel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.OfficeroomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.OfficeroomTable,
			Columns: []string{doctorinfo.OfficeroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: officeroom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.OfficeroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.OfficeroomTable,
			Columns: []string{doctorinfo.OfficeroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: officeroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.PrenameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.PrenameTable,
			Columns: []string{doctorinfo.PrenameColumn},
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
	if nodes := du.mutation.PrenameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.PrenameTable,
			Columns: []string{doctorinfo.PrenameColumn},
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
	if nodes := du.mutation.RemovedDoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctorinfo.DoctorTable,
			Columns: []string{doctorinfo.DoctorColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctorinfo.DoctorTable,
			Columns: []string{doctorinfo.DoctorColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctorinfo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DoctorinfoUpdateOne is the builder for updating a single Doctorinfo entity.
type DoctorinfoUpdateOne struct {
	config
	hooks    []Hook
	mutation *DoctorinfoMutation
}

// SetDoctorname sets the doctorname field.
func (duo *DoctorinfoUpdateOne) SetDoctorname(s string) *DoctorinfoUpdateOne {
	duo.mutation.SetDoctorname(s)
	return duo
}

// SetDoctorsurname sets the doctorsurname field.
func (duo *DoctorinfoUpdateOne) SetDoctorsurname(s string) *DoctorinfoUpdateOne {
	duo.mutation.SetDoctorsurname(s)
	return duo
}

// SetTelephonenumber sets the telephonenumber field.
func (duo *DoctorinfoUpdateOne) SetTelephonenumber(s string) *DoctorinfoUpdateOne {
	duo.mutation.SetTelephonenumber(s)
	return duo
}

// SetLicensenumber sets the licensenumber field.
func (duo *DoctorinfoUpdateOne) SetLicensenumber(s string) *DoctorinfoUpdateOne {
	duo.mutation.SetLicensenumber(s)
	return duo
}

// SetDepartmentID sets the department edge to Department by id.
func (duo *DoctorinfoUpdateOne) SetDepartmentID(id int) *DoctorinfoUpdateOne {
	duo.mutation.SetDepartmentID(id)
	return duo
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (duo *DoctorinfoUpdateOne) SetNillableDepartmentID(id *int) *DoctorinfoUpdateOne {
	if id != nil {
		duo = duo.SetDepartmentID(*id)
	}
	return duo
}

// SetDepartment sets the department edge to Department.
func (duo *DoctorinfoUpdateOne) SetDepartment(d *Department) *DoctorinfoUpdateOne {
	return duo.SetDepartmentID(d.ID)
}

// SetEducationlevelID sets the educationlevel edge to Educationlevel by id.
func (duo *DoctorinfoUpdateOne) SetEducationlevelID(id int) *DoctorinfoUpdateOne {
	duo.mutation.SetEducationlevelID(id)
	return duo
}

// SetNillableEducationlevelID sets the educationlevel edge to Educationlevel by id if the given value is not nil.
func (duo *DoctorinfoUpdateOne) SetNillableEducationlevelID(id *int) *DoctorinfoUpdateOne {
	if id != nil {
		duo = duo.SetEducationlevelID(*id)
	}
	return duo
}

// SetEducationlevel sets the educationlevel edge to Educationlevel.
func (duo *DoctorinfoUpdateOne) SetEducationlevel(e *Educationlevel) *DoctorinfoUpdateOne {
	return duo.SetEducationlevelID(e.ID)
}

// SetOfficeroomID sets the officeroom edge to Officeroom by id.
func (duo *DoctorinfoUpdateOne) SetOfficeroomID(id int) *DoctorinfoUpdateOne {
	duo.mutation.SetOfficeroomID(id)
	return duo
}

// SetNillableOfficeroomID sets the officeroom edge to Officeroom by id if the given value is not nil.
func (duo *DoctorinfoUpdateOne) SetNillableOfficeroomID(id *int) *DoctorinfoUpdateOne {
	if id != nil {
		duo = duo.SetOfficeroomID(*id)
	}
	return duo
}

// SetOfficeroom sets the officeroom edge to Officeroom.
func (duo *DoctorinfoUpdateOne) SetOfficeroom(o *Officeroom) *DoctorinfoUpdateOne {
	return duo.SetOfficeroomID(o.ID)
}

// SetPrenameID sets the prename edge to Prename by id.
func (duo *DoctorinfoUpdateOne) SetPrenameID(id int) *DoctorinfoUpdateOne {
	duo.mutation.SetPrenameID(id)
	return duo
}

// SetNillablePrenameID sets the prename edge to Prename by id if the given value is not nil.
func (duo *DoctorinfoUpdateOne) SetNillablePrenameID(id *int) *DoctorinfoUpdateOne {
	if id != nil {
		duo = duo.SetPrenameID(*id)
	}
	return duo
}

// SetPrename sets the prename edge to Prename.
func (duo *DoctorinfoUpdateOne) SetPrename(p *Prename) *DoctorinfoUpdateOne {
	return duo.SetPrenameID(p.ID)
}

// AddDoctorIDs adds the doctor edge to Doctor by ids.
func (duo *DoctorinfoUpdateOne) AddDoctorIDs(ids ...int) *DoctorinfoUpdateOne {
	duo.mutation.AddDoctorIDs(ids...)
	return duo
}

// AddDoctor adds the doctor edges to Doctor.
func (duo *DoctorinfoUpdateOne) AddDoctor(d ...*Doctor) *DoctorinfoUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDoctorIDs(ids...)
}

// Mutation returns the DoctorinfoMutation object of the builder.
func (duo *DoctorinfoUpdateOne) Mutation() *DoctorinfoMutation {
	return duo.mutation
}

// ClearDepartment clears the department edge to Department.
func (duo *DoctorinfoUpdateOne) ClearDepartment() *DoctorinfoUpdateOne {
	duo.mutation.ClearDepartment()
	return duo
}

// ClearEducationlevel clears the educationlevel edge to Educationlevel.
func (duo *DoctorinfoUpdateOne) ClearEducationlevel() *DoctorinfoUpdateOne {
	duo.mutation.ClearEducationlevel()
	return duo
}

// ClearOfficeroom clears the officeroom edge to Officeroom.
func (duo *DoctorinfoUpdateOne) ClearOfficeroom() *DoctorinfoUpdateOne {
	duo.mutation.ClearOfficeroom()
	return duo
}

// ClearPrename clears the prename edge to Prename.
func (duo *DoctorinfoUpdateOne) ClearPrename() *DoctorinfoUpdateOne {
	duo.mutation.ClearPrename()
	return duo
}

// RemoveDoctorIDs removes the doctor edge to Doctor by ids.
func (duo *DoctorinfoUpdateOne) RemoveDoctorIDs(ids ...int) *DoctorinfoUpdateOne {
	duo.mutation.RemoveDoctorIDs(ids...)
	return duo
}

// RemoveDoctor removes doctor edges to Doctor.
func (duo *DoctorinfoUpdateOne) RemoveDoctor(d ...*Doctor) *DoctorinfoUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDoctorIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DoctorinfoUpdateOne) Save(ctx context.Context) (*Doctorinfo, error) {
	if v, ok := duo.mutation.Doctorname(); ok {
		if err := doctorinfo.DoctornameValidator(v); err != nil {
			return nil, &ValidationError{Name: "doctorname", err: fmt.Errorf("ent: validator failed for field \"doctorname\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Doctorsurname(); ok {
		if err := doctorinfo.DoctorsurnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "doctorsurname", err: fmt.Errorf("ent: validator failed for field \"doctorsurname\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Telephonenumber(); ok {
		if err := doctorinfo.TelephonenumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "telephonenumber", err: fmt.Errorf("ent: validator failed for field \"telephonenumber\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Licensenumber(); ok {
		if err := doctorinfo.LicensenumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "licensenumber", err: fmt.Errorf("ent: validator failed for field \"licensenumber\": %w", err)}
		}
	}

	var (
		err  error
		node *Doctorinfo
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorinfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DoctorinfoUpdateOne) SaveX(ctx context.Context) *Doctorinfo {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DoctorinfoUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DoctorinfoUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DoctorinfoUpdateOne) sqlSave(ctx context.Context) (d *Doctorinfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctorinfo.Table,
			Columns: doctorinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctorinfo.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Doctorinfo.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Doctorname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctorinfo.FieldDoctorname,
		})
	}
	if value, ok := duo.mutation.Doctorsurname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctorinfo.FieldDoctorsurname,
		})
	}
	if value, ok := duo.mutation.Telephonenumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctorinfo.FieldTelephonenumber,
		})
	}
	if value, ok := duo.mutation.Licensenumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctorinfo.FieldLicensenumber,
		})
	}
	if duo.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.DepartmentTable,
			Columns: []string{doctorinfo.DepartmentColumn},
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
	if nodes := duo.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.DepartmentTable,
			Columns: []string{doctorinfo.DepartmentColumn},
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
	if duo.mutation.EducationlevelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.EducationlevelTable,
			Columns: []string{doctorinfo.EducationlevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: educationlevel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.EducationlevelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.EducationlevelTable,
			Columns: []string{doctorinfo.EducationlevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: educationlevel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.OfficeroomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.OfficeroomTable,
			Columns: []string{doctorinfo.OfficeroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: officeroom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.OfficeroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.OfficeroomTable,
			Columns: []string{doctorinfo.OfficeroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: officeroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.PrenameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.PrenameTable,
			Columns: []string{doctorinfo.PrenameColumn},
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
	if nodes := duo.mutation.PrenameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctorinfo.PrenameTable,
			Columns: []string{doctorinfo.PrenameColumn},
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
	if nodes := duo.mutation.RemovedDoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctorinfo.DoctorTable,
			Columns: []string{doctorinfo.DoctorColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctorinfo.DoctorTable,
			Columns: []string{doctorinfo.DoctorColumn},
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
	d = &Doctorinfo{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctorinfo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
