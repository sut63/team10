package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Symptomseverity holds the schema definition for the Symptomseverity entity.
type Symptomseverity struct {
	ent.Schema
}

// Fields of the Symptomseverity.
func (Symptomseverity) Fields() []ent.Field {
	return []ent.Field{
		field.String("symptomseverity"),
	}
}

// Edges of the Symptomseverity.
func (Symptomseverity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("historytaking", Historytaking.Type).StorageKey(edge.Column("symptomseverity_id")),
	}
}
