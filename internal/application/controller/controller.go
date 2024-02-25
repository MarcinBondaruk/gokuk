package controller

import "github.com/MarcinBondaruk/gokuk/internal/application/service"

type Controller struct {
	recipeService *service.RecipeService
	userService   *service.UserService
}

func NewController(recipeService *service.RecipeService, userService *service.UserService) *Controller {
	return &Controller{
		recipeService: recipeService,
		userService:   userService,
	}
}
