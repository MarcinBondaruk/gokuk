package service

import "github.com/MarcinBondaruk/gokuk/internal/application/api/request"

type RecipeService struct {
}

func (r RecipeService) AddRecipe(req *request.RecipeRequest) error {
	// add recipe
	return nil
}

func (r RecipeService) GetRecipes(req *request.RecipeRequest) error {
	return nil
}

func (r RecipeService) GetRecipeById(recipeId string) error {
	return nil
}
