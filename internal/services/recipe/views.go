package recipe

import "github.com/google/uuid"

type RecipeView struct {
	ID          uuid.UUID        `json:"id"`
	AuthorID    string           `json:"author_id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Ingredients []IngredientView `json:"ingredients"`
}

type IngredientView struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}
