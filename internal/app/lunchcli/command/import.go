package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

type Ingredient struct {
	Title		string	`json:"title,omitempty"`
	BestBefore	string	`json:"best-before,omitempty"`
	UseBy		string	`json:"use-by,omitempty"`
}

type Ingredients struct {
	Ingredients	[]*Ingredient  `json:"ingredients,omitempty"`
}

type ListIngredients struct {
	Ingredients	Ingredients
}

// NewImportCommand creates a new cobra.Command for importing recipe and ingredients data.
func NewImportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "import",
		Aliases: []string{"i"},
		Short:   "Import Data",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceErrors = true
			cmd.SilenceUsage = true

			return runImport()
		},
	}

	return cmd
}

func runImport() error {
	fmt.Println(os.Getwd())
	ingredientPath := "etc/fixtures/ingredients.json"
	recipePath := "etc/fixtures/recipes.json"
	fmt.Printf("Import from  %q \n", ingredientPath)
	ingredientsFile, err := os.Open(ingredientPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Successfully Opened ingredients fixtures")
	b, err := ioutil.ReadAll(ingredientsFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var listIngredients ListIngredients
	json.Unmarshal(b, &listIngredients.Ingredients)

	var ingredients []*Ingredient
	ingredients = listIngredients.Ingredients.Ingredients
	fmt.Println(ingredients)
	fmt.Println("Successfully Get ingredients fixtures")

	for _, ingredient := range ingredients {
		fmt.Println(ingredient)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer ingredientsFile.Close()


	fmt.Printf("Import from  %q ", recipePath)

	return nil
}
