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
		field.Floats("้hight").NotEmpty(),            //ส่วนสูง[เซนติเมตร]
		field.Floats("weight").NotEmpty(),            //น้ำหนัก[กิโลกรัม]
		field.Floats("้temperature").NotEmpty(),      //อุณภูมิร่างกาย[เซลเซียส]
		field.Int("้pulse").NotEmpty(),               //ชีพจร[ครั้งต่อวินาที]
		field.Int("้respiration").NotEmpty(),         //การหายใจ[ครั้งต่อนาที]
		field.Int("้bloodpressure").NotEmpty(),       //ความดันโลหิต[มิลลิเมตรปรอท]
		field.Floats("้oxygensaturation").NotEmpty(), //ออกซิเจนในเลือด[เปอร์เซ็นต์]
		field.Time("datetime"),
	}
}

// Edges of the Historytaking.
func (Historytaking) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("ผู้ป่วย", ผู้ป่วย.Type).Ref("historytaking").Unique(),
		edge.From("hospitaldepartment", Hospitaldepartment.Type).Ref("historytaking").Unique(),
		edge.From("symptomseverity", Symptomseverity.Type).Ref("historytaking").Unique(),
		edge.From("nurse", Nurse.Type).Ref("historytaking").Unique(),
	}
}
