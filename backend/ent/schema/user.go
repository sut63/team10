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
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("financier", Financier.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.To("Nurse", Nurse.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.To("UserPatientrights", Patientrights.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.To("medicalrecordstaff", Medicalrecordstaff.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.To("user2registrar", Registrar.Type).StorageKey(edge.Column("user_id")).Unique(),
		edge.From("userstatus", Userstatus.Type).Ref("user").Unique(),
	}
}
