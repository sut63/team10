package schema

import (
	"regexp"

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
		field.String("Symptom").NotEmpty().Match(regexp.MustCompile("[ก-ฮ]")),
		field.String("Treat").NotEmpty().Match(regexp.MustCompile("[ก-ฮ]")),
		field.String("Medicine").NotEmpty().Match(regexp.MustCompile("[ก-ฮ]")),
		field.Time("Datetreat"),
	}
}

// Edges of the Treatment.
func (Treatment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("EdgesOfTypetreatment", Typetreatment.Type).Ref("EdgesOfTreatment").Unique(),
		edge.From("EdgesOfPatientrecord", Patientrecord.Type).Ref("EdgesOfTreatment").Unique(),
		edge.From("EdgesOfDoctor", Doctor.Type).Ref("EdgesOfTreatment").Unique(),
		edge.To("EdgesOfUnpaybills", Unpaybill.Type).StorageKey(edge.Column("treatment_id")).Unique(),
	}
}
