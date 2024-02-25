package controller

import "github.com/MarcinBondaruk/gokuk/internal/application/service"

type Controller struct {
	RecipeService *service.RecipeService
	UserService   *service.UserService
}
