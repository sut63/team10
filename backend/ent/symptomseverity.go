// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/b6109868/app/ent/symptomseverity"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Symptomseverity is the model entity for the Symptomseverity schema.
type Symptomseverity struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Symptomseverity holds the value of the "symptomseverity" field.
	Symptomseverity string `json:"symptomseverity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SymptomseverityQuery when eager-loading is set.
	Edges SymptomseverityEdges `json:"edges"`
}

// SymptomseverityEdges holds the relations/edges for other nodes in the graph.
type SymptomseverityEdges struct {
	// Historytaking holds the value of the historytaking edge.
	Historytaking []*Historytaking
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HistorytakingOrErr returns the Historytaking value or an error if the edge
// was not loaded in eager-loading.
func (e SymptomseverityEdges) HistorytakingOrErr() ([]*Historytaking, error) {
	if e.loadedTypes[0] {
		return e.Historytaking, nil
	}
	return nil, &NotLoadedError{edge: "historytaking"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Symptomseverity) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // symptomseverity
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Symptomseverity fields.
func (s *Symptomseverity) assignValues(values ...interface{}) error {
	if m, n := len(values), len(symptomseverity.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field symptomseverity", values[0])
	} else if value.Valid {
		s.Symptomseverity = value.String
	}
	return nil
}

// QueryHistorytaking queries the historytaking edge of the Symptomseverity.
func (s *Symptomseverity) QueryHistorytaking() *HistorytakingQuery {
	return (&SymptomseverityClient{config: s.config}).QueryHistorytaking(s)
}

// Update returns a builder for updating this Symptomseverity.
// Note that, you need to call Symptomseverity.Unwrap() before calling this method, if this Symptomseverity
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Symptomseverity) Update() *SymptomseverityUpdateOne {
	return (&SymptomseverityClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Symptomseverity) Unwrap() *Symptomseverity {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Symptomseverity is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Symptomseverity) String() string {
	var builder strings.Builder
	builder.WriteString("Symptomseverity(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", symptomseverity=")
	builder.WriteString(s.Symptomseverity)
	builder.WriteByte(')')
	return builder.String()
}

// Symptomseverities is a parsable slice of Symptomseverity.
type Symptomseverities []*Symptomseverity

func (s Symptomseverities) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
