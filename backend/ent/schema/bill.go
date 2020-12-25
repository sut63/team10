package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.String("Amount").NotEmpty(),
		field.Time("Date"),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("paytype", Paytype.Type).Ref("bills").Unique(),
		edge.From("officer", Financialofficer.Type).Ref("bills").Unique(),
		edge.From("treatment", Treatment.Type).Ref("bills").Unique(),
	}
}
