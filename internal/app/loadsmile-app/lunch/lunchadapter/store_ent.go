package lunchadapter

import (
	"context"
	"fmt"

	"emperror.dev/errors"

	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchdriver"

	// ingredient "github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/ingredient"
	recipeP "github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/recipe"
)

type entStore struct {
	client *ent.Client
}

// NewEntStore returns a new todo store backed by Ent ORM.
func NewEntStore(client *ent.Client) lunch.Store {
	return entStore{
		client: client,
	}
}

// func (s entStore) Store(ctx context.Context, lunch lunch.Recipe) error {
// 	// existing, err := s.client.Todo.Query().Where(todop.UID(todo.ID)).First(ctx)
// 	// if ent.IsNotFound(err) {
// 	// 	_, err := s.client.Todo.Create().
// 	// 		SetUID(todo.ID).
// 	// 		SetText(todo.Text).
// 	// 		SetDone(todo.Done).
// 	// 		Save(ctx)
// 	// 	if err != nil {
// 	// 		return err
// 	// 	}
// 	//
// 	// 	return nil
// 	// }
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	//
// 	// _, err = s.client.Todo.UpdateOneID(existing.ID).
// 	// 	SetText(todo.Text).
// 	// 	SetDone(todo.Done).
// 	// 	Save(ctx)
// 	// if err != nil {
// 	// 	return err
// 	// }
//
// 	return nil
// }

type ListRecipesArgs struct{
	Limit	string
	Offset	string
	Search 	string
}

func (s entStore) AllRecipes(ctx context.Context, args interface{}) ([]*ent.Recipe, error) {
	queryArgs := args.(lunchdriver.ListRecipesRequest)
	fmt.Println(queryArgs)
	recipeDatum, err := s.client.Recipe.Query().WithIngredients().All(ctx)
	if err != nil {
		return nil, err
	}

	return recipeDatum, nil
}

func (s entStore) GetLunch(ctx context.Context) ([]*ent.Recipe, error) {
	recipeDatum, err := s.client.Recipe.Query().WithIngredients().All(ctx)
	if err != nil {
		return nil, err
	}

	return recipeDatum, nil
}

func (s entStore) GetRecipe(ctx context.Context, id string) (*ent.Recipe, error) {
	recipeModel, err := s.client.Recipe.Query().Where(recipeP.UID(id)).First(ctx)
	if ent.IsNotFound(err) {
		return nil, errors.WithStack(lunch.NotFoundError{ID: id})
	}
	return recipeModel, nil
}
