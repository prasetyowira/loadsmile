// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/migrate"

	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/ingredient"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/recipe"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Ingredient is the client for interacting with the Ingredient builders.
	Ingredient *IngredientClient
	// Recipe is the client for interacting with the Recipe builders.
	Recipe *RecipeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config:     c,
		Schema:     migrate.NewSchema(c.driver),
		Ingredient: NewIngredientClient(c),
		Recipe:     NewRecipeClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config:     cfg,
		Ingredient: NewIngredientClient(cfg),
		Recipe:     NewRecipeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Ingredient.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config:     cfg,
		Schema:     migrate.NewSchema(cfg.driver),
		Ingredient: NewIngredientClient(cfg),
		Recipe:     NewRecipeClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// IngredientClient is a client for the Ingredient schema.
type IngredientClient struct {
	config
}

// NewIngredientClient returns a client for the Ingredient from the given config.
func NewIngredientClient(c config) *IngredientClient {
	return &IngredientClient{config: c}
}

// Create returns a create builder for Ingredient.
func (c *IngredientClient) Create() *IngredientCreate {
	return &IngredientCreate{config: c.config}
}

// Update returns an update builder for Ingredient.
func (c *IngredientClient) Update() *IngredientUpdate {
	return &IngredientUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *IngredientClient) UpdateOne(i *Ingredient) *IngredientUpdateOne {
	return c.UpdateOneID(i.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *IngredientClient) UpdateOneID(id int64) *IngredientUpdateOne {
	return &IngredientUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Ingredient.
func (c *IngredientClient) Delete() *IngredientDelete {
	return &IngredientDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *IngredientClient) DeleteOne(i *Ingredient) *IngredientDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *IngredientClient) DeleteOneID(id int64) *IngredientDeleteOne {
	return &IngredientDeleteOne{c.Delete().Where(ingredient.ID(id))}
}

// Create returns a query builder for Ingredient.
func (c *IngredientClient) Query() *IngredientQuery {
	return &IngredientQuery{config: c.config}
}

// Get returns a Ingredient entity by its id.
func (c *IngredientClient) Get(ctx context.Context, id int64) (*Ingredient, error) {
	return c.Query().Where(ingredient.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IngredientClient) GetX(ctx context.Context, id int64) *Ingredient {
	i, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return i
}

// QueryRecipes queries the recipes edge of a Ingredient.
func (c *IngredientClient) QueryRecipes(i *Ingredient) *RecipeQuery {
	query := &RecipeQuery{config: c.config}
	id := i.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(ingredient.Table, ingredient.FieldID, id),
		sqlgraph.To(recipe.Table, recipe.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ingredient.RecipesTable, ingredient.RecipesPrimaryKey...),
	)
	query.sql = sqlgraph.Neighbors(i.driver.Dialect(), step)

	return query
}

// RecipeClient is a client for the Recipe schema.
type RecipeClient struct {
	config
}

// NewRecipeClient returns a client for the Recipe from the given config.
func NewRecipeClient(c config) *RecipeClient {
	return &RecipeClient{config: c}
}

// Create returns a create builder for Recipe.
func (c *RecipeClient) Create() *RecipeCreate {
	return &RecipeCreate{config: c.config}
}

// Update returns an update builder for Recipe.
func (c *RecipeClient) Update() *RecipeUpdate {
	return &RecipeUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *RecipeClient) UpdateOne(r *Recipe) *RecipeUpdateOne {
	return c.UpdateOneID(r.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *RecipeClient) UpdateOneID(id int64) *RecipeUpdateOne {
	return &RecipeUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Recipe.
func (c *RecipeClient) Delete() *RecipeDelete {
	return &RecipeDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RecipeClient) DeleteOne(r *Recipe) *RecipeDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RecipeClient) DeleteOneID(id int64) *RecipeDeleteOne {
	return &RecipeDeleteOne{c.Delete().Where(recipe.ID(id))}
}

// Create returns a query builder for Recipe.
func (c *RecipeClient) Query() *RecipeQuery {
	return &RecipeQuery{config: c.config}
}

// Get returns a Recipe entity by its id.
func (c *RecipeClient) Get(ctx context.Context, id int64) (*Recipe, error) {
	return c.Query().Where(recipe.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RecipeClient) GetX(ctx context.Context, id int64) *Recipe {
	r, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return r
}

// QueryIngredients queries the ingredients edge of a Recipe.
func (c *RecipeClient) QueryIngredients(r *Recipe) *IngredientQuery {
	query := &IngredientQuery{config: c.config}
	id := r.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(recipe.Table, recipe.FieldID, id),
		sqlgraph.To(ingredient.Table, ingredient.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, recipe.IngredientsTable, recipe.IngredientsPrimaryKey...),
	)
	query.sql = sqlgraph.Neighbors(r.driver.Dialect(), step)

	return query
}
