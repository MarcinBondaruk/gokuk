package request

type RecipeRequest struct {
	Description string              `json:"description"`
	Ingredients []IngredientRequest `json:"ingredients"`
	Title       string              `json:"title"`
}
