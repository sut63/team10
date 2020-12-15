package schema

import "github.com/facebookincubator/ent"

// Gender holds the schema definition for the Gender entity.
type Gender struct {
	ent.Schema
}

// Fields of the Gender.
func (Gender) Fields() []ent.Field {
	return nil
}

// Edges of the Gender.
func (Gender) Edges() []ent.Edge {
	return nil
}
