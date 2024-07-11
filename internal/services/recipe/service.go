package recipe

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RecipeService struct {
	recipeRepository *repository
}

type CreateRecipeCmd struct {
	AuthorID    int64
	Title       string
	Description string
	Ingredients []CmdIngredient
}

type CmdIngredient struct {
	Name  string
	Value int
	Unit  string
}

func NewRecipeService(db *pgxpool.Pool) *RecipeService {
	return &RecipeService{
		recipeRepository: newRecipeRepository(db),
	}
}

func (rs *RecipeService) CreateRecipe(ctx context.Context, cmd CreateRecipeCmd) error {
	ingredients := make([]Ingredient, 0, len(cmd.Ingredients))
	for _, ingredient := range cmd.Ingredients {
		ingredients = append(ingredients, Ingredient(ingredient))
	}

	recipe := NewRecipe(cmd.AuthorID, cmd.Title, cmd.Description, ingredients)

	return rs.recipeRepository.Add(ctx, recipe)
}

func (rs *RecipeService) GetRecipes(ctx context.Context) ([]RecipeView, error) {
	return rs.recipeRepository.GetRecipes(ctx)
}
