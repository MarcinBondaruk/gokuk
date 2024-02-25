package service

import (
	"github.com/MarcinBondaruk/gokuk/internal/application/api/request"
	"github.com/MarcinBondaruk/gokuk/internal/domain/recipe"
)

type RecipeService struct {
	recipeRepository *recipe.RecipeRepository
}

func (r *RecipeService) CreateRecipe(req *request.RecipeRequest) error {
	// add recipe
	return nil
}

func (r *RecipeService) GetRecipes(req *request.RecipeRequest) error {
	return nil
}

func (r *RecipeService) GetRecipeById(recipeId string) error {
	return nil
}
