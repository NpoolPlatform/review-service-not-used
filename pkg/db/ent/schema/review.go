package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("entity_type"),
		field.String("domain"),
		field.UUID("object_id", uuid.UUID{}),
		field.UUID("reviewer_id", uuid.UUID{}),
		field.Enum("state").
			Values("wait", "approved", "rejected"),
		field.String("message"),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return nil
}
