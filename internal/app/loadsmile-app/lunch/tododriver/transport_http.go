package tododriver

import (
	"context"
	"encoding/json"
	"net/http"

	"emperror.dev/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	appkithttp "github.com/sagikazarmark/appkit/transport/http"
	kitxhttp "github.com/sagikazarmark/kitx/transport/http"

	api "github.com/prasetyowira/loadsmile/.gen/api/openapi/lunch/go"
)

// RegisterHTTPHandlers mounts all of the service endpoints into a router.
func RegisterHTTPHandlers(endpoints Endpoints, router *mux.Router, options ...kithttp.ServerOption) {
	errorEncoder := kitxhttp.NewJSONProblemErrorResponseEncoder(appkithttp.NewDefaultProblemConverter())

	// router.Methods(http.MethodPost).Path("").Handler(kithttp.NewServer(
	// 	endpoints.CreateTodo,
	// 	decodeCreateTodoHTTPRequest,
	// 	kitxhttp.ErrorResponseEncoder(encodeCreateTodoHTTPResponse, errorEncoder),
	// 	options...,
	// ))

	router.Methods(http.MethodGet).Path("recipe").Handler(kithttp.NewServer(
		endpoints.ListBuildings,
		kithttp.NopRequestDecoder,
		kitxhttp.ErrorResponseEncoder(encodeListRecipesHTTPResponse, errorEncoder),
		options...,
	))

	// router.Methods(http.MethodPost).Path("/{id}/done").Handler(kithttp.NewServer(
	// 	endpoints.MarkAsDone,
	// 	decodeMarkAsDoneHTTPRequest,
	// 	kitxhttp.ErrorResponseEncoder(kitxhttp.NopResponseEncoder, errorEncoder),
	// 	options...,
	// ))
}

// func decodeCreateTodoHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var apiRequest api.CreateTodoRequest
//
// 	err := json.NewDecoder(r.Body).Decode(&apiRequest)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to decode request")
// 	}
//
// 	return CreateTodoRequest{
// 		Text: apiRequest.Text,
// 	}, nil
// }

func decodeListBuildingsHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var apiRequest api.CreateTodoRequest

	err := json.NewDecoder(r.Body).Decode(&apiRequest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request")
	}

	return CreateTodoRequest{
		Text: apiRequest.Text,
	}, nil
}

// func encodeCreateTodoHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
// 	resp := response.(CreateTodoResponse)
//
// 	apiResponse := api.CreateTodoResponse{
// 		Id: resp.Id,
// 	}
//
// 	return kitxhttp.JSONResponseEncoder(ctx, w, kitxhttp.WithStatusCode(apiResponse, http.StatusCreated))
// }

func encodeListRecipesHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(ListRecipesResponse)

	apiResponse := api.RecipeList{}

	for _, recipe := range resp.Recipes {
		var listIngredients []api.Ingredient
		for _, ingredient := range recipe.Edges.Ingredients {
			listIngredients = append(listIngredients, api.Ingredient{
				Id:    ingredient.UID,
				Title: ingredient.Title,
			})
		}
		apiResponse.Recipes = append(apiResponse.Recipes, api.Recipe{
			Id:   			string(recipe.ID),
			Title:			recipe.Title,
			Ingredients: 	listIngredients,
		})
	}

	return kitxhttp.JSONResponseEncoder(ctx, w, apiResponse)
}

// func decodeMarkAsDoneHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	vars := mux.Vars(r)
//
// 	id, ok := vars["id"]
// 	if !ok || id == "" {
// 		return nil, errors.NewWithDetails("missing parameter from the URL", "param", "id")
// 	}
//
// 	return MarkAsDoneRequest{
// 		Id: id,
// 	}, nil
// }
