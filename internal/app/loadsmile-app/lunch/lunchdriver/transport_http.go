package lunchdriver

import (
	"context"
	"fmt"
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

	router.Methods(http.MethodGet).Path("/lunch").Handler(kithttp.NewServer(
		endpoints.GetLunch,
		kithttp.NopRequestDecoder,
		kitxhttp.ErrorResponseEncoder(encodeListRecipesHTTPResponse, errorEncoder),
		options...,
	))

	router.Methods(http.MethodGet).Path("/recipe").Handler(kithttp.NewServer(
		endpoints.ListRecipes,
		decodeListRecipeHTTPRequest,
		kitxhttp.ErrorResponseEncoder(encodeListRecipesHTTPResponse, errorEncoder),
		options...,
	))

	router.Methods(http.MethodGet).Path("/recipe/{id}").Handler(kithttp.NewServer(
		endpoints.GetRecipe,
		decodeGetRecipeHTTPRequest,
		kitxhttp.ErrorResponseEncoder(encodeGetLunchHTTPResponse, errorEncoder),
		options...,
	))
}

func decodeListRecipeHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	query := r.URL.Query()
	fmt.Println(query)

	limit := query.Get("limit")
	offset := query.Get("offset")
	search := query.Get("search")
	fmt.Println(limit)


	return ListRecipesRequest{
		Limit: limit,
		Offset: offset,
		Search: search,
	}, nil
}

func encodeListRecipesHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(ListRecipesResponse)

	apiResponse := api.RecipeList{}

	for _, recipe := range resp.Recipes {
		var listIngredients []api.Ingredient
		for _, ingredient := range recipe.Edges.Ingredients {
			listIngredients = append(listIngredients, api.Ingredient{
				Id:    ingredient.ID,
				Title: ingredient.Title,
			})
		}
		apiResponse.Recipes = append(apiResponse.Recipes, api.Recipe{
			Id:   			recipe.ID,
			Title:			recipe.Title,
			Ingredients: 	listIngredients,
		})
	}

	return kitxhttp.JSONResponseEncoder(ctx, w, apiResponse)
}

func decodeGetRecipeHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok || id == "" {
		return nil, errors.NewWithDetails("missing parameter from the URL", "param", "id")
	}

	return GetRecipeRequest{
		Id: id,
	}, nil
}

func encodeGetRecipeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetRecipeResponse)

	var listIngredients []api.Ingredient
	for _, ingredient := range resp.Recipe.Edges.Ingredients {
		listIngredients = append(listIngredients, api.Ingredient{
			Id:    ingredient.ID,
			Title: ingredient.Title,
		})
	}

	apiResponse := api.Recipe{
		Id: resp.Recipe.ID,
		Title: resp.Recipe.Title,
		Ingredients: listIngredients,
	}

	return kitxhttp.JSONResponseEncoder(ctx, w, apiResponse)
}

func encodeGetLunchHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetLunchResponse)

	apiResponse := api.RecipeList{}

	for _, recipe := range resp.Recipes {
		var listIngredients []api.Ingredient
		for _, ingredient := range recipe.Edges.Ingredients {
			listIngredients = append(listIngredients, api.Ingredient{
				Id:    ingredient.ID,
				Title: ingredient.Title,
			})
		}
		apiResponse.Recipes = append(apiResponse.Recipes, api.Recipe{
			Id:   			recipe.ID,
			Title:			recipe.Title,
			Ingredients: 	listIngredients,
		})
	}

	return kitxhttp.JSONResponseEncoder(ctx, w, apiResponse)
}
