// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team10/app/ent/registrar"
	"github.com/team10/app/ent/user"
)

// Registrar is the model entity for the Registrar schema.
type Registrar struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RegistrarQuery when eager-loading is set.
	Edges   RegistrarEdges `json:"edges"`
	user_id *int
}

// RegistrarEdges holds the relations/edges for other nodes in the graph.
type RegistrarEdges struct {
	// EdgesOfUser holds the value of the EdgesOfUser edge.
	EdgesOfUser *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EdgesOfUserOrErr returns the EdgesOfUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RegistrarEdges) EdgesOfUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.EdgesOfUser == nil {
			// The edge EdgesOfUser was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.EdgesOfUser, nil
	}
	return nil, &NotLoadedError{edge: "EdgesOfUser"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Registrar) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Registrar) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Registrar fields.
func (r *Registrar) assignValues(values ...interface{}) error {
	if m, n := len(values), len(registrar.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Name", values[0])
	} else if value.Valid {
		r.Name = value.String
	}
	values = values[1:]
	if len(values) == len(registrar.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_id", value)
		} else if value.Valid {
			r.user_id = new(int)
			*r.user_id = int(value.Int64)
		}
	}
	return nil
}

// QueryEdgesOfUser queries the EdgesOfUser edge of the Registrar.
func (r *Registrar) QueryEdgesOfUser() *UserQuery {
	return (&RegistrarClient{config: r.config}).QueryEdgesOfUser(r)
}

// Update returns a builder for updating this Registrar.
// Note that, you need to call Registrar.Unwrap() before calling this method, if this Registrar
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Registrar) Update() *RegistrarUpdateOne {
	return (&RegistrarClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Registrar) Unwrap() *Registrar {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Registrar is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Registrar) String() string {
	var builder strings.Builder
	builder.WriteString("Registrar(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", Name=")
	builder.WriteString(r.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Registrars is a parsable slice of Registrar.
type Registrars []*Registrar

func (r Registrars) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
