package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edge"

)

// Paytype holds the schema definition for the Paytype entity.
type Paytype struct {
	ent.Schema
}

// Fields of the Paytype.
func (Paytype) Fields() []ent.Field {
	return []ent.Field{
		field.String("paytype").NotEmpty(),
	}
}

// Edges of the Paytype.
func (Paytype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bills", Bill.Type).StorageKey(edge.Column("paytype_id")),
	}
}
