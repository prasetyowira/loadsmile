package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// holds the schema definition for the Ingredient entity.
type Ingredient struct {
	ent.Schema
}

// Fields of the Ingredient.
func (Ingredient) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").
			MaxLen(26).
			NotEmpty().
			Unique().
			Immutable(),
		field.String("title"),
		field.Time("best_before"),
		field.Time("use_by").
			Optional().
			Nillable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Ingredient.
func (Ingredient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipes", Recipe.Type).
			Ref("ingredients"),
	}
}
