package recipe

import (
	"github.com/google/uuid"
)

type recipe struct {
	id          uuid.UUID
	authorId    uuid.UUID
	title       string
	description string
	ingredients []Ingredient
}

func NewRecipe(id, authorId uuid.UUID, title, description string, ingredients []Ingredient) *recipe {
	return &recipe{
		id:          id,
		authorId:    authorId,
		title:       title,
		description: description,
		ingredients: ingredients,
	}
}
