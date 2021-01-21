package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Patientrecord holds the schema definition for the Patientrecord entity.
type Patientrecord struct {
	ent.Schema
}

// Fields of the Patientrecord.
func (Patientrecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").NotEmpty(),
		field.String("Idcardnumber").MinLen(13).MaxLen(13),
		field.Int("Age").Min(0),
		field.String("Disease").NotEmpty(),
		field.String("Allergic").NotEmpty(),
		field.String("Phonenumber").MinLen(10).MaxLen(10),
		field.String("Email").NotEmpty(),
		field.String("Home").NotEmpty(),
		field.Time("Date"),
	}
}

// Edges of the Patientrecord.
func (Patientrecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("EdgesOfGender", Gender.Type).Ref("EdgesOfPatientrecord").Unique(),
		edge.From("EdgesOfBloodtype", Bloodtype.Type).Ref("EdgesOfPatientrecord").Unique(),
		edge.From("EdgesOfMedicalrecordstaff", Medicalrecordstaff.Type).Ref("EdgesOfPatientrecord").Unique(),
		edge.From("EdgesOfPrename", Prename.Type).Ref("EdgesOfPatientrecord").Unique(),
		edge.To("EdgesOfHistorytaking", Historytaking.Type).StorageKey(edge.Column("patientrecord_id")),
		edge.To("EdgesOfTreatment", Treatment.Type).StorageKey(edge.Column("patientrecord_id")),
		edge.To("EdgesOfPatientrecordPatientrights", Patientrights.Type).StorageKey(edge.Column("patientrecord_id")),
	}
}
