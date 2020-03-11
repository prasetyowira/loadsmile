// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package tododriver

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"

	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent"
)

// endpointError identifies an error that should be returned as an endpoint error.
type endpointError interface {
	EndpointError() bool
}

// serviceError identifies an error that should be returned as a service error.
type serviceError interface {
	ServiceError() bool
}

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateTodo 		endpoint.Endpoint
	ListTodos  		endpoint.Endpoint
	MarkAsDone 		endpoint.Endpoint
	ListBuildings 	endpoint.Endpoint
}

// MakeEndpoints returns a(n) Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service todo.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Combine(middleware...)

	return Endpoints{
		CreateTodo: kitxendpoint.OperationNameMiddleware("todo.CreateTodo")(mw(MakeCreateTodoEndpoint(service))),
		ListTodos:  kitxendpoint.OperationNameMiddleware("todo.ListTodos")(mw(MakeListTodosEndpoint(service))),
		MarkAsDone: kitxendpoint.OperationNameMiddleware("todo.MarkAsDone")(mw(MakeMarkAsDoneEndpoint(service))),
	}
}

// CreateTodoRequest is a request struct for CreateTodo endpoint.
type CreateTodoRequest struct {
	Text string
}

// CreateTodoResponse is a response struct for CreateTodo endpoint.
type CreateTodoResponse struct {
	Id  string
	Err error
}

func (r CreateTodoResponse) Failed() error {
	return r.Err
}

// MakeCreateTodoEndpoint returns an endpoint for the matching method of the underlying service.
func MakeCreateTodoEndpoint(service todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTodoRequest)

		id, err := service.CreateTodo(ctx, req.Text)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return CreateTodoResponse{
					Err: err,
					Id:  id,
				}, nil
			}

			return CreateTodoResponse{
				Err: err,
				Id:  id,
			}, err
		}

		return CreateTodoResponse{Id: id}, nil
	}
}

// ListTodosRequest is a request struct for ListTodos endpoint.
type ListTodosRequest struct{}

// ListTodosResponse is a response struct for ListTodos endpoint.
type ListTodosResponse struct {
	Todos []todo.Todo
	Err   error
}

func (r ListTodosResponse) Failed() error {
	return r.Err
}

// MakeListTodosEndpoint returns an endpoint for the matching method of the underlying service.
func MakeListTodosEndpoint(service todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		todos, err := service.ListTodos(ctx)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ListTodosResponse{
					Err:   err,
					Todos: todos,
				}, nil
			}

			return ListTodosResponse{
				Err:   err,
				Todos: todos,
			}, err
		}

		return ListTodosResponse{Todos: todos}, nil
	}
}

// MarkAsDoneRequest is a request struct for MarkAsDone endpoint.
type MarkAsDoneRequest struct {
	Id string
}

// MarkAsDoneResponse is a response struct for MarkAsDone endpoint.
type MarkAsDoneResponse struct {
	Err error
}

func (r MarkAsDoneResponse) Failed() error {
	return r.Err
}

// MakeMarkAsDoneEndpoint returns an endpoint for the matching method of the underlying service.
func MakeMarkAsDoneEndpoint(service todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MarkAsDoneRequest)

		err := service.MarkAsDone(ctx, req.Id)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return MarkAsDoneResponse{Err: err}, nil
			}

			return MarkAsDoneResponse{Err: err}, err
		}

		return MarkAsDoneResponse{}, nil
	}
}


// ListTodosRequest is a request struct for ListTodos endpoint.
type ListRecipesRequest struct{
	Limit	string
	Offset	string
	Search 	string
}

// ListTodosResponse is a response struct for ListTodos endpoint.
type ListRecipesResponse struct {
	Recipes 	[]ent.Recipe
	Err			error
}

func (r ListRecipesResponse) Failed() error {
	return r.Err
}

// MakeListTodosEndpoint returns an endpoint for the matching method of the underlying service.
func MakeListRecipesEndpoint(service lunch.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		recipes, err := service.ListRecipes(ctx)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ListRecipesResponse{
					Err:   err,
					Recipes: recipes,
				}, nil
			}

			return ListRecipesResponse{
				Err:   err,
				Recipes: recipes,
			}, err
		}

		return ListRecipesResponse{Recipes: recipes}, nil
	}
}
