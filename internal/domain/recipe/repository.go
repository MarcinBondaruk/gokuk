package recipe

type RecipeRepository interface {
	Add(recipe *Recipe) error
	Retrieve(id string) (*Recipe, error)
}
