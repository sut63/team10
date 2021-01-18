// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team10/app/ent/bill"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/paytype"
	"github.com/team10/app/ent/unpaybill"
)

// Bill is the model entity for the Bill schema.
type Bill struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "Amount" field.
	Amount string `json:"Amount,omitempty"`
	// Date holds the value of the "Date" field.
	Date time.Time `json:"Date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BillQuery when eager-loading is set.
	Edges        BillEdges `json:"edges"`
	officer_id   *int
	paytype_id   *int
	treatment_id *int
}

// BillEdges holds the relations/edges for other nodes in the graph.
type BillEdges struct {
	// EdgesOfPaytype holds the value of the EdgesOfPaytype edge.
	EdgesOfPaytype *Paytype
	// EdgesOfOfficer holds the value of the EdgesOfOfficer edge.
	EdgesOfOfficer *Financier
	// EdgesOfTreatment holds the value of the EdgesOfTreatment edge.
	EdgesOfTreatment *Unpaybill
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// EdgesOfPaytypeOrErr returns the EdgesOfPaytype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) EdgesOfPaytypeOrErr() (*Paytype, error) {
	if e.loadedTypes[0] {
		if e.EdgesOfPaytype == nil {
			// The edge EdgesOfPaytype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: paytype.Label}
		}
		return e.EdgesOfPaytype, nil
	}
	return nil, &NotLoadedError{edge: "EdgesOfPaytype"}
}

// EdgesOfOfficerOrErr returns the EdgesOfOfficer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) EdgesOfOfficerOrErr() (*Financier, error) {
	if e.loadedTypes[1] {
		if e.EdgesOfOfficer == nil {
			// The edge EdgesOfOfficer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: financier.Label}
		}
		return e.EdgesOfOfficer, nil
	}
	return nil, &NotLoadedError{edge: "EdgesOfOfficer"}
}

// EdgesOfTreatmentOrErr returns the EdgesOfTreatment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) EdgesOfTreatmentOrErr() (*Unpaybill, error) {
	if e.loadedTypes[2] {
		if e.EdgesOfTreatment == nil {
			// The edge EdgesOfTreatment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: unpaybill.Label}
		}
		return e.EdgesOfTreatment, nil
	}
	return nil, &NotLoadedError{edge: "EdgesOfTreatment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bill) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Amount
		&sql.NullTime{},   // Date
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Bill) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // officer_id
		&sql.NullInt64{}, // paytype_id
		&sql.NullInt64{}, // treatment_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bill fields.
func (b *Bill) assignValues(values ...interface{}) error {
	if m, n := len(values), len(bill.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Amount", values[0])
	} else if value.Valid {
		b.Amount = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field Date", values[1])
	} else if value.Valid {
		b.Date = value.Time
	}
	values = values[2:]
	if len(values) == len(bill.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field officer_id", value)
		} else if value.Valid {
			b.officer_id = new(int)
			*b.officer_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field paytype_id", value)
		} else if value.Valid {
			b.paytype_id = new(int)
			*b.paytype_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field treatment_id", value)
		} else if value.Valid {
			b.treatment_id = new(int)
			*b.treatment_id = int(value.Int64)
		}
	}
	return nil
}

// QueryEdgesOfPaytype queries the EdgesOfPaytype edge of the Bill.
func (b *Bill) QueryEdgesOfPaytype() *PaytypeQuery {
	return (&BillClient{config: b.config}).QueryEdgesOfPaytype(b)
}

// QueryEdgesOfOfficer queries the EdgesOfOfficer edge of the Bill.
func (b *Bill) QueryEdgesOfOfficer() *FinancierQuery {
	return (&BillClient{config: b.config}).QueryEdgesOfOfficer(b)
}

// QueryEdgesOfTreatment queries the EdgesOfTreatment edge of the Bill.
func (b *Bill) QueryEdgesOfTreatment() *UnpaybillQuery {
	return (&BillClient{config: b.config}).QueryEdgesOfTreatment(b)
}

// Update returns a builder for updating this Bill.
// Note that, you need to call Bill.Unwrap() before calling this method, if this Bill
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bill) Update() *BillUpdateOne {
	return (&BillClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Bill) Unwrap() *Bill {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bill is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bill) String() string {
	var builder strings.Builder
	builder.WriteString("Bill(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", Amount=")
	builder.WriteString(b.Amount)
	builder.WriteString(", Date=")
	builder.WriteString(b.Date.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Bills is a parsable slice of Bill.
type Bills []*Bill

func (b Bills) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
