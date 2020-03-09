// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/ingredient"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/predicate"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/recipe"
)

// RecipeUpdate is the builder for updating Recipe entities.
type RecipeUpdate struct {
	config

	title              *string
	created_at         *time.Time
	updated_at         *time.Time
	ingredients        map[int]struct{}
	removedIngredients map[int]struct{}
	predicates         []predicate.Recipe
}

// Where adds a new predicate for the builder.
func (ru *RecipeUpdate) Where(ps ...predicate.Recipe) *RecipeUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetTitle sets the title field.
func (ru *RecipeUpdate) SetTitle(s string) *RecipeUpdate {
	ru.title = &s
	return ru
}

// SetCreatedAt sets the created_at field.
func (ru *RecipeUpdate) SetCreatedAt(t time.Time) *RecipeUpdate {
	ru.created_at = &t
	return ru
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (ru *RecipeUpdate) SetNillableCreatedAt(t *time.Time) *RecipeUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetUpdatedAt sets the updated_at field.
func (ru *RecipeUpdate) SetUpdatedAt(t time.Time) *RecipeUpdate {
	ru.updated_at = &t
	return ru
}

// AddIngredientIDs adds the ingredients edge to Ingredient by ids.
func (ru *RecipeUpdate) AddIngredientIDs(ids ...int) *RecipeUpdate {
	if ru.ingredients == nil {
		ru.ingredients = make(map[int]struct{})
	}
	for i := range ids {
		ru.ingredients[ids[i]] = struct{}{}
	}
	return ru
}

// AddIngredients adds the ingredients edges to Ingredient.
func (ru *RecipeUpdate) AddIngredients(i ...*Ingredient) *RecipeUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ru.AddIngredientIDs(ids...)
}

// RemoveIngredientIDs removes the ingredients edge to Ingredient by ids.
func (ru *RecipeUpdate) RemoveIngredientIDs(ids ...int) *RecipeUpdate {
	if ru.removedIngredients == nil {
		ru.removedIngredients = make(map[int]struct{})
	}
	for i := range ids {
		ru.removedIngredients[ids[i]] = struct{}{}
	}
	return ru
}

// RemoveIngredients removes ingredients edges to Ingredient.
func (ru *RecipeUpdate) RemoveIngredients(i ...*Ingredient) *RecipeUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ru.RemoveIngredientIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RecipeUpdate) Save(ctx context.Context) (int, error) {
	if ru.updated_at == nil {
		v := recipe.UpdateDefaultUpdatedAt()
		ru.updated_at = &v
	}
	return ru.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecipeUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecipeUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecipeUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RecipeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recipe.Table,
			Columns: recipe.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recipe.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := ru.title; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: recipe.FieldTitle,
		})
	}
	if value := ru.created_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: recipe.FieldCreatedAt,
		})
	}
	if value := ru.updated_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: recipe.FieldUpdatedAt,
		})
	}
	if nodes := ru.removedIngredients; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.IngredientsTable,
			Columns: recipe.IngredientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ingredient.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.ingredients; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.IngredientsTable,
			Columns: recipe.IngredientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ingredient.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipe.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RecipeUpdateOne is the builder for updating a single Recipe entity.
type RecipeUpdateOne struct {
	config
	id int

	title              *string
	created_at         *time.Time
	updated_at         *time.Time
	ingredients        map[int]struct{}
	removedIngredients map[int]struct{}
}

// SetTitle sets the title field.
func (ruo *RecipeUpdateOne) SetTitle(s string) *RecipeUpdateOne {
	ruo.title = &s
	return ruo
}

// SetCreatedAt sets the created_at field.
func (ruo *RecipeUpdateOne) SetCreatedAt(t time.Time) *RecipeUpdateOne {
	ruo.created_at = &t
	return ruo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (ruo *RecipeUpdateOne) SetNillableCreatedAt(t *time.Time) *RecipeUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetUpdatedAt sets the updated_at field.
func (ruo *RecipeUpdateOne) SetUpdatedAt(t time.Time) *RecipeUpdateOne {
	ruo.updated_at = &t
	return ruo
}

// AddIngredientIDs adds the ingredients edge to Ingredient by ids.
func (ruo *RecipeUpdateOne) AddIngredientIDs(ids ...int) *RecipeUpdateOne {
	if ruo.ingredients == nil {
		ruo.ingredients = make(map[int]struct{})
	}
	for i := range ids {
		ruo.ingredients[ids[i]] = struct{}{}
	}
	return ruo
}

// AddIngredients adds the ingredients edges to Ingredient.
func (ruo *RecipeUpdateOne) AddIngredients(i ...*Ingredient) *RecipeUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ruo.AddIngredientIDs(ids...)
}

// RemoveIngredientIDs removes the ingredients edge to Ingredient by ids.
func (ruo *RecipeUpdateOne) RemoveIngredientIDs(ids ...int) *RecipeUpdateOne {
	if ruo.removedIngredients == nil {
		ruo.removedIngredients = make(map[int]struct{})
	}
	for i := range ids {
		ruo.removedIngredients[ids[i]] = struct{}{}
	}
	return ruo
}

// RemoveIngredients removes ingredients edges to Ingredient.
func (ruo *RecipeUpdateOne) RemoveIngredients(i ...*Ingredient) *RecipeUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ruo.RemoveIngredientIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ruo *RecipeUpdateOne) Save(ctx context.Context) (*Recipe, error) {
	if ruo.updated_at == nil {
		v := recipe.UpdateDefaultUpdatedAt()
		ruo.updated_at = &v
	}
	return ruo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecipeUpdateOne) SaveX(ctx context.Context) *Recipe {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RecipeUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecipeUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RecipeUpdateOne) sqlSave(ctx context.Context) (r *Recipe, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recipe.Table,
			Columns: recipe.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  ruo.id,
				Type:   field.TypeInt,
				Column: recipe.FieldID,
			},
		},
	}
	if value := ruo.title; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: recipe.FieldTitle,
		})
	}
	if value := ruo.created_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: recipe.FieldCreatedAt,
		})
	}
	if value := ruo.updated_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: recipe.FieldUpdatedAt,
		})
	}
	if nodes := ruo.removedIngredients; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.IngredientsTable,
			Columns: recipe.IngredientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ingredient.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.ingredients; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.IngredientsTable,
			Columns: recipe.IngredientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ingredient.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Recipe{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipe.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
