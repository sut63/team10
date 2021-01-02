// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/theuo/app/ent/paytype"
)

// Paytype is the model entity for the Paytype schema.
type Paytype struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Paytype holds the value of the "paytype" field.
	Paytype string `json:"paytype,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaytypeQuery when eager-loading is set.
	Edges PaytypeEdges `json:"edges"`
}

// PaytypeEdges holds the relations/edges for other nodes in the graph.
type PaytypeEdges struct {
	// Bills holds the value of the bills edge.
	Bills []*Bill
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BillsOrErr returns the Bills value or an error if the edge
// was not loaded in eager-loading.
func (e PaytypeEdges) BillsOrErr() ([]*Bill, error) {
	if e.loadedTypes[0] {
		return e.Bills, nil
	}
	return nil, &NotLoadedError{edge: "bills"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Paytype) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // paytype
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Paytype fields.
func (pa *Paytype) assignValues(values ...interface{}) error {
	if m, n := len(values), len(paytype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field paytype", values[0])
	} else if value.Valid {
		pa.Paytype = value.String
	}
	return nil
}

// QueryBills queries the bills edge of the Paytype.
func (pa *Paytype) QueryBills() *BillQuery {
	return (&PaytypeClient{config: pa.config}).QueryBills(pa)
}

// Update returns a builder for updating this Paytype.
// Note that, you need to call Paytype.Unwrap() before calling this method, if this Paytype
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Paytype) Update() *PaytypeUpdateOne {
	return (&PaytypeClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Paytype) Unwrap() *Paytype {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Paytype is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Paytype) String() string {
	var builder strings.Builder
	builder.WriteString("Paytype(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", paytype=")
	builder.WriteString(pa.Paytype)
	builder.WriteByte(')')
	return builder.String()
}

// Paytypes is a parsable slice of Paytype.
type Paytypes []*Paytype

func (pa Paytypes) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
