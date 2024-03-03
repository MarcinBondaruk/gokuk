package repository

import (
	"context"

	"github.com/MarcinBondaruk/gokuk/internal/domain/recipe"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PostgresRecipeRepo struct {
	Connection *pgx.Conn
}

func (rr PostgresRecipeRepo) Add(r *recipe.Recipe) error {
	_, err := rr.Connection.Exec(context.Background(), "INSERT INTO recipes (id, author_id, title, description) VALUES ($1, $2, $3, $4)", r.Id(), r.AuthorId(), r.Title(), r.Description())

	if err != nil {
		return err
	}

	return nil
}

func (rr PostgresRecipeRepo) Retrieve(id string) (*recipe.Recipe, error) {
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

	recipe := recipe.NewRecipe(recipeIdUUID, authorIdUUID, title, description, nil)
	return recipe, nil
}
