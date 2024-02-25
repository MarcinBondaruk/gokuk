package repository

import "github.com/MarcinBondaruk/gokuk/internal/domain/recipe"

type PostgresRecipeRepo struct {
}

func (rr PostgresRecipeRepo) Add(r *recipe.Recipe) error {
	return nil
}

func (rr PostgresRecipeRepo) Retrieve(id string) (*recipe.Recipe, error) {
	return nil, nil
}
