package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	//"github.com/facebookincubator/ent/schema/field"
)

// Doctor holds the schema definition for the Doctor entity.
type Doctor struct {
	ent.Schema
}

// Fields of the Doctor.
func (Doctor) Fields() []ent.Field {
	return nil
}

// Edges of the Doctor.
func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("doctorinfo", Doctorinfo.Type).Ref("doctor").Unique(),
		edge.From("user", User.Type).Ref("doctor").Unique(),
		edge.To("treatment", Treatment.Type).StorageKey(edge.Column("doctor_id")),
	}
}
