// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/user"
)

// Medicalrecordstaff is the model entity for the Medicalrecordstaff schema.
type Medicalrecordstaff struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MedicalrecordstaffQuery when eager-loading is set.
	Edges   MedicalrecordstaffEdges `json:"edges"`
	user_id *int
}

// MedicalrecordstaffEdges holds the relations/edges for other nodes in the graph.
type MedicalrecordstaffEdges struct {
	// EdgesOfPatientrecord holds the value of the EdgesOfPatientrecord edge.
	EdgesOfPatientrecord []*Patientrecord
	// EdgesOfMedicalrecordstaffPatientrights holds the value of the EdgesOfMedicalrecordstaffPatientrights edge.
	EdgesOfMedicalrecordstaffPatientrights []*Patientrights
	// EdgesOfUser holds the value of the EdgesOfUser edge.
	EdgesOfUser *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// EdgesOfPatientrecordOrErr returns the EdgesOfPatientrecord value or an error if the edge
// was not loaded in eager-loading.
func (e MedicalrecordstaffEdges) EdgesOfPatientrecordOrErr() ([]*Patientrecord, error) {
	if e.loadedTypes[0] {
		return e.EdgesOfPatientrecord, nil
	}
	return nil, &NotLoadedError{edge: "EdgesOfPatientrecord"}
}

// EdgesOfMedicalrecordstaffPatientrightsOrErr returns the EdgesOfMedicalrecordstaffPatientrights value or an error if the edge
// was not loaded in eager-loading.
func (e MedicalrecordstaffEdges) EdgesOfMedicalrecordstaffPatientrightsOrErr() ([]*Patientrights, error) {
	if e.loadedTypes[1] {
		return e.EdgesOfMedicalrecordstaffPatientrights, nil
	}
	return nil, &NotLoadedError{edge: "EdgesOfMedicalrecordstaffPatientrights"}
}

// EdgesOfUserOrErr returns the EdgesOfUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalrecordstaffEdges) EdgesOfUserOrErr() (*User, error) {
	if e.loadedTypes[2] {
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
func (*Medicalrecordstaff) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Medicalrecordstaff) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Medicalrecordstaff fields.
func (m *Medicalrecordstaff) assignValues(values ...interface{}) error {
	if m, n := len(values), len(medicalrecordstaff.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Name", values[0])
	} else if value.Valid {
		m.Name = value.String
	}
	values = values[1:]
	if len(values) == len(medicalrecordstaff.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_id", value)
		} else if value.Valid {
			m.user_id = new(int)
			*m.user_id = int(value.Int64)
		}
	}
	return nil
}

// QueryEdgesOfPatientrecord queries the EdgesOfPatientrecord edge of the Medicalrecordstaff.
func (m *Medicalrecordstaff) QueryEdgesOfPatientrecord() *PatientrecordQuery {
	return (&MedicalrecordstaffClient{config: m.config}).QueryEdgesOfPatientrecord(m)
}

// QueryEdgesOfMedicalrecordstaffPatientrights queries the EdgesOfMedicalrecordstaffPatientrights edge of the Medicalrecordstaff.
func (m *Medicalrecordstaff) QueryEdgesOfMedicalrecordstaffPatientrights() *PatientrightsQuery {
	return (&MedicalrecordstaffClient{config: m.config}).QueryEdgesOfMedicalrecordstaffPatientrights(m)
}

// QueryEdgesOfUser queries the EdgesOfUser edge of the Medicalrecordstaff.
func (m *Medicalrecordstaff) QueryEdgesOfUser() *UserQuery {
	return (&MedicalrecordstaffClient{config: m.config}).QueryEdgesOfUser(m)
}

// Update returns a builder for updating this Medicalrecordstaff.
// Note that, you need to call Medicalrecordstaff.Unwrap() before calling this method, if this Medicalrecordstaff
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Medicalrecordstaff) Update() *MedicalrecordstaffUpdateOne {
	return (&MedicalrecordstaffClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Medicalrecordstaff) Unwrap() *Medicalrecordstaff {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Medicalrecordstaff is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Medicalrecordstaff) String() string {
	var builder strings.Builder
	builder.WriteString("Medicalrecordstaff(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", Name=")
	builder.WriteString(m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Medicalrecordstaffs is a parsable slice of Medicalrecordstaff.
type Medicalrecordstaffs []*Medicalrecordstaff

func (m Medicalrecordstaffs) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
