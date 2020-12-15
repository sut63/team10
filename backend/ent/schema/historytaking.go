package schema

import "github.com/facebookincubator/ent"

// HistoryTaking holds the schema definition for the HistoryTaking entity.
type HistoryTaking struct {
	ent.Schema
}

// Fields of the HistoryTaking.
func (HistoryTaking) Fields() []ent.Field {
	return nil
}

// Edges of the HistoryTaking.
func (HistoryTaking) Edges() []ent.Edge {
	return nil
}
