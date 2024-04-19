package recipe

type recipe struct {
	ID          string
	authorID    string
	title       string
	description string
	ingredients []Ingredient
}

func NewRecipe(id, authorId string, title, description string, ingredients []Ingredient) *recipe {
	return &recipe{
		ID:          id,
		authorID:    authorId,
		title:       title,
		description: description,
		ingredients: ingredients,
	}
}
