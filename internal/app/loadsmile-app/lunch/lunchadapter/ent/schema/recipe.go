package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Recipe struct {
	ent.Schema
}

// Fields of the Todo.
func (Recipe) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").
			MaxLen(26).
			NotEmpty().
			Unique().
			Immutable(),
		field.String("title"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Todo.
func (Recipe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ingredients", Ingredient.Type),
	}
}
