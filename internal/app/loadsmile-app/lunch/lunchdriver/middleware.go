package lunchdriver

import (
	"context"

	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent"
)

// Middleware describes a service middleware.
type Middleware func(lunch.Service) lunch.Service

// LoggingMiddleware is a service level logging middleware for TodoList.
func LoggingMiddleware(logger lunch.Logger) Middleware {
	return func(next lunch.Service) lunch.Service {
		return loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   lunch.Service
	logger lunch.Logger
}

func (mw loggingMiddleware) ListRecipes(ctx context.Context, args interface{}) ([]*ent.Recipe, error) {
	logger := mw.logger.WithContext(ctx)
	logger.Info("listing recipes")
	return mw.next.ListRecipes(ctx, args)
}

func (mw loggingMiddleware) GetRecipe(ctx context.Context, id string) (*ent.Recipe, error) {
	logger := mw.logger.WithContext(ctx)
	logger.Info("get recipe", map[string]interface{}{
		"id": id,
	})

	return mw.next.GetRecipe(ctx, id)
}

func (mw loggingMiddleware) GetLunch(ctx context.Context) ([]*ent.Recipe, error) {
	logger := mw.logger.WithContext(ctx)
	logger.Info("get lunch")

	return mw.next.GetLunch(ctx)
}

