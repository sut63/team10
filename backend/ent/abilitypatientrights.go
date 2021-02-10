// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team10/app/ent/abilitypatientrights"
)

// Abilitypatientrights is the model entity for the Abilitypatientrights schema.
type Abilitypatientrights struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Operative holds the value of the "Operative" field.
	Operative int `json:"Operative,omitempty"`
	// MedicalSupplies holds the value of the "MedicalSupplies" field.
	MedicalSupplies int `json:"MedicalSupplies,omitempty"`
	// Examine holds the value of the "Examine" field.
	Examine int `json:"Examine,omitempty"`
	// StayInHospital holds the value of the "StayInHospital" field.
	StayInHospital int `json:"StayInHospital,omitempty"`
	// Check holds the value of the "check" field.
	Check string `json:"check,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AbilitypatientrightsQuery when eager-loading is set.
	Edges AbilitypatientrightsEdges `json:"edges"`
}

// AbilitypatientrightsEdges holds the relations/edges for other nodes in the graph.
type AbilitypatientrightsEdges struct {
	// EdgesOfAbilitypatientrightsPatientrights holds the value of the EdgesOfAbilitypatientrightsPatientrights edge.
	EdgesOfAbilitypatientrightsPatientrights []*Patientrights
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EdgesOfAbilitypatientrightsPatientrightsOrErr returns the EdgesOfAbilitypatientrightsPatientrights value or an error if the edge
// was not loaded in eager-loading.
func (e AbilitypatientrightsEdges) EdgesOfAbilitypatientrightsPatientrightsOrErr() ([]*Patientrights, error) {
	if e.loadedTypes[0] {
		return e.EdgesOfAbilitypatientrightsPatientrights, nil
	}
	return nil, &NotLoadedError{edge: "EdgesOfAbilitypatientrightsPatientrights"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Abilitypatientrights) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // Operative
		&sql.NullInt64{},  // MedicalSupplies
		&sql.NullInt64{},  // Examine
		&sql.NullInt64{},  // StayInHospital
		&sql.NullString{}, // check
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Abilitypatientrights fields.
func (a *Abilitypatientrights) assignValues(values ...interface{}) error {
	if m, n := len(values), len(abilitypatientrights.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Operative", values[0])
	} else if value.Valid {
		a.Operative = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field MedicalSupplies", values[1])
	} else if value.Valid {
		a.MedicalSupplies = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Examine", values[2])
	} else if value.Valid {
		a.Examine = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field StayInHospital", values[3])
	} else if value.Valid {
		a.StayInHospital = int(value.Int64)
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field check", values[4])
	} else if value.Valid {
		a.Check = value.String
	}
	return nil
}

// QueryEdgesOfAbilitypatientrightsPatientrights queries the EdgesOfAbilitypatientrightsPatientrights edge of the Abilitypatientrights.
func (a *Abilitypatientrights) QueryEdgesOfAbilitypatientrightsPatientrights() *PatientrightsQuery {
	return (&AbilitypatientrightsClient{config: a.config}).QueryEdgesOfAbilitypatientrightsPatientrights(a)
}

// Update returns a builder for updating this Abilitypatientrights.
// Note that, you need to call Abilitypatientrights.Unwrap() before calling this method, if this Abilitypatientrights
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Abilitypatientrights) Update() *AbilitypatientrightsUpdateOne {
	return (&AbilitypatientrightsClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Abilitypatientrights) Unwrap() *Abilitypatientrights {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Abilitypatientrights is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Abilitypatientrights) String() string {
	var builder strings.Builder
	builder.WriteString("Abilitypatientrights(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", Operative=")
	builder.WriteString(fmt.Sprintf("%v", a.Operative))
	builder.WriteString(", MedicalSupplies=")
	builder.WriteString(fmt.Sprintf("%v", a.MedicalSupplies))
	builder.WriteString(", Examine=")
	builder.WriteString(fmt.Sprintf("%v", a.Examine))
	builder.WriteString(", StayInHospital=")
	builder.WriteString(fmt.Sprintf("%v", a.StayInHospital))
	builder.WriteString(", check=")
	builder.WriteString(a.Check)
	builder.WriteByte(')')
	return builder.String()
}

// AbilitypatientrightsSlice is a parsable slice of Abilitypatientrights.
type AbilitypatientrightsSlice []*Abilitypatientrights

func (a AbilitypatientrightsSlice) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
