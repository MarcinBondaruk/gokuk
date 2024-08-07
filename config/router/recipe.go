package router

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/config/middleware"
	"github.com/MarcinBondaruk/gokuk/internal/app_interface/http/recipe"
)

func RegisterRecipeRoutes(router *http.ServeMux, handler *recipe.RecipeHandler) *http.ServeMux {
	router.HandleFunc("GET /api/v1/recipes", middleware.Pipe(handler.GetRecipes, middleware.Logger()))
	router.HandleFunc("POST /api/v1/recipes", middleware.Pipe(handler.CreateRecipe, middleware.Logger()))

	return router
}
