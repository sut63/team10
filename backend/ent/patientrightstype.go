// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/theuo/app/ent/abilitypatientrights"
	"github.com/theuo/app/ent/patientrightstype"
)

// Patientrightstype is the model entity for the Patientrightstype schema.
type Patientrightstype struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Permission holds the value of the "Permission" field.
	Permission string `json:"Permission,omitempty"`
	// PermissionArea holds the value of the "PermissionArea" field.
	PermissionArea string `json:"PermissionArea,omitempty"`
	// Responsible holds the value of the "Responsible" field.
	Responsible string `json:"Responsible,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientrightstypeQuery when eager-loading is set.
	Edges                   PatientrightstypeEdges `json:"edges"`
	Abilitypatientrights_id *int
}

// PatientrightstypeEdges holds the relations/edges for other nodes in the graph.
type PatientrightstypeEdges struct {
	// PatientrightstypePatientrights holds the value of the PatientrightstypePatientrights edge.
	PatientrightstypePatientrights []*Patientrights
	// PatientrightstypeAbilitypatientrights holds the value of the PatientrightstypeAbilitypatientrights edge.
	PatientrightstypeAbilitypatientrights *Abilitypatientrights
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PatientrightstypePatientrightsOrErr returns the PatientrightstypePatientrights value or an error if the edge
// was not loaded in eager-loading.
func (e PatientrightstypeEdges) PatientrightstypePatientrightsOrErr() ([]*Patientrights, error) {
	if e.loadedTypes[0] {
		return e.PatientrightstypePatientrights, nil
	}
	return nil, &NotLoadedError{edge: "PatientrightstypePatientrights"}
}

// PatientrightstypeAbilitypatientrightsOrErr returns the PatientrightstypeAbilitypatientrights value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientrightstypeEdges) PatientrightstypeAbilitypatientrightsOrErr() (*Abilitypatientrights, error) {
	if e.loadedTypes[1] {
		if e.PatientrightstypeAbilitypatientrights == nil {
			// The edge PatientrightstypeAbilitypatientrights was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: abilitypatientrights.Label}
		}
		return e.PatientrightstypeAbilitypatientrights, nil
	}
	return nil, &NotLoadedError{edge: "PatientrightstypeAbilitypatientrights"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Patientrightstype) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Permission
		&sql.NullString{}, // PermissionArea
		&sql.NullString{}, // Responsible
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Patientrightstype) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // Abilitypatientrights_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Patientrightstype fields.
func (pa *Patientrightstype) assignValues(values ...interface{}) error {
	if m, n := len(values), len(patientrightstype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Permission", values[0])
	} else if value.Valid {
		pa.Permission = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PermissionArea", values[1])
	} else if value.Valid {
		pa.PermissionArea = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Responsible", values[2])
	} else if value.Valid {
		pa.Responsible = value.String
	}
	values = values[3:]
	if len(values) == len(patientrightstype.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Abilitypatientrights_id", value)
		} else if value.Valid {
			pa.Abilitypatientrights_id = new(int)
			*pa.Abilitypatientrights_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPatientrightstypePatientrights queries the PatientrightstypePatientrights edge of the Patientrightstype.
func (pa *Patientrightstype) QueryPatientrightstypePatientrights() *PatientrightsQuery {
	return (&PatientrightstypeClient{config: pa.config}).QueryPatientrightstypePatientrights(pa)
}

// QueryPatientrightstypeAbilitypatientrights queries the PatientrightstypeAbilitypatientrights edge of the Patientrightstype.
func (pa *Patientrightstype) QueryPatientrightstypeAbilitypatientrights() *AbilitypatientrightsQuery {
	return (&PatientrightstypeClient{config: pa.config}).QueryPatientrightstypeAbilitypatientrights(pa)
}

// Update returns a builder for updating this Patientrightstype.
// Note that, you need to call Patientrightstype.Unwrap() before calling this method, if this Patientrightstype
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Patientrightstype) Update() *PatientrightstypeUpdateOne {
	return (&PatientrightstypeClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Patientrightstype) Unwrap() *Patientrightstype {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Patientrightstype is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Patientrightstype) String() string {
	var builder strings.Builder
	builder.WriteString("Patientrightstype(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", Permission=")
	builder.WriteString(pa.Permission)
	builder.WriteString(", PermissionArea=")
	builder.WriteString(pa.PermissionArea)
	builder.WriteString(", Responsible=")
	builder.WriteString(pa.Responsible)
	builder.WriteByte(')')
	return builder.String()
}

// Patientrightstypes is a parsable slice of Patientrightstype.
type Patientrightstypes []*Patientrightstype

func (pa Patientrightstypes) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
