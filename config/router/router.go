package router

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/config/dependencies"
	"github.com/MarcinBondaruk/gokuk/internal/app_interface/http/recipe"
	"github.com/MarcinBondaruk/gokuk/internal/app_interface/http/user"
)

// creates new router instance and registeres routes
func NewRouter(di *dependencies.DependencyContainer) *http.ServeMux {
	router := http.NewServeMux()

	RegisterRecipeRoutes(router, recipe.NewRecipeHandler(di.RecipeSvc))
	RegisterUserRoutes(router, user.NewUserHandler(di.UserSvc))

	return router
}
