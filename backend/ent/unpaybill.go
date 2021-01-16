// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team10/app/ent/bill"
	"github.com/team10/app/ent/treatment"
	"github.com/team10/app/ent/unpaybill"
)

// Unpaybill is the model entity for the Unpaybill schema.
type Unpaybill struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "Status" field.
	Status string `json:"Status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UnpaybillQuery when eager-loading is set.
	Edges        UnpaybillEdges `json:"edges"`
	treatment_id *int
}

// UnpaybillEdges holds the relations/edges for other nodes in the graph.
type UnpaybillEdges struct {
	// EdgesOfTreatment holds the value of the EdgesOfTreatment edge.
	EdgesOfTreatment *Treatment
	// EdgesOfBills holds the value of the EdgesOfBills edge.
	EdgesOfBills *Bill
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EdgesOfTreatmentOrErr returns the EdgesOfTreatment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UnpaybillEdges) EdgesOfTreatmentOrErr() (*Treatment, error) {
	if e.loadedTypes[0] {
		if e.EdgesOfTreatment == nil {
			// The edge EdgesOfTreatment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: treatment.Label}
		}
		return e.EdgesOfTreatment, nil
	}
	return nil, &NotLoadedError{edge: "EdgesOfTreatment"}
}

// EdgesOfBillsOrErr returns the EdgesOfBills value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UnpaybillEdges) EdgesOfBillsOrErr() (*Bill, error) {
	if e.loadedTypes[1] {
		if e.EdgesOfBills == nil {
			// The edge EdgesOfBills was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: bill.Label}
		}
		return e.EdgesOfBills, nil
	}
	return nil, &NotLoadedError{edge: "EdgesOfBills"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Unpaybill) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Status
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Unpaybill) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // treatment_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Unpaybill fields.
func (u *Unpaybill) assignValues(values ...interface{}) error {
	if m, n := len(values), len(unpaybill.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Status", values[0])
	} else if value.Valid {
		u.Status = value.String
	}
	values = values[1:]
	if len(values) == len(unpaybill.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field treatment_id", value)
		} else if value.Valid {
			u.treatment_id = new(int)
			*u.treatment_id = int(value.Int64)
		}
	}
	return nil
}

// QueryEdgesOfTreatment queries the EdgesOfTreatment edge of the Unpaybill.
func (u *Unpaybill) QueryEdgesOfTreatment() *TreatmentQuery {
	return (&UnpaybillClient{config: u.config}).QueryEdgesOfTreatment(u)
}

// QueryEdgesOfBills queries the EdgesOfBills edge of the Unpaybill.
func (u *Unpaybill) QueryEdgesOfBills() *BillQuery {
	return (&UnpaybillClient{config: u.config}).QueryEdgesOfBills(u)
}

// Update returns a builder for updating this Unpaybill.
// Note that, you need to call Unpaybill.Unwrap() before calling this method, if this Unpaybill
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Unpaybill) Update() *UnpaybillUpdateOne {
	return (&UnpaybillClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *Unpaybill) Unwrap() *Unpaybill {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Unpaybill is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Unpaybill) String() string {
	var builder strings.Builder
	builder.WriteString("Unpaybill(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", Status=")
	builder.WriteString(u.Status)
	builder.WriteByte(')')
	return builder.String()
}

// Unpaybills is a parsable slice of Unpaybill.
type Unpaybills []*Unpaybill

func (u Unpaybills) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
