package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"

)

// Unpaybill holds the schema definition for the Unpaybill entity.
type Unpaybill struct {
	ent.Schema
}

// Fields of the Unpaybill.
func (Unpaybill) Fields() []ent.Field {
	return []ent.Field{
		field.String("Status").NotEmpty(),
	}
}

// Edges of the Unpaybill.
func (Unpaybill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("EdgesOfTreatment", Treatment.Type).Ref("EdgesOfUnpaybills").Unique(),
		edge.To("EdgesOfBills",Bill.Type).StorageKey(edge.Column("treatment_id")).Unique(),
	}
}
