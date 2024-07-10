package dependencies

import (
	"github.com/MarcinBondaruk/gokuk/config/database"
	"github.com/MarcinBondaruk/gokuk/internal/services/recipe"
	"github.com/MarcinBondaruk/gokuk/internal/services/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DependencyContainer struct {
	database *pgxpool.Pool

	UserSvc   *user.UserService
	RecipeSvc *recipe.RecipeService
}

func Initialize() *DependencyContainer {
	database := database.InitConnPool()
	recipeSvc := recipe.NewRecipeService(database)
	userSvc := user.NewUserService(database)

	return &DependencyContainer{
		database:  database,
		RecipeSvc: recipeSvc,
		UserSvc:   userSvc,
	}
}

func (di *DependencyContainer) TearDown() {
	di.database.Close()
}
