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
					return errors.New("doctor name must begin with uppercase")
				}
				return nil
			}),
		field.String("doctorsurname").NotEmpty().Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("surname must begin with uppercase")
				}
				return nil
			}),
		field.String("telephonenumber").NotEmpty().MaxLen(10).MinLen(10),
		field.String("licensenumber").NotEmpty().MaxLen(11).MinLen(11),
	}
}

// Edges of the Doctorinfo.
func (Doctorinfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("department", Department.Type).Ref("department2doctorinfo").Unique(),
		edge.From("educationlevel", Educationlevel.Type).Ref("educationlevel2doctorinfo").Unique(),
		edge.From("officeroom", Officeroom.Type).Ref("officeroom2doctorinfo").Unique(),
		edge.From("prename", Prename.Type).Ref("prename2doctorinfo").Unique(),
		edge.To("doctor", Doctor.Type).StorageKey(edge.Column("doctorinfo_id")).Unique(),
	}
}
