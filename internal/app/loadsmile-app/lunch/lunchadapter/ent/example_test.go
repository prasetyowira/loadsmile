// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleIngredient() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the ingredient's edges.

	// create ingredient vertex with its edges.
	i := client.Ingredient.
		Create().
		SetUID("string").
		SetTitle("string").
		SetBestBefore(time.Now()).
		SetUseBy(time.Now()).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SaveX(ctx)
	log.Println("ingredient created:", i)

	// query edges.

	// Output:
}
func ExampleRecipe() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the recipe's edges.
	i0 := client.Ingredient.
		Create().
		SetUID("string").
		SetTitle("string").
		SetBestBefore(time.Now()).
		SetUseBy(time.Now()).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SaveX(ctx)
	log.Println("ingredient created:", i0)

	// create recipe vertex with its edges.
	r := client.Recipe.
		Create().
		SetUID("string").
		SetTitle("string").
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		AddIngredients(i0).
		SaveX(ctx)
	log.Println("recipe created:", r)

	// query edges.
	i0, err = r.QueryIngredients().First(ctx)
	if err != nil {
		log.Fatalf("failed querying ingredients: %v", err)
	}
	log.Println("ingredients found:", i0)

	// Output:
}
