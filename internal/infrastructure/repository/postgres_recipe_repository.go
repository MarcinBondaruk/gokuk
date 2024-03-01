package repository

import (
	"fmt"

	"github.com/MarcinBondaruk/gokuk/internal/domain/recipe"
)

type PostgresRecipeRepo struct {
}

func (rr PostgresRecipeRepo) Add(r *recipe.Recipe) error {
	fmt.Println("ADDING\n", r)
	return nil
}

func (rr PostgresRecipeRepo) Retrieve(id string) (*recipe.Recipe, error) {
	fmt.Println("looking for recipe with id: ", id)
	return nil, nil
}
