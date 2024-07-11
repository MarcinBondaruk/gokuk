package recipe

type RecipeView struct {
	ID          int64            `json:"id"`
	AuthorID    int64            `json:"author_id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Ingredients []IngredientView `json:"ingredients"`
}

type IngredientView struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}
