package main

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/configs"
	"github.com/MarcinBondaruk/gokuk/internal/application/controller"
	"github.com/MarcinBondaruk/gokuk/internal/application/router"
	"github.com/MarcinBondaruk/gokuk/internal/application/service"
)

func main() {
	postgres := configs.GetPostgresConnection()
	defer configs.ClosePostgresConnection(postgres)

	recipeService := &service.RecipeService{}
	userService := &service.UserService{}

	controller := controller.NewController(recipeService, userService)
	router := router.NewRouter(controller)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
