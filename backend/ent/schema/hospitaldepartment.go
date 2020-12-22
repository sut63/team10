package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Hospitaldepartment holds the schema definition for the Hospitaldepartment entity.
type Hospitaldepartment struct {
	ent.Schema
}

// Fields of the Hospitaldepartment.
func (Hospitaldepartment) Fields() []ent.Field {
	return []ent.Field{
		field.String("department"),
	}
}

// Edges of the Hospitaldepartment.
func (Hospitaldepartment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("historytaking", Historytaking.Type).StorageKey(edge.Column("hospitaldepartment_id")),
	}
}
