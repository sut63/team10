package schema

import (
    
    "github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	_"github.com/facebook/ent/schema/edge"
)

// Patientrights holds the schema definition for the Patientrights entity.
type Patientrights struct {
	ent.Schema
}

// Fields of the Patientrights.
func (Patientrights) Fields() []ent.Field {
	return []ent.Field{
        field.String("PermissionNumber"),
			
    }
}
// Edges of the Patientrights.
func (Patientrights) Edges() []ent.Edge {
	return []ent.Edge{
/*
		edge.From("PatientrightsAbilitypatientrights", Abilitypatientrights.Type).
			Ref("AbilitypatientrightsPatientrights").
			Unique(),

		edge.From("Patientrights___", ___.Type).
			Ref("___Patientrights").
			Unique(),

		edge.From("PatientrightsPatientrightstype", Patientrightstype.Type).
			Ref("PatientrightstypePatientrights").
			Unique(),

		
		edge.To("Patientrights___", ___.Type),

		
*/
	}
}
