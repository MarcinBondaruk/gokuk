package cooking_book

type RecipeRepository interface {
	Add(Recipe) error
	Retrieve(id string) (Recipe, error)
}
