package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Financialofficer holds the schema definition for the Financialofficer entity.
type Financialofficer struct {
	ent.Schema
}

// Fields of the Financialofficer.
func (Financialofficer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty(),
	}
}

// Edges of the Financialofficer.
func (Financialofficer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bills", Bill.Type).StorageKey(edge.Column("officer_id")),
	}
}
