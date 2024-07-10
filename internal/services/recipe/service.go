package recipe

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RecipeService struct {
	recipeRepository *repository
}

type CreateRecipeCmd struct {
	AuthorID    string
	Title       string
	Description string
	Ingredients []CmdIngredient
}

type CmdIngredient struct {
	Name     string
	Quantity int
	Unit     string
}

func NewRecipeService(db *pgxpool.Pool) *RecipeService {
	return &RecipeService{
		recipeRepository: newRecipeRepository(db),
	}
}

func (rs *RecipeService) CreateRecipe(cmd CreateRecipeCmd) error {
	recipeID, err := uuid.NewV7()
	if err != nil {
		return err
	}

	ingredients := make([]Ingredient, 0, len(cmd.Ingredients))
	for _, ingredient := range cmd.Ingredients {
		ingredients = append(ingredients, Ingredient(ingredient))
	}

	authorID, _ := uuid.Parse(cmd.AuthorID)

	recipe := NewRecipe(recipeID, authorID, cmd.Title, cmd.Description, ingredients)

	return rs.recipeRepository.Add(recipe)
}

func (rs *RecipeService) GetRecipes() ([]RecipeView, error) {
	return rs.recipeRepository.GetRecipes()
}
