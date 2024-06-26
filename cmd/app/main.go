package main

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/config"
	"github.com/MarcinBondaruk/gokuk/internal/services/recipe"
	"github.com/MarcinBondaruk/gokuk/internal/services/user"
	"github.com/MarcinBondaruk/gokuk/internal/user_interface/http_handler"
)

func main() {
	// check envs
	// boonstrap repositories, services
	// register routes
	postgres := config.GetPostgresConnection()
	defer config.ClosePostgresConnection(postgres)

	userService := user.NewUserService(user.NewUserRepository(postgres))
	userHandler := http_handler.NewUserHandler(userService)

	recipeService := recipe.NewRecipeService(recipe.NewRecipeRepository(postgres))
	recipeHandler := http_handler.NewRecipeHandler(recipeService)

	router := http.NewServeMux()
	router.HandleFunc("POST /api/v1/users", userHandler.CreateUser)

	router.HandleFunc("POST /api/v1/recipes", recipeHandler.CreateRecipe)
	router.HandleFunc("GET /api/v1/recipes", recipeHandler.GetRecipes)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
