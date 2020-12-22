package schema

import (
	"github.com/facebookincubator/ent"
	//"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Treatment holds the schema definition for the Treatment entity.
type Treatment struct {
	ent.Schema
}

// Fields of the Treatment.
func (Treatment) Fields() []ent.Field {
	return []ent.Field{
		field.String("Treatment").NotEmpty(),
		field.Time("Datetime"),
	}
}

// Edges of the Treatment.
func (Treatment) Edges() []ent.Edge {
	return nil
}
