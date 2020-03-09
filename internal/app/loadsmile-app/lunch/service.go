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
	// CreateTodo adds a new todo to the todo list.
	// CreateTodo(ctx context.Context, text string) (id string, err error)

	// ListTodos returns the list of todos.
	// ListTodos(ctx context.Context) (todos []Todo, err error)

	// MarkAsDone marks a todo as done.
	// MarkAsDone(ctx context.Context, id string) error

	// ListBuildings returns the list of buildings.
	ListRecipes(ctx context.Context) (recipes []*ent.Recipe, err error)
	GetLunch(ctx context.Context) ([]*ent.Recipe, error)
	GetRecipe(ctx context.Context, id string) (*ent.Recipe, error)
}

type service struct {
	idgenerator IDGenerator
	store       Store
}

// IDGenerator generates a new ID.
type IDGenerator interface {
	// Generate generates a new ID.
	Generate() (string, error)
}

// Store provides todo persistence.
type Store interface {
	// // Store stores a todo.
	// Store(ctx context.Context, todo Todo) error

	// All returns all todos.
	All(ctx context.Context) ([]*ent.Recipe, error)

	// Get returns a single todo by its ID.
	Get(ctx context.Context, id string) (*ent.Recipe, error)

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
func NewService(idgenerator IDGenerator, store Store) Service {
	return &service{
		idgenerator: idgenerator,
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

func (s service) ListRecipes(ctx context.Context) ([]*ent.Recipe, error) {
	return s.store.All(ctx)
}

func (s service) GetLunch(ctx context.Context) ([]*ent.Recipe, error) {
	return s.store.GetLunch(ctx)
}

func (s service) GetRecipe(ctx context.Context, id string) (*ent.Recipe, error) {
	recipe, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to mark todo as done")
	}
	return recipe, nil
}
