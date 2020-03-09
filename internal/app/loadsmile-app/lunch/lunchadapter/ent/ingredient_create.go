// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/ingredient"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/recipe"
)

// IngredientCreate is the builder for creating a Ingredient entity.
type IngredientCreate struct {
	config
	uid         *string
	title       *string
	best_before *time.Time
	use_by      *time.Time
	created_at  *time.Time
	updated_at  *time.Time
	recipes     map[int]struct{}
}

// SetUID sets the uid field.
func (ic *IngredientCreate) SetUID(s string) *IngredientCreate {
	ic.uid = &s
	return ic
}

// SetTitle sets the title field.
func (ic *IngredientCreate) SetTitle(s string) *IngredientCreate {
	ic.title = &s
	return ic
}

// SetBestBefore sets the best_before field.
func (ic *IngredientCreate) SetBestBefore(t time.Time) *IngredientCreate {
	ic.best_before = &t
	return ic
}

// SetUseBy sets the use_by field.
func (ic *IngredientCreate) SetUseBy(t time.Time) *IngredientCreate {
	ic.use_by = &t
	return ic
}

// SetNillableUseBy sets the use_by field if the given value is not nil.
func (ic *IngredientCreate) SetNillableUseBy(t *time.Time) *IngredientCreate {
	if t != nil {
		ic.SetUseBy(*t)
	}
	return ic
}

// SetCreatedAt sets the created_at field.
func (ic *IngredientCreate) SetCreatedAt(t time.Time) *IngredientCreate {
	ic.created_at = &t
	return ic
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (ic *IngredientCreate) SetNillableCreatedAt(t *time.Time) *IngredientCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the updated_at field.
func (ic *IngredientCreate) SetUpdatedAt(t time.Time) *IngredientCreate {
	ic.updated_at = &t
	return ic
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (ic *IngredientCreate) SetNillableUpdatedAt(t *time.Time) *IngredientCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// AddRecipeIDs adds the recipes edge to Recipe by ids.
func (ic *IngredientCreate) AddRecipeIDs(ids ...int) *IngredientCreate {
	if ic.recipes == nil {
		ic.recipes = make(map[int]struct{})
	}
	for i := range ids {
		ic.recipes[ids[i]] = struct{}{}
	}
	return ic
}

// AddRecipes adds the recipes edges to Recipe.
func (ic *IngredientCreate) AddRecipes(r ...*Recipe) *IngredientCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ic.AddRecipeIDs(ids...)
}

// Save creates the Ingredient in the database.
func (ic *IngredientCreate) Save(ctx context.Context) (*Ingredient, error) {
	if ic.uid == nil {
		return nil, errors.New("ent: missing required field \"uid\"")
	}
	if err := ingredient.UIDValidator(*ic.uid); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"uid\": %v", err)
	}
	if ic.title == nil {
		return nil, errors.New("ent: missing required field \"title\"")
	}
	if ic.best_before == nil {
		return nil, errors.New("ent: missing required field \"best_before\"")
	}
	if ic.created_at == nil {
		v := ingredient.DefaultCreatedAt()
		ic.created_at = &v
	}
	if ic.updated_at == nil {
		v := ingredient.DefaultUpdatedAt()
		ic.updated_at = &v
	}
	return ic.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IngredientCreate) SaveX(ctx context.Context) *Ingredient {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ic *IngredientCreate) sqlSave(ctx context.Context) (*Ingredient, error) {
	var (
		i     = &Ingredient{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ingredient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ingredient.FieldID,
			},
		}
	)
	if value := ic.uid; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: ingredient.FieldUID,
		})
		i.UID = *value
	}
	if value := ic.title; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: ingredient.FieldTitle,
		})
		i.Title = *value
	}
	if value := ic.best_before; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: ingredient.FieldBestBefore,
		})
		i.BestBefore = *value
	}
	if value := ic.use_by; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: ingredient.FieldUseBy,
		})
		i.UseBy = value
	}
	if value := ic.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: ingredient.FieldCreatedAt,
		})
		i.CreatedAt = *value
	}
	if value := ic.updated_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: ingredient.FieldUpdatedAt,
		})
		i.UpdatedAt = *value
	}
	if nodes := ic.recipes; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ingredient.RecipesTable,
			Columns: ingredient.RecipesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recipe.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	i.ID = int(id)
	return i, nil
}