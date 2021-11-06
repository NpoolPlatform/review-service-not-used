package schema

import "entgo.io/ent"

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return nil
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return nil
}
