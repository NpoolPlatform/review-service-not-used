package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	constant "github.com/NpoolPlatform/review-service/pkg/const" //nolint

	"github.com/google/uuid"
)

// Review holds the schema definition for the Review object.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("object_type"),
		field.String("domain"),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("object_id", uuid.UUID{}),
		field.UUID("reviewer_id", uuid.UUID{}),
		field.Enum("state").
			Values(constant.StateWait,
				constant.StateApproved,
				constant.StateRejected),
		field.String("message"),
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

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return nil
}

// Indexs of the Review.
func (Review) Indexs() []ent.Index {
	return nil
}
