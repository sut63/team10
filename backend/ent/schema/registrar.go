package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Registrar holds the schema definition for the Registrar entity.
type Registrar struct {
	ent.Schema
}

// Fields of the Registrar.
func (Registrar) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
	}
}

// Edges of the Registrar.
func (Registrar) Edges() []ent.Edge {
	return []ent.Edge{
		
		edge.From("user", User.Type).Ref("user2registrar").Unique(),
	}
}
