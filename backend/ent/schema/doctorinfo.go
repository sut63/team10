package schema

import (
	"errors"
	"regexp"
	"strings"

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
		field.String("doctorname").NotEmpty().Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("Doctor name must begin with uppercase")
				}
				return nil
			}),
		field.String("doctorsurname").NotEmpty().Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("Doctor surname must begin with uppercase")
				}
				return nil
			}),
		field.String("telephonenumber").NotEmpty().MinLen(10).MaxLen(10),
		field.String("licensenumber").NotEmpty().MinLen(11).MaxLen(11),
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
