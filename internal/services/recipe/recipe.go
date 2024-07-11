package recipe

type recipe struct {
	ID          int64
	authorID    int64
	title       string
	description string
	ingredients []Ingredient
}

func NewRecipe(authorId int64, title, description string, ingredients []Ingredient) *recipe {
	return &recipe{
		authorID:    authorId,
		title:       title,
		description: description,
		ingredients: ingredients,
	}
}
