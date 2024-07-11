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

//go:embed sql/InsertRecipe.sql
var insertRecipeSql string

//go:embed sql/InsertRecipeIngredient.sql
var insertRecipeIngredientSql string

func (rr *repository) Add(ctx context.Context, r *recipe) error {
	tx, err := rr.connPool.Begin(ctx) //TODO: move to middleware
	if err != nil {
		return err
	}

	recipeArgs := pgx.NamedArgs{
		"author_id":   r.authorID,
		"title":       r.title,
		"description": r.description,
	}

	err = tx.QueryRow(ctx, insertRecipeSql, recipeArgs).Scan(r.ID)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	for _, ingredient := range r.ingredients {
		inredientArgs := pgx.NamedArgs{
			"recipe_id": r.ID,
			"name":      ingredient.Name,
			"quantity":  ingredient.Value,
			"unit":      ingredient.Unit,
		}
		_, err := tx.Exec(ctx, insertRecipeIngredientSql, inredientArgs)

		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	err = tx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (rr *repository) GetRecipes(ctx context.Context) ([]RecipeView, error) {
	recipeViews := []RecipeView{}
	recipes, err := rr.connPool.Query(ctx, "SELECT id, author_id, title, description FROM recipe")
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

		ingredients, err := rr.connPool.Query(ctx, "SELECT name, quantity, unit FROM recipe_ingredient WHERE recipe_id = $1", recipeView.ID)
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
