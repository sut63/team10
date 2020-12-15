package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Officer holds the schema definition for the Officer entity.
type Officer struct {
	ent.Schema
}

// Fields of the Officer.
func (Officer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty(),
	}
}

// Edges of the Officer.
func (Officer) Edges() []ent.Edge {
	return nil
}
