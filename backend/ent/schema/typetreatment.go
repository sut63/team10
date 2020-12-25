package schema

import (
	"github.com/facebookincubator/ent"
	//"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Typetreatment holds the schema definition for the Typetreatment entity.
type Typetreatment struct {
	ent.Schema
}

// Fields of the Typetreatment.
func (Typetreatment) Fields() []ent.Field {
	return []ent.Field{
		field.String("Type").NotEmpty(),
	}
}

// Edges of the Typetreatment.
func (Typetreatment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("typetreatmenttreatment", Treatment.Type).StorageKey(edge.Column("typetreatment_id")),
	}
}
}
