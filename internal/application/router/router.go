package router

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(ctrl *controller.Controller) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	apiV1 := router.Group("/api/v1")

	recipes := apiV1.Group("/recipes")
	recipes.POST("", ctrl.RecipeController.CreateRecipe)
	recipes.GET("", ctrl.RecipeController.GetRecipes)
	recipes.GET("/shopping-list", ctrl.RecipeController.CreateShoppingListFromRecipes)
	recipes.GET("/:recipeId", ctrl.RecipeController.GetRecipeById)

	mealPlan := apiV1.Group("/meal-plans")
	mealPlan.POST("", ctrl.MenuController.CreateMealPlan)
	mealPlan.GET("/:mealPlanId/shopping-list", ctrl.MenuController.GetShoppingListForMealPlan)

	products := apiV1.Group("/products")
	products.GET("", ctrl.ProductController.GetProducts)
	products.POST("/bulk-add", ctrl.ProductController.BulkAddProducts)

	users := apiV1.Group("/users")
	users.POST("", ctrl.UserController.CreateUser)
	users.GET("/:userId", ctrl.UserController.GetUser)

	return router
}
