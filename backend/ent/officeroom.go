// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team10/app/ent/officeroom"
)

// Officeroom is the model entity for the Officeroom schema.
type Officeroom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Roomnumber holds the value of the "roomnumber" field.
	Roomnumber int `json:"roomnumber,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OfficeroomQuery when eager-loading is set.
	Edges OfficeroomEdges `json:"edges"`
}

// OfficeroomEdges holds the relations/edges for other nodes in the graph.
type OfficeroomEdges struct {
	// Officeroom2doctorinfo holds the value of the officeroom2doctorinfo edge.
	Officeroom2doctorinfo []*Doctorinfo
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// Officeroom2doctorinfoOrErr returns the Officeroom2doctorinfo value or an error if the edge
// was not loaded in eager-loading.
func (e OfficeroomEdges) Officeroom2doctorinfoOrErr() ([]*Doctorinfo, error) {
	if e.loadedTypes[0] {
		return e.Officeroom2doctorinfo, nil
	}
	return nil, &NotLoadedError{edge: "officeroom2doctorinfo"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Officeroom) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // roomnumber
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Officeroom fields.
func (o *Officeroom) assignValues(values ...interface{}) error {
	if m, n := len(values), len(officeroom.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	o.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field roomnumber", values[0])
	} else if value.Valid {
		o.Roomnumber = int(value.Int64)
	}
	return nil
}

// QueryOfficeroom2doctorinfo queries the officeroom2doctorinfo edge of the Officeroom.
func (o *Officeroom) QueryOfficeroom2doctorinfo() *DoctorinfoQuery {
	return (&OfficeroomClient{config: o.config}).QueryOfficeroom2doctorinfo(o)
}

// Update returns a builder for updating this Officeroom.
// Note that, you need to call Officeroom.Unwrap() before calling this method, if this Officeroom
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Officeroom) Update() *OfficeroomUpdateOne {
	return (&OfficeroomClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (o *Officeroom) Unwrap() *Officeroom {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Officeroom is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Officeroom) String() string {
	var builder strings.Builder
	builder.WriteString("Officeroom(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", roomnumber=")
	builder.WriteString(fmt.Sprintf("%v", o.Roomnumber))
	builder.WriteByte(')')
	return builder.String()
}

// Officerooms is a parsable slice of Officeroom.
type Officerooms []*Officeroom

func (o Officerooms) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
