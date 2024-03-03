package controller

type Controller struct {
	menuController    *MenuController
	productController *ProductController
	recipeController  *RecipeController
	userController    *UserController
}

func NewController(
	menuController *MenuController,
	productController *ProductController,
	recipeController *RecipeController,
	userController *UserController,
) *Controller {
	return &Controller{
		menuController:    menuController,
		productController: productController,
		recipeController:  recipeController,
		userController:    userController,
	}
}
