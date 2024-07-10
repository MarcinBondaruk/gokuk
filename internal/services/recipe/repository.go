package recipe

import (
	"context"
	_ "embed"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	connPool *pgxpool.Pool
}

func newRecipeRepository(connection *pgxpool.Pool) *repository {
	return &repository{
		connPool: connection,
	}
}

func (rr *repository) Add(r *recipe) error {
	tx, err := rr.connPool.Begin(context.Background())
	if err != nil {
		return err
	}

	recipeArgs := pgx.NamedArgs{
		"id":          r.ID,
		"author_id":   r.authorID,
		"title":       r.title,
		"description": r.description,
	}

	_, err = tx.Exec(context.Background(), "INSERT INTO recipes (id, author_id, title, description) VALUES (@id, @author_id, @title, @description)", recipeArgs)

	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	for _, ingredient := range r.ingredients {
		inredientArgs := pgx.NamedArgs{
			"recipe_id": r.ID,
			"name":      ingredient.Name,
			"quantity":  ingredient.Quantity,
			"unit":      ingredient.Unit,
		}
		_, err = tx.Exec(context.Background(), "INSERT INTO recipe_ingredient (recipe_id, name, quantity, unit) VALUES (@recipe_id, @name, @quantity, @unit)", inredientArgs)
		if err != nil {
			tx.Rollback(context.Background())
			return err
		}
	}

	err = tx.Commit(context.Background())

	if err != nil {
		return err
	}

	return nil
}

// TODO: Add pagination, maybe optimize n+1 problem
func (rr *repository) GetRecipes() ([]RecipeView, error) {
	recipeViews := []RecipeView{}
	recipes, err := rr.connPool.Query(context.Background(), "SELECT id, author_id, title, description FROM recipes")
	if err != nil {
		return nil, err
	}

	for recipes.Next() {
		recipeView := RecipeView{}

		err = recipes.Scan(&recipeView.ID, &recipeView.AuthorID, &recipeView.Title, &recipeView.Description)
		if err != nil {
			log.Println(err)
			continue
		}

		ingredients, err := rr.connPool.Query(context.Background(), "SELECT name, quantity, unit FROM recipe_ingredient WHERE recipe_id = $1", recipeView.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		for ingredients.Next() {
			ingredientView := IngredientView{}

			err := ingredients.Scan(&ingredientView.Name, &ingredientView.Quantity, &ingredientView.Unit)
			if err != nil {
				log.Println(err)
				continue
			}

			recipeView.Ingredients = append(recipeView.Ingredients, ingredientView)
		}

		recipeViews = append(recipeViews, recipeView)
	}

	return recipeViews, nil
}
