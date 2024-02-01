package request

type MealPlanRequest struct {
	Id        string   `json:"id"`
	RecipeIds []string `json:"recipeIds"`
}
