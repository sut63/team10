package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("images"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("EdgesOfFinancier", Financier.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.To("EdgesOfNurse", Nurse.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.To("EdgesOfUserPatientrights", Patientrights.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.To("EdgesOfMedicalrecordstaff", Medicalrecordstaff.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.To("EdgesOfUser2registrar", Registrar.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.To("EdgesOfDoctor", Doctor.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.From("EdgesOfUserstatus", Userstatus.Type).Ref("EdgesOfUser").Unique(),
	}
}
