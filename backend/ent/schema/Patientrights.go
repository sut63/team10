package schema

//เป็น OUTPUT หลักของ B6103866
import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Patientrights holds the schema definition for the Patientrights entity.
type Patientrights struct {
	ent.Schema
}

// Fields of the Patientrights.
func (Patientrights) Fields() []ent.Field {
	return []ent.Field{
		field.Time("PermissionDate"), //เวลาที่สร้างสิทธิ์

		field.String("Permission").
			Validate(func(s string) error {
				match, _ := regexp.MatchString("[\\w]", s)

				if len(s) == 0 {
					return errors.New("กรุณากรอก เลขสิทธิ์")
				}
				if !match {
					return errors.New("เลขสิทธิ์ ไม่ถูกต้อง")
				}
				return nil
			}).NotEmpty(), //ชื่อสิทธิ์

		field.String("PermissionArea").
			Validate(func(s string) error {

				if len(s) == 0 {
					return errors.New("กรุณากรอก พื้นที่ที่ใช้สิทธิ์")
				}
				return nil
			}).NotEmpty(), //พื้นที่ หรือเขต หรือโรงพยาบาล ที่สามารถใช้สิทได้

		field.String("Responsible").
			Validate(func(s string) error {

				if len(s) == 0 {
					return errors.New("กรุณากรอก ผู้รับผิดชอบดูแลสิทธิ์")
				}
				return nil
			}).NotEmpty(), //เจ้าหน้าที่ประกัน

	}
}

// Edges of the Patientrights.
func (Patientrights) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("EdgesOfPatientrightsAbilitypatientrights", Abilitypatientrights.Type).
			Ref("EdgesOfAbilitypatientrightsPatientrights").
			Unique(),
		//บอกว่า สิทธิ์นี้มีความสามารถอะไรบ้าง

		edge.From("EdgesOfPatientrightsInsurance", Insurance.Type).
			Ref("EdgesOfInsurancePatientrights").
			Unique(),
		//ผู้จ่ายเงิน

		edge.From("EdgesOfPatientrightsPatientrecord", Patientrecord.Type).
			Ref("EdgesOfPatientrecordPatientrights").
			Unique(),
		//ผู้ป่วย ผู้รับสิทธิ์

		edge.From("EdgesOfPatientrightsMedicalrecordstaff", Medicalrecordstaff.Type).
			Ref("EdgesOfMedicalrecordstaffPatientrights").
			Unique(),
		//พนักงานผู้กรอกข้อมูล

	}
}
