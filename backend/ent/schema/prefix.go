package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Prefix holds the schema definition for the Prefix entity.
type Prefix struct {
	ent.Schema
}

// Fields of the Prefix.
func (Prefix) Fields() []ent.Field {
	return nil
}

// Edges of the Prefix.
func (Prefix) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("patientrecord", Patientrecord.Type).StorageKey(edge.Column("prefix_id")),
	}
}
