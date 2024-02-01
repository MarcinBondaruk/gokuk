package response

type RecipeResponse struct {
	Description string               `json:"description"`
	Id          string               `json:"id"`
	Ingredients []IngredientResponse `json:"ingredients"`
	Title       string               `json:"title"`
}
