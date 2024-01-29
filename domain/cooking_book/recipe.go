package cooking_book

import "github.com/google/uuid"

type Recipe struct {
	ID          uuid.UUID
	title       string
	ingredients []Ingredient
}
