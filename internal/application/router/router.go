package router

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/controller"
	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
)

func NewRouter(ctrl controller.Controller) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	apiV1 := router.Group("/api/v1")
	recipes := apiV1.Group("/recipes")
	recipes.POST("", ctrl.CreateRecipe)
	recipes.GET("", ctrl.GetRecipes)
	recipes.GET("/shopping-list", ctrl.CreateShoppingListFromRecipes)
	recipes.GET("/:recipeId", ctrl.GetRecipeById)

	mealPlan := apiV1.Group("/meal-plans")
	mealPlan.POST("", ctrl.CreateMealPlan)
	mealPlan.GET("/:mealPlanId/shopping-list", ctrl.GetShoppingListForMealPlan)

	products := apiV1.Group("/products")
	products.GET("", ctrl.GetProducts)

	pprof.Register(router)
	return router
}
