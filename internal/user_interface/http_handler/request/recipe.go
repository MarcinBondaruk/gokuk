package request

import (
	"errors"

	"github.com/MarcinBondaruk/gokuk/internal/services/recipe"
)

const (
	ErrDescriptionEmpty          = "description cannot be empty"
	ErrIngredientsEmpty          = "ingredients cannot be empty"
	ErrIngredientNameEmpty       = "ingredient name cannot be empty"
	ErrIngredientQuantityInvalid = "ingredient quantity must be greater than 0"
	ErrIngredientUnitEmpty       = "ingredient unit cannot be empty"
	ErrTitleEmpty                = "title cannot be empty"
)

type CreateRecipeRequest struct {
	AuthorID    string       `json:"author_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

func (crr *CreateRecipeRequest) Validate() error {
	if crr.Title == "" {
		return errors.New(ErrTitleEmpty)
	}

	if crr.Description == "" {
		return errors.New(ErrDescriptionEmpty)
	}

	if len(crr.Ingredients) == 0 {
		return errors.New(ErrIngredientsEmpty)
	}

	for _, ingredient := range crr.Ingredients {
		if ingredient.Name == "" {
			return errors.New(ErrIngredientNameEmpty)
		}

		if ingredient.Quantity <= 0 {
			return errors.New(ErrIngredientQuantityInvalid)
		}

		if ingredient.Unit == "" {
			return errors.New(ErrIngredientUnitEmpty)
		}
	}

	return nil
}

func (crr *CreateRecipeRequest) ToCommand() recipe.CreateRecipeCmd {
	ingredients := make([]recipe.CmdIngredient, 0, len(crr.Ingredients))

	for _, ingredient := range crr.Ingredients {
		ingredients = append(ingredients, recipe.CmdIngredient{
			Name:     ingredient.Name,
			Quantity: ingredient.Quantity,
			Unit:     ingredient.Unit,
		})
	}

	return recipe.CreateRecipeCmd{
		AuthorID:    crr.AuthorID,
		Title:       crr.Title,
		Description: crr.Description,
		Ingredients: ingredients,
	}
}
