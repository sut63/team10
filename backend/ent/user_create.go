// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/doctor"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/patientrights"
	"github.com/team10/app/ent/registrar"
	"github.com/team10/app/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetPassword sets the password field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetImages sets the images field.
func (uc *UserCreate) SetImages(s string) *UserCreate {
	uc.mutation.SetImages(s)
	return uc
}

// SetEdgesOfFinancierID sets the EdgesOfFinancier edge to Financier by id.
func (uc *UserCreate) SetEdgesOfFinancierID(id int) *UserCreate {
	uc.mutation.SetEdgesOfFinancierID(id)
	return uc
}

// SetNillableEdgesOfFinancierID sets the EdgesOfFinancier edge to Financier by id if the given value is not nil.
func (uc *UserCreate) SetNillableEdgesOfFinancierID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetEdgesOfFinancierID(*id)
	}
	return uc
}

// SetEdgesOfFinancier sets the EdgesOfFinancier edge to Financier.
func (uc *UserCreate) SetEdgesOfFinancier(f *Financier) *UserCreate {
	return uc.SetEdgesOfFinancierID(f.ID)
}

// SetEdgesOfNurseID sets the EdgesOfNurse edge to Nurse by id.
func (uc *UserCreate) SetEdgesOfNurseID(id int) *UserCreate {
	uc.mutation.SetEdgesOfNurseID(id)
	return uc
}

// SetNillableEdgesOfNurseID sets the EdgesOfNurse edge to Nurse by id if the given value is not nil.
func (uc *UserCreate) SetNillableEdgesOfNurseID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetEdgesOfNurseID(*id)
	}
	return uc
}

// SetEdgesOfNurse sets the EdgesOfNurse edge to Nurse.
func (uc *UserCreate) SetEdgesOfNurse(n *Nurse) *UserCreate {
	return uc.SetEdgesOfNurseID(n.ID)
}

// SetEdgesOfUserPatientrightsID sets the EdgesOfUserPatientrights edge to Patientrights by id.
func (uc *UserCreate) SetEdgesOfUserPatientrightsID(id int) *UserCreate {
	uc.mutation.SetEdgesOfUserPatientrightsID(id)
	return uc
}

// SetNillableEdgesOfUserPatientrightsID sets the EdgesOfUserPatientrights edge to Patientrights by id if the given value is not nil.
func (uc *UserCreate) SetNillableEdgesOfUserPatientrightsID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetEdgesOfUserPatientrightsID(*id)
	}
	return uc
}

// SetEdgesOfUserPatientrights sets the EdgesOfUserPatientrights edge to Patientrights.
func (uc *UserCreate) SetEdgesOfUserPatientrights(p *Patientrights) *UserCreate {
	return uc.SetEdgesOfUserPatientrightsID(p.ID)
}

// SetEdgesOfMedicalrecordstaffID sets the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff by id.
func (uc *UserCreate) SetEdgesOfMedicalrecordstaffID(id int) *UserCreate {
	uc.mutation.SetEdgesOfMedicalrecordstaffID(id)
	return uc
}

// SetNillableEdgesOfMedicalrecordstaffID sets the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff by id if the given value is not nil.
func (uc *UserCreate) SetNillableEdgesOfMedicalrecordstaffID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetEdgesOfMedicalrecordstaffID(*id)
	}
	return uc
}

// SetEdgesOfMedicalrecordstaff sets the EdgesOfMedicalrecordstaff edge to Medicalrecordstaff.
func (uc *UserCreate) SetEdgesOfMedicalrecordstaff(m *Medicalrecordstaff) *UserCreate {
	return uc.SetEdgesOfMedicalrecordstaffID(m.ID)
}

// SetEdgesOfUser2registrarID sets the EdgesOfUser2registrar edge to Registrar by id.
func (uc *UserCreate) SetEdgesOfUser2registrarID(id int) *UserCreate {
	uc.mutation.SetEdgesOfUser2registrarID(id)
	return uc
}

// SetNillableEdgesOfUser2registrarID sets the EdgesOfUser2registrar edge to Registrar by id if the given value is not nil.
func (uc *UserCreate) SetNillableEdgesOfUser2registrarID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetEdgesOfUser2registrarID(*id)
	}
	return uc
}

// SetEdgesOfUser2registrar sets the EdgesOfUser2registrar edge to Registrar.
func (uc *UserCreate) SetEdgesOfUser2registrar(r *Registrar) *UserCreate {
	return uc.SetEdgesOfUser2registrarID(r.ID)
}

// SetEdgesOfDoctorID sets the EdgesOfDoctor edge to Doctor by id.
func (uc *UserCreate) SetEdgesOfDoctorID(id int) *UserCreate {
	uc.mutation.SetEdgesOfDoctorID(id)
	return uc
}

// SetNillableEdgesOfDoctorID sets the EdgesOfDoctor edge to Doctor by id if the given value is not nil.
func (uc *UserCreate) SetNillableEdgesOfDoctorID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetEdgesOfDoctorID(*id)
	}
	return uc
}

// SetEdgesOfDoctor sets the EdgesOfDoctor edge to Doctor.
func (uc *UserCreate) SetEdgesOfDoctor(d *Doctor) *UserCreate {
	return uc.SetEdgesOfDoctorID(d.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if _, ok := uc.mutation.Email(); !ok {
		return nil, &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return nil, &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Images(); !ok {
		return nil, &ValidationError{Name: "images", err: errors.New("ent: missing required field \"images\"")}
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	u, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		u.Email = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
		u.Password = value
	}
	if value, ok := uc.mutation.Images(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldImages,
		})
		u.Images = value
	}
	if nodes := uc.mutation.EdgesOfFinancierIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.EdgesOfNurseIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.EdgesOfUserPatientrightsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.EdgesOfMedicalrecordstaffIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.EdgesOfUser2registrarIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.EdgesOfDoctorIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return u, _spec
}
