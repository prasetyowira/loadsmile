package lunch

import (
	"context"
	// "time"

	"emperror.dev/errors"

	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent"
)

// Todo is a note describing a task to be done.
// type Ingredient struct {
// 	ID   			string
// 	Title 			string
// 	BestBefore 		time.Time
// 	UseBy			time.Time
// }
//
// type Recipe struct {
// 	ID   			string
// 	Title 			string
// 	Ingredients 	[]*Ingredient
// }


// +kit:endpoint:errorStrategy=service

// Service manages a list of todos.
type Service interface {
	ListRecipes(ctx context.Context, args interface{}) (recipes []*ent.Recipe, err error)
	GetLunch(ctx context.Context) ([]*ent.Recipe, error)
	GetRecipe(ctx context.Context, id string) (*ent.Recipe, error)
}

type service struct {
	store       Store
}

// Store provides todo persistence.
type Store interface {
	// // Store stores a todo.
	// Store(ctx context.Context, todo Todo) error

	// All returns all todos.
	AllRecipes(ctx context.Context, args interface{}) ([]*ent.Recipe, error)

	// Get returns a single todo by its ID.
	GetRecipe(ctx context.Context, id string) (*ent.Recipe, error)

	GetLunch(ctx context.Context) ([]*ent.Recipe, error)
}

// NotFoundError is returned if a todo cannot be found.
type NotFoundError struct {
	ID string
}

// Error implements the error interface.
func (NotFoundError) Error() string {
	return "recipe not found"
}

// Details returns error details.
func (e NotFoundError) Details() []interface{} {
	return []interface{}{"recipe_id", e.ID}
}

// NotFound tells a client that this error is related to a resource being not found.
// Can be used to translate the error to eg. status code.
func (NotFoundError) NotFound() bool {
	return true
}

// ServiceError tells the transport layer whether this error should be translated into the transport format
// or an internal error should be returned instead.
func (NotFoundError) ServiceError() bool {
	return true
}

// NewService returns a new Service.
func NewService(store Store) Service {
	return &service{
		store:       store,
	}
}

type validationError struct {
	violations map[string][]string
}

func (validationError) Error() string {
	return "invalid recipes"
}

func (e validationError) Violations() map[string][]string {
	return e.violations
}

// Validation tells a client that this error is related to a resource being invalid.
// Can be used to translate the error to eg. status code.
func (validationError) Validation() bool {
	return true
}

// ServiceError tells the transport layer whether this error should be translated into the transport format
// or an internal error should be returned instead.
func (validationError) ServiceError() bool {
	return true
}

func (s service) ListRecipes(ctx context.Context, args interface{}) ([]*ent.Recipe, error) {
	return s.store.AllRecipes(ctx, args)
}

func (s service) GetLunch(ctx context.Context) ([]*ent.Recipe, error) {
	return s.store.GetLunch(ctx)
}

func (s service) GetRecipe(ctx context.Context, id string) (*ent.Recipe, error) {
	recipe, err := s.store.GetRecipe(ctx, id)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to mark todo as done")
	}
	return recipe, nil
}
