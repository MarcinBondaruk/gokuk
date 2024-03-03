package main

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/configs"
	"github.com/MarcinBondaruk/gokuk/internal/application/controller"
	"github.com/MarcinBondaruk/gokuk/internal/application/router"
	"github.com/MarcinBondaruk/gokuk/internal/application/service"
	"github.com/MarcinBondaruk/gokuk/internal/infrastructure/repository"
)

func main() {
	postgres := configs.GetPostgresConnection()
	defer configs.ClosePostgresConnection(postgres)
	userRepository := &repository.UserRepositoryImpl{
		Connection: postgres,
	}

	userService := &service.UserService{UserRepository: userRepository}

	controller := controller.NewController(
		controller.NewMenuController(),
		controller.NewProductController(&service.ProductService{}),
		controller.NewRecipeController(&service.RecipeService{}),
		controller.NewUserController(userService),
	)

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
