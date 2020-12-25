package schema

import (
    
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Patientrights holds the schema definition for the Patientrights entity.
type Patientrights struct {
	ent.Schema
}

// Fields of the Patientrights.
func (Patientrights) Fields() []ent.Field {
	return []ent.Field{
		field.String("PermissionDate"),
		
			
    }
}
// Edges of the Patientrights.
func (Patientrights) Edges() []ent.Edge {
	return []ent.Edge{

		
		edge.From("PatientrightsPatientrightstype", Patientrightstype.Type).
			Ref("PatientrightstypePatientrights").
			Unique(),

		/*
		edge.To("Patientrights___", ___.Type),

		*/

		edge.From("PatientrightsInsurance", Insurance.Type).
			Ref("InsurancePatientrights").
			Unique(),

		edge.From("PatientrightsPatientrecord", Patientrecord.Type).
			Ref("PatientrecordPatientrights").
			Unique(),

		edge.From("PatientrightsUser", User.Type).
			Ref("UserPatientrights").
			Unique(),

	}
}
