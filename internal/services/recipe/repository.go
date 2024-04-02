package recipe

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRecipeRepo struct {
	Connection *pgxpool.Pool
}

func (rr PostgresRecipeRepo) Add(r *recipe) error {
	_, err := rr.Connection.Exec(context.Background(), "INSERT INTO recipes (id, author_id, title, description) VALUES ($1, $2, $3, $4)", r.id, r.authorId, r.title, r.description)

	if err != nil {
		return err
	}

	return nil
}

func (rr PostgresRecipeRepo) Retrieve(id string) (*recipe, error) {
	var recipeId string
	var authorId string
	var title string
	var description string

	// need to query for reciepies with ingredients
	sql := "SELECT id, author_id, title, description FROM recipes JOIN recipe_ingredient ON recipes.id = recipe_ingredient.recipe_id WHERE id = $1"
	err := rr.Connection.QueryRow(context.Background(), sql, id).Scan(&recipeId, &authorId, &title, &description)

	if err != nil {
		return nil, err
	}

	recipeIdUUID, err := uuid.Parse(recipeId)
	if err != nil {
		return nil, err
	}

	authorIdUUID, err := uuid.Parse(authorId)
	if err != nil {
		return nil, err
	}

	recipe := NewRecipe(recipeIdUUID, authorIdUUID, title, description, nil)
	return recipe, nil
}
