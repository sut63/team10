package schema

import (
	"errors"
	"regexp"

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
		field.String("Name").NotEmpty().Match(regexp.MustCompile("[ก-ฮ]")),
		field.String("Idcardnumber").MinLen(13).MaxLen(13).Unique().Match(regexp.MustCompile("[0-9]")),
		field.Int("Age").Min(0),
		field.String("Disease").NotEmpty().Match(regexp.MustCompile("[ก-ฮ]")),
		field.String("Allergic").NotEmpty().Match(regexp.MustCompile("[a-z]")),
		field.String("Phonenumber").MinLen(10).MaxLen(10).
			Validate(func(s string) error {
				match, _ := regexp.MatchString("[0]\\d", s)
				if !match {
					return errors.New("เบอร์โทรต้องเป็นตัวเลขและขึ้นต้นด้วย 0")
				}
				return nil
			}),
		field.String("Email").NotEmpty().Match(regexp.MustCompile("[a-zA-Z0-9]+@([a-zA-Z0-9]+.)+[A-Za-z]")),
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
