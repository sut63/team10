package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Patientrecord holds the schema definition for the Patientrecord entity.
type Patientrecord struct {
	ent.Schema
}

// Fields of the Patientrecord.
func (Patientrecord) Fields() []ent.Field {
	return nil
}

// Edges of the Patientrecord.
func (Patientrecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("historytaking", Historytaking.Type).StorageKey(edge.Column("patientrecord_id")),
	}
}