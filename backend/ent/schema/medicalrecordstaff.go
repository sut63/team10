package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Medicalrecordstaff holds the schema definition for the Medicalrecordstaff entity.
type Medicalrecordstaff struct {
	ent.Schema
}

// Fields of the Medicalrecordstaff.
func (Medicalrecordstaff) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
	}
}

// Edges of the Medicalrecordstaff.
func (Medicalrecordstaff) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("EdgesOfPatientrecord", Patientrecord.Type).StorageKey(edge.Column("medicalrecordstaff_id")),
		edge.To("EdgesOfMedicalrecordstaffPatientrights", Patientrights.Type).StorageKey(edge.Column("medicalrecordstaff_id")),
		edge.From("EdgesOfUser", User.Type).Ref("EdgesOfMedicalrecordstaff").Unique(),
	}
}
