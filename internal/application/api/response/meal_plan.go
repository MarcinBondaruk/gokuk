package response

type MealPlanResponse struct {
	Id        string   `json:"id"`
	RecipeIds []string `json:"recipeIds"`
}
