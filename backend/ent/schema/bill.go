package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"

)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.String("Amount").NotEmpty(),
		field.Time("Date"),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("EdgesOfPaytype", Paytype.Type).Ref("EdgesOfBills").Unique(),
		edge.From("EdgesOfOfficer", Financier.Type).Ref("EdgesOfBills").Unique(),
		edge.From("EdgesOfTreatment", Unpaybill.Type).Ref("EdgesOfBills").Unique(),
	}
}
