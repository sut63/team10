package schema
//เป็นส่วนประกอบของ Patientrightstype เพื่อบอกว่าสิทธิ์แบบใดมีคุณสมบัติใดบ้าง
import (
    
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Abilitypatientrights holds the schema definition for the Abilitypatientrights entity.
type Abilitypatientrights struct {
	ent.Schema
}

// Fields of the Abilitypatientrights.
func (Abilitypatientrights) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Operative"),//หัตถการ
		field.Int("MedicalSupplies"),//เวชภัณฑ์
		field.Int("Examine"),//ตรวจสุขภาพ และ ค่า แลป
			
    }
}

// Edges of the Abilitypatientrights.
func (Abilitypatientrights) Edges() []ent.Edge {
	return []ent.Edge{
		
		edge.To("EdgesOfAbilitypatientrightsPatientrightstype", Patientrightstype.Type).StorageKey(edge.Column("Abilitypatientrights_id")),//เป็นส่วนประกอบของ Patientrightstype
    }
}
