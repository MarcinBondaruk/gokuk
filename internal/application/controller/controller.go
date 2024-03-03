package controller

type Controller struct {
	MenuController    *MenuController
	ProductController *ProductController
	RecipeController  *RecipeController
	UserController    *UserController
}

func NewController(
	menuController *MenuController,
	productController *ProductController,
	recipeController *RecipeController,
	userController *UserController,
) *Controller {
	return &Controller{
		MenuController:    menuController,
		ProductController: productController,
		RecipeController:  recipeController,
		UserController:    userController,
	}
}
