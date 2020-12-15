package schema

import "github.com/facebookincubator/ent"

// Patientrights holds the schema definition for the Patientrights entity.
type Patientrights struct {
	ent.Schema
}

// Fields of the Patientrights.
func (Patientrights) Fields() []ent.Field {
	return nil
}

// Edges of the Patientrights.
func (Patientrights) Edges() []ent.Edge {
	return nil
}
