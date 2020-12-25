package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
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
	return []ent.Edge{
		edge.To("bills", Bill.Type).StorageKey(edge.Column("treatment_id")).Unique(),
		edge.From("typetreatment", Typetreatment.Type).Ref("typetreatmenttreatment").Unique(),
		edge.From("doctorinfo", Doctorinfo.Type).Ref("doctorinfotreatment").Unique(),
		edge.From("patientrecord", Patientrecord.Type).Ref("patientrecordtreatment").Unique(),
	}
}
}
