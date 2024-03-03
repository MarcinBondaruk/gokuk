package recipe

import (
	"github.com/google/uuid"
)

type Recipe struct {
	id          uuid.UUID
	authorId    uuid.UUID
	title       string
	description string
	ingredients []Ingredient
}

func NewRecipe(id, authorId uuid.UUID, title, description string, ingredients []Ingredient) *Recipe {
	return &Recipe{
		id:          id,
		authorId:    authorId,
		title:       title,
		description: description,
		ingredients: ingredients,
	}
}

func (r *Recipe) Id() string {
	return r.id.String()
}

func (r *Recipe) AuthorId() string {
	return r.authorId.String()
}

func (r *Recipe) Title() string {
	return r.title
}

func (r *Recipe) Description() string {
	return r.description
}

func (r *Recipe) Ingredients() []Ingredient {
	return r.ingredients
}
