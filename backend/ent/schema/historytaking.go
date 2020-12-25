package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Historytaking holds the schema definition for the Historytaking entity.
type Historytaking struct {
	ent.Schema
}

// Fields of the Historytaking.
func (Historytaking) Fields() []ent.Field {
	return []ent.Field{
		field.Floats("้hight"),            //ส่วนสูง[เซนติเมตร]
		field.Floats("weight"),            //น้ำหนัก[กิโลกรัม]
		field.Floats("temperature"),       //อุณภูมิร่างกาย[เซลเซียส]
		field.Int("้pulse"),               //ชีพจร[ครั้งต่อวินาที]
		field.Int("้respiration"),         //การหายใจ[ครั้งต่อนาที]
		field.Int("้bloodpressure"),       //ความดันโลหิต[มิลลิเมตรปรอท]
		field.Floats("oxygensaturation"), //ออกซิเจนในเลือด[เปอร์เซ็นต์]
		field.String("symptom"),           //อาการ
		field.Time("datetime"),
	}
}

// Edges of the Historytaking.
func (Historytaking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nurse", Nurse.Type).Ref("historytaking").Unique(),
		edge.From("department", Department.Type).Ref("historytaking").Unique(),
		edge.From("symptomseverity", Symptomseverity.Type).Ref("historytaking").Unique(),
		edge.From("patientrecord", Patientrecord.Type).Ref("historytaking").Unique(),
	}
}
