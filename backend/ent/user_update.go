// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/doctor"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/patientrights"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/registrar"
	"github.com/team10/app/ent/user"
	"github.com/team10/app/ent/userstatus"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks      []Hook
	mutation   *UserMutation
	predicates []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetEmail sets the email field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPassword sets the password field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetImages sets the images field.
func (uu *UserUpdate) SetImages(s string) *UserUpdate {
	uu.mutation.SetImages(s)
	return uu
}

// SetEdgesOfFinancierID sets the EdgesOfFinancier edge to Financier by id.
func (uu *UserUpdate) SetEdgesOfFinancierID(id int) *UserUpdate {
	uu.mutation.SetEdgesOfFinancierID(id)
	return uu
}

// SetNillableEdgesOfFinancierID sets the EdgesOfFinancier edge to Financier by id if the given value is not nil.
func (uu *UserUpdate) SetNillableEdgesOfFinancierID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetEdgesOfFinancierID(*id)
	}
	return uu
}

// SetEdgesOfFinancier sets the EdgesOfFinancier edge to Financier.
func (uu *UserUpdate) SetEdgesOfFinancier(f *Financier) *UserUpdate {
	return uu.SetEdgesOfFinancierID(f.ID)
}

// SetEdgesOfNurseID sets the EdgesOfNurse edge to Nurse by id.
func (uu *UserUpdate) SetEdgesOfNurseID(id int) *UserUpdate {
	uu.mutation.SetEdgesOfNurseID(id)
	return uu
}

// SetNillableEdgesOfNurseID sets the EdgesOfNurse edge to Nurse by id if the given value is not nil.
func (uu *UserUpdate) SetNillableEdgesOfNurseID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetEdgesOfNurseID(*id)
	}
	return uu
}

// SetEdgesOfNurse sets the EdgesOfNurse edge to Nurse.
func (uu *UserUpdate) SetEdgesOfNurse(n *Nurse) *UserUpdate {
	return uu.SetEdgesOfNurseID(n.ID)
}

// SetEdgesOfUserPatientrightsID sets the EdgesOfUserPatientrights edge to Patientrights by id.
func (uu *UserUpdate) SetEdgesOfUserPatientrightsID(id int) *UserUpdate {
	uu.mutation.SetEdgesOfUserPatientrightsID(id)
	return uu
}

// SetNillableEdgesOfUserPatientrightsID sets the EdgesOfUserPatientrights edge to Patientrights by id if the given value is not nil.
func (uu *UserUpdate) SetNillableEdgesOfUserPatientrightsID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetEdgesOfUserPatientrightsID(*id)
	}
	return uu
}

// SetEdgesOfUserPatientrights sets the EdgesOfUserPatientrights edge to Patientrights.
func (uu *UserUpdate) SetEdgesOfUserPatientrights(p *Patientrights) *UserUpdate {
	return uu.SetEdgesOfUserPatientrightsID(p.ID)
}

// SetEdgesOfMedicalrecordstaffID sets the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff by id.
func (uu *UserUpdate) SetEdgesOfMedicalrecordstaffID(id int) *UserUpdate {
	uu.mutation.SetEdgesOfMedicalrecordstaffID(id)
	return uu
}

// SetNillableEdgesOfMedicalrecordstaffID sets the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff by id if the given value is not nil.
func (uu *UserUpdate) SetNillableEdgesOfMedicalrecordstaffID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetEdgesOfMedicalrecordstaffID(*id)
	}
	return uu
}

// SetEdgesOfMedicalrecordstaff sets the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff.
func (uu *UserUpdate) SetEdgesOfMedicalrecordstaff(m *Medicalrecordstaff) *UserUpdate {
	return uu.SetEdgesOfMedicalrecordstaffID(m.ID)
}

// SetEdgesOfUser2registrarID sets the EdgesOfUser2registrar edge to Registrar by id.
func (uu *UserUpdate) SetEdgesOfUser2registrarID(id int) *UserUpdate {
	uu.mutation.SetEdgesOfUser2registrarID(id)
	return uu
}

// SetNillableEdgesOfUser2registrarID sets the EdgesOfUser2registrar edge to Registrar by id if the given value is not nil.
func (uu *UserUpdate) SetNillableEdgesOfUser2registrarID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetEdgesOfUser2registrarID(*id)
	}
	return uu
}

// SetEdgesOfUser2registrar sets the EdgesOfUser2registrar edge to Registrar.
func (uu *UserUpdate) SetEdgesOfUser2registrar(r *Registrar) *UserUpdate {
	return uu.SetEdgesOfUser2registrarID(r.ID)
}

// SetEdgesOfDoctorID sets the EdgesOfDoctor edge to Doctor by id.
func (uu *UserUpdate) SetEdgesOfDoctorID(id int) *UserUpdate {
	uu.mutation.SetEdgesOfDoctorID(id)
	return uu
}

// SetNillableEdgesOfDoctorID sets the EdgesOfDoctor edge to Doctor by id if the given value is not nil.
func (uu *UserUpdate) SetNillableEdgesOfDoctorID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetEdgesOfDoctorID(*id)
	}
	return uu
}

// SetEdgesOfDoctor sets the EdgesOfDoctor edge to Doctor.
func (uu *UserUpdate) SetEdgesOfDoctor(d *Doctor) *UserUpdate {
	return uu.SetEdgesOfDoctorID(d.ID)
}

// SetEdgesOfUserstatusID sets the EdgesOfUserstatus edge to Userstatus by id.
func (uu *UserUpdate) SetEdgesOfUserstatusID(id int) *UserUpdate {
	uu.mutation.SetEdgesOfUserstatusID(id)
	return uu
}

// SetNillableEdgesOfUserstatusID sets the EdgesOfUserstatus edge to Userstatus by id if the given value is not nil.
func (uu *UserUpdate) SetNillableEdgesOfUserstatusID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetEdgesOfUserstatusID(*id)
	}
	return uu
}

// SetEdgesOfUserstatus sets the EdgesOfUserstatus edge to Userstatus.
func (uu *UserUpdate) SetEdgesOfUserstatus(u *Userstatus) *UserUpdate {
	return uu.SetEdgesOfUserstatusID(u.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearEdgesOfFinancier clears the EdgesOfFinancier edge to Financier.
func (uu *UserUpdate) ClearEdgesOfFinancier() *UserUpdate {
	uu.mutation.ClearEdgesOfFinancier()
	return uu
}

// ClearEdgesOfNurse clears the EdgesOfNurse edge to Nurse.
func (uu *UserUpdate) ClearEdgesOfNurse() *UserUpdate {
	uu.mutation.ClearEdgesOfNurse()
	return uu
}

// ClearEdgesOfUserPatientrights clears the EdgesOfUserPatientrights edge to Patientrights.
func (uu *UserUpdate) ClearEdgesOfUserPatientrights() *UserUpdate {
	uu.mutation.ClearEdgesOfUserPatientrights()
	return uu
}

// ClearEdgesOfMedicalrecordstaff clears the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff.
func (uu *UserUpdate) ClearEdgesOfMedicalrecordstaff() *UserUpdate {
	uu.mutation.ClearEdgesOfMedicalrecordstaff()
	return uu
}

// ClearEdgesOfUser2registrar clears the EdgesOfUser2registrar edge to Registrar.
func (uu *UserUpdate) ClearEdgesOfUser2registrar() *UserUpdate {
	uu.mutation.ClearEdgesOfUser2registrar()
	return uu
}

// ClearEdgesOfDoctor clears the EdgesOfDoctor edge to Doctor.
func (uu *UserUpdate) ClearEdgesOfDoctor() *UserUpdate {
	uu.mutation.ClearEdgesOfDoctor()
	return uu
}

// ClearEdgesOfUserstatus clears the EdgesOfUserstatus edge to Userstatus.
func (uu *UserUpdate) ClearEdgesOfUserstatus() *UserUpdate {
	uu.mutation.ClearEdgesOfUserstatus()
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return 0, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uu.mutation.Images(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldImages,
		})
	}
	if uu.mutation.EdgesOfFinancierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfFinancierTable,
			Columns: []string{user.EdgesOfFinancierColumn},
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
	if nodes := uu.mutation.EdgesOfFinancierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfFinancierTable,
			Columns: []string{user.EdgesOfFinancierColumn},
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
	if uu.mutation.EdgesOfNurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfNurseTable,
			Columns: []string{user.EdgesOfNurseColumn},
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
	if nodes := uu.mutation.EdgesOfNurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfNurseTable,
			Columns: []string{user.EdgesOfNurseColumn},
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
	if uu.mutation.EdgesOfUserPatientrightsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.EdgesOfUserPatientrightsTable,
			Columns: []string{user.EdgesOfUserPatientrightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrights.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.EdgesOfUserPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.EdgesOfUserPatientrightsTable,
			Columns: []string{user.EdgesOfUserPatientrightsColumn},
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
	if uu.mutation.EdgesOfMedicalrecordstaffCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfMedicalrecordstaffTable,
			Columns: []string{user.EdgesOfMedicalrecordstaffColumn},
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
	if nodes := uu.mutation.EdgesOfMedicalrecordstaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfMedicalrecordstaffTable,
			Columns: []string{user.EdgesOfMedicalrecordstaffColumn},
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
	if uu.mutation.EdgesOfUser2registrarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfUser2registrarTable,
			Columns: []string{user.EdgesOfUser2registrarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: registrar.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.EdgesOfUser2registrarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfUser2registrarTable,
			Columns: []string{user.EdgesOfUser2registrarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: registrar.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.EdgesOfDoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfDoctorTable,
			Columns: []string{user.EdgesOfDoctorColumn},
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
	if nodes := uu.mutation.EdgesOfDoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfDoctorTable,
			Columns: []string{user.EdgesOfDoctorColumn},
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
	if uu.mutation.EdgesOfUserstatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.EdgesOfUserstatusTable,
			Columns: []string{user.EdgesOfUserstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.EdgesOfUserstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.EdgesOfUserstatusTable,
			Columns: []string{user.EdgesOfUserstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the email field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPassword sets the password field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetImages sets the images field.
func (uuo *UserUpdateOne) SetImages(s string) *UserUpdateOne {
	uuo.mutation.SetImages(s)
	return uuo
}

// SetEdgesOfFinancierID sets the EdgesOfFinancier edge to Financier by id.
func (uuo *UserUpdateOne) SetEdgesOfFinancierID(id int) *UserUpdateOne {
	uuo.mutation.SetEdgesOfFinancierID(id)
	return uuo
}

// SetNillableEdgesOfFinancierID sets the EdgesOfFinancier edge to Financier by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEdgesOfFinancierID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetEdgesOfFinancierID(*id)
	}
	return uuo
}

// SetEdgesOfFinancier sets the EdgesOfFinancier edge to Financier.
func (uuo *UserUpdateOne) SetEdgesOfFinancier(f *Financier) *UserUpdateOne {
	return uuo.SetEdgesOfFinancierID(f.ID)
}

// SetEdgesOfNurseID sets the EdgesOfNurse edge to Nurse by id.
func (uuo *UserUpdateOne) SetEdgesOfNurseID(id int) *UserUpdateOne {
	uuo.mutation.SetEdgesOfNurseID(id)
	return uuo
}

// SetNillableEdgesOfNurseID sets the EdgesOfNurse edge to Nurse by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEdgesOfNurseID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetEdgesOfNurseID(*id)
	}
	return uuo
}

// SetEdgesOfNurse sets the EdgesOfNurse edge to Nurse.
func (uuo *UserUpdateOne) SetEdgesOfNurse(n *Nurse) *UserUpdateOne {
	return uuo.SetEdgesOfNurseID(n.ID)
}

// SetEdgesOfUserPatientrightsID sets the EdgesOfUserPatientrights edge to Patientrights by id.
func (uuo *UserUpdateOne) SetEdgesOfUserPatientrightsID(id int) *UserUpdateOne {
	uuo.mutation.SetEdgesOfUserPatientrightsID(id)
	return uuo
}

// SetNillableEdgesOfUserPatientrightsID sets the EdgesOfUserPatientrights edge to Patientrights by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEdgesOfUserPatientrightsID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetEdgesOfUserPatientrightsID(*id)
	}
	return uuo
}

// SetEdgesOfUserPatientrights sets the EdgesOfUserPatientrights edge to Patientrights.
func (uuo *UserUpdateOne) SetEdgesOfUserPatientrights(p *Patientrights) *UserUpdateOne {
	return uuo.SetEdgesOfUserPatientrightsID(p.ID)
}

// SetEdgesOfMedicalrecordstaffID sets the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff by id.
func (uuo *UserUpdateOne) SetEdgesOfMedicalrecordstaffID(id int) *UserUpdateOne {
	uuo.mutation.SetEdgesOfMedicalrecordstaffID(id)
	return uuo
}

// SetNillableEdgesOfMedicalrecordstaffID sets the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEdgesOfMedicalrecordstaffID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetEdgesOfMedicalrecordstaffID(*id)
	}
	return uuo
}

// SetEdgesOfMedicalrecordstaff sets the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff.
func (uuo *UserUpdateOne) SetEdgesOfMedicalrecordstaff(m *Medicalrecordstaff) *UserUpdateOne {
	return uuo.SetEdgesOfMedicalrecordstaffID(m.ID)
}

// SetEdgesOfUser2registrarID sets the EdgesOfUser2registrar edge to Registrar by id.
func (uuo *UserUpdateOne) SetEdgesOfUser2registrarID(id int) *UserUpdateOne {
	uuo.mutation.SetEdgesOfUser2registrarID(id)
	return uuo
}

// SetNillableEdgesOfUser2registrarID sets the EdgesOfUser2registrar edge to Registrar by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEdgesOfUser2registrarID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetEdgesOfUser2registrarID(*id)
	}
	return uuo
}

// SetEdgesOfUser2registrar sets the EdgesOfUser2registrar edge to Registrar.
func (uuo *UserUpdateOne) SetEdgesOfUser2registrar(r *Registrar) *UserUpdateOne {
	return uuo.SetEdgesOfUser2registrarID(r.ID)
}

// SetEdgesOfDoctorID sets the EdgesOfDoctor edge to Doctor by id.
func (uuo *UserUpdateOne) SetEdgesOfDoctorID(id int) *UserUpdateOne {
	uuo.mutation.SetEdgesOfDoctorID(id)
	return uuo
}

// SetNillableEdgesOfDoctorID sets the EdgesOfDoctor edge to Doctor by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEdgesOfDoctorID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetEdgesOfDoctorID(*id)
	}
	return uuo
}

// SetEdgesOfDoctor sets the EdgesOfDoctor edge to Doctor.
func (uuo *UserUpdateOne) SetEdgesOfDoctor(d *Doctor) *UserUpdateOne {
	return uuo.SetEdgesOfDoctorID(d.ID)
}

// SetEdgesOfUserstatusID sets the EdgesOfUserstatus edge to Userstatus by id.
func (uuo *UserUpdateOne) SetEdgesOfUserstatusID(id int) *UserUpdateOne {
	uuo.mutation.SetEdgesOfUserstatusID(id)
	return uuo
}

// SetNillableEdgesOfUserstatusID sets the EdgesOfUserstatus edge to Userstatus by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEdgesOfUserstatusID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetEdgesOfUserstatusID(*id)
	}
	return uuo
}

// SetEdgesOfUserstatus sets the EdgesOfUserstatus edge to Userstatus.
func (uuo *UserUpdateOne) SetEdgesOfUserstatus(u *Userstatus) *UserUpdateOne {
	return uuo.SetEdgesOfUserstatusID(u.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearEdgesOfFinancier clears the EdgesOfFinancier edge to Financier.
func (uuo *UserUpdateOne) ClearEdgesOfFinancier() *UserUpdateOne {
	uuo.mutation.ClearEdgesOfFinancier()
	return uuo
}

// ClearEdgesOfNurse clears the EdgesOfNurse edge to Nurse.
func (uuo *UserUpdateOne) ClearEdgesOfNurse() *UserUpdateOne {
	uuo.mutation.ClearEdgesOfNurse()
	return uuo
}

// ClearEdgesOfUserPatientrights clears the EdgesOfUserPatientrights edge to Patientrights.
func (uuo *UserUpdateOne) ClearEdgesOfUserPatientrights() *UserUpdateOne {
	uuo.mutation.ClearEdgesOfUserPatientrights()
	return uuo
}

// ClearEdgesOfMedicalrecordstaff clears the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff.
func (uuo *UserUpdateOne) ClearEdgesOfMedicalrecordstaff() *UserUpdateOne {
	uuo.mutation.ClearEdgesOfMedicalrecordstaff()
	return uuo
}

// ClearEdgesOfUser2registrar clears the EdgesOfUser2registrar edge to Registrar.
func (uuo *UserUpdateOne) ClearEdgesOfUser2registrar() *UserUpdateOne {
	uuo.mutation.ClearEdgesOfUser2registrar()
	return uuo
}

// ClearEdgesOfDoctor clears the EdgesOfDoctor edge to Doctor.
func (uuo *UserUpdateOne) ClearEdgesOfDoctor() *UserUpdateOne {
	uuo.mutation.ClearEdgesOfDoctor()
	return uuo
}

// ClearEdgesOfUserstatus clears the EdgesOfUserstatus edge to Userstatus.
func (uuo *UserUpdateOne) ClearEdgesOfUserstatus() *UserUpdateOne {
	uuo.mutation.ClearEdgesOfUserstatus()
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}

	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uuo.mutation.Images(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldImages,
		})
	}
	if uuo.mutation.EdgesOfFinancierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfFinancierTable,
			Columns: []string{user.EdgesOfFinancierColumn},
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
	if nodes := uuo.mutation.EdgesOfFinancierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfFinancierTable,
			Columns: []string{user.EdgesOfFinancierColumn},
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
	if uuo.mutation.EdgesOfNurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfNurseTable,
			Columns: []string{user.EdgesOfNurseColumn},
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
	if nodes := uuo.mutation.EdgesOfNurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfNurseTable,
			Columns: []string{user.EdgesOfNurseColumn},
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
	if uuo.mutation.EdgesOfUserPatientrightsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.EdgesOfUserPatientrightsTable,
			Columns: []string{user.EdgesOfUserPatientrightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrights.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.EdgesOfUserPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.EdgesOfUserPatientrightsTable,
			Columns: []string{user.EdgesOfUserPatientrightsColumn},
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
	if uuo.mutation.EdgesOfMedicalrecordstaffCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfMedicalrecordstaffTable,
			Columns: []string{user.EdgesOfMedicalrecordstaffColumn},
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
	if nodes := uuo.mutation.EdgesOfMedicalrecordstaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfMedicalrecordstaffTable,
			Columns: []string{user.EdgesOfMedicalrecordstaffColumn},
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
	if uuo.mutation.EdgesOfUser2registrarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfUser2registrarTable,
			Columns: []string{user.EdgesOfUser2registrarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: registrar.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.EdgesOfUser2registrarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfUser2registrarTable,
			Columns: []string{user.EdgesOfUser2registrarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: registrar.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.EdgesOfDoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfDoctorTable,
			Columns: []string{user.EdgesOfDoctorColumn},
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
	if nodes := uuo.mutation.EdgesOfDoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EdgesOfDoctorTable,
			Columns: []string{user.EdgesOfDoctorColumn},
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
	if uuo.mutation.EdgesOfUserstatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.EdgesOfUserstatusTable,
			Columns: []string{user.EdgesOfUserstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.EdgesOfUserstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.EdgesOfUserstatusTable,
			Columns: []string{user.EdgesOfUserstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
