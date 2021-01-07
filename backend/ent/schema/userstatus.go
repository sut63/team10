package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Userstatus holds the schema definition for the Userstatus entity.
type Userstatus struct {
	ent.Schema
}

// Fields of the Userstatus.
func (Userstatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("Userstatus"),
	}
}

// Edges of the Userstatus.
func (Userstatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).StorageKey(edge.Column("userstatus_id")),
	}
}
