package schema

//เป็นส่วนประกอบของ Patientrightstype เพื่อบอกว่าสิทธิ์แบบใดมีคุณสมบัติใดบ้าง
import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Abilitypatientrights holds the schema definition for the Abilitypatientrights entity.
type Abilitypatientrights struct {
	ent.Schema
}

// Fields of the Abilitypatientrights.
func (Abilitypatientrights) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Operative").Positive().Max(101),       //หัตถการ
		field.Int("MedicalSupplies").Positive().Max(101), //เวชภัณฑ์
		field.Int("Examine").Positive().Max(101),         //ตรวจสุขภาพ และ ค่า แลป
		field.Int("StayInHospital").Positive().Max(101),  //ค่า นอนโรงพยาบาล
		
		field.String("check").Unique(),                   //field ควบคุม

	}
}

// Edges of the Abilitypatientrights.
func (Abilitypatientrights) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("EdgesOfAbilitypatientrightsPatientrights", Patientrights.Type).StorageKey(edge.Column("Abilitypatientrights_id")), //เป็นส่วนประกอบของ Patientrightstype
	}
}
