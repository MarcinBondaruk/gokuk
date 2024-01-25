package router

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/infrastructure/controller"

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
	recipes.POST("", ctrl.CreateMeal)
	recipes.GET("/:recipeId", ctrl.GetMeal)

	return router
}
