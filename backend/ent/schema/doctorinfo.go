package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Doctorinfo holds the schema definition for the Doctorinfo entity.
type Doctorinfo struct {
	ent.Schema
}

// Fields of the Doctorinfo.
func (Doctorinfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("doctorname").NotEmpty(),
		field.String("doctorsurname").NotEmpty(),
		field.String("telephonenumber").NotEmpty(),
		field.String("licensenumber").NotEmpty(),
	}
}

// Edges of the Doctorinfo.
func (Doctorinfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("EdgesOfDepartment", Department.Type).Ref("EdgesOfDepartment2doctorinfo").Unique(),
		edge.From("EdgesOfEducationlevel", Educationlevel.Type).Ref("EdgesOfEducationlevel2doctorinfo").Unique(),
		edge.From("EdgesOfOfficeroom", Officeroom.Type).Ref("EdgesOfOfficeroom2doctorinfo").Unique(),
		edge.From("EdgesOfPrename", Prename.Type).Ref("EdgesOfPrename2doctorinfo").Unique(),
		edge.To("EdgesOfDoctor", Doctor.Type).StorageKey(edge.Column("doctorinfo_id")).Unique(),
	}
}
