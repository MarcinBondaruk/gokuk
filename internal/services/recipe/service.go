package recipe

type RecipeService struct {
	recipeRepository RecipeRepository
}

func NewRecipeService(recipeRepository RecipeRepository) RecipeService {
	return RecipeService{
		recipeRepository: recipeRepository,
	}
}

func (rs *RecipeService) GetRecipes() ([]*RecipeView, error) {
	return rs.recipeRepository.GetRecipes()
}
