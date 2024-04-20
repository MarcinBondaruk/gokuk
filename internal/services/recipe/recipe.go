package recipe

import "github.com/google/uuid"

type recipe struct {
	ID          uuid.UUID
	authorID    uuid.UUID
	title       string
	description string
	ingredients []Ingredient
}

func NewRecipe(id, authorId uuid.UUID, title, description string, ingredients []Ingredient) *recipe {
	return &recipe{
		ID:          id,
		authorID:    authorId,
		title:       title,
		description: description,
		ingredients: ingredients,
	}
}
