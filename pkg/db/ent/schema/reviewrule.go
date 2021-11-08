package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// ReviewRule holds the schema definition for the ReviewRule entity.
type ReviewRule struct {
	ent.Schema
}

// Fields of the ReviewRule.
func (ReviewRule) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("entity_type"),
		field.String("domain"),
		field.String("rules").
			Default("{}"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the ReviewRule.
func (ReviewRule) Edges() []ent.Edge {
	return nil
}

// Indexs of the ReviewRule.
func (ReviewRule) Indexs() []ent.Index {
	return []ent.Index{
		index.Fields("domain", "entity_type").
			Unique(),
	}
}
