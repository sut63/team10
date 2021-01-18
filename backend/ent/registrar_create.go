// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/registrar"
	"github.com/team10/app/ent/user"
)

// RegistrarCreate is the builder for creating a Registrar entity.
type RegistrarCreate struct {
	config
	mutation *RegistrarMutation
	hooks    []Hook
}

// SetName sets the Name field.
func (rc *RegistrarCreate) SetName(s string) *RegistrarCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetEdgesOfUserID sets the EdgesOfUser edge to User by id.
func (rc *RegistrarCreate) SetEdgesOfUserID(id int) *RegistrarCreate {
	rc.mutation.SetEdgesOfUserID(id)
	return rc
}

// SetNillableEdgesOfUserID sets the EdgesOfUser edge to User by id if the given value is not nil.
func (rc *RegistrarCreate) SetNillableEdgesOfUserID(id *int) *RegistrarCreate {
	if id != nil {
		rc = rc.SetEdgesOfUserID(*id)
	}
	return rc
}

// SetEdgesOfUser sets the EdgesOfUser edge to User.
func (rc *RegistrarCreate) SetEdgesOfUser(u *User) *RegistrarCreate {
	return rc.SetEdgesOfUserID(u.ID)
}

// Mutation returns the RegistrarMutation object of the builder.
func (rc *RegistrarCreate) Mutation() *RegistrarMutation {
	return rc.mutation
}

// Save creates the Registrar in the database.
func (rc *RegistrarCreate) Save(ctx context.Context) (*Registrar, error) {
	if _, ok := rc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "Name", err: errors.New("ent: missing required field \"Name\"")}
	}
	var (
		err  error
		node *Registrar
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegistrarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RegistrarCreate) SaveX(ctx context.Context) *Registrar {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RegistrarCreate) sqlSave(ctx context.Context) (*Registrar, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *RegistrarCreate) createSpec() (*Registrar, *sqlgraph.CreateSpec) {
	var (
		r     = &Registrar{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: registrar.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: registrar.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: registrar.FieldName,
		})
		r.Name = value
	}
	if nodes := rc.mutation.EdgesOfUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   registrar.EdgesOfUserTable,
			Columns: []string{registrar.EdgesOfUserColumn},
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
	return r, _spec
}
