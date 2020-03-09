// Code generated by entc, DO NOT EDIT.

package recipe

import (
	"time"

	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/schema"
)

const (
	// Label holds the string label denoting the recipe type in the database.
	Label = "recipe"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid vertex property in the database.
	FieldUID = "uid"
	// FieldTitle holds the string denoting the title vertex property in the database.
	FieldTitle = "title"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"

	// Table holds the table name of the recipe in the database.
	Table = "recipes"
	// IngredientsTable is the table the holds the ingredients relation/edge. The primary key declared below.
	IngredientsTable = "recipe_ingredients"
	// IngredientsInverseTable is the table name for the Ingredient entity.
	// It exists in this package in order to avoid circular dependency with the "ingredient" package.
	IngredientsInverseTable = "ingredients"
)

// Columns holds all SQL columns for recipe fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldTitle,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// IngredientsPrimaryKey and IngredientsColumn2 are the table columns denoting the
	// primary key for the ingredients relation (M2M).
	IngredientsPrimaryKey = []string{"recipe_id", "ingredient_id"}
)

var (
	fields = schema.Recipe{}.Fields()

	// descUID is the schema descriptor for uid field.
	descUID = fields[0].Descriptor()
	// UIDValidator is a validator for the "uid" field. It is called by the builders before save.
	UIDValidator = func() func(string) error {
		validators := descUID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(uid string) error {
			for _, fn := range fns {
				if err := fn(uid); err != nil {
					return err
				}
			}
			return nil
		}
	}()

	// descCreatedAt is the schema descriptor for created_at field.
	descCreatedAt = fields[2].Descriptor()
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt = descCreatedAt.Default.(func() time.Time)

	// descUpdatedAt is the schema descriptor for updated_at field.
	descUpdatedAt = fields[3].Descriptor()
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt = descUpdatedAt.Default.(func() time.Time)
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt = descUpdatedAt.UpdateDefault.(func() time.Time)
)
