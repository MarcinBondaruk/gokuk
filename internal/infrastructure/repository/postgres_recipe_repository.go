package repository

import "github.com/MarcinBondaruk/gokuk/domain/cooking_book"

type PostgresRecipeRepo struct {
}

func (rr PostgresRecipeRepo) Add(r cooking_book.Recipe) error {
	return nil
}

func (rr PostgresRecipeRepo) Retrieve(id string) (*cooking_book.Recipe, error) {
	return nil, nil
}
