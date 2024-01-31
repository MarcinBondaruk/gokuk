package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/application/api/request"
	"github.com/gin-gonic/gin"
)

func (c Controller) CreateMeal(ctx *gin.Context) {
	var reqBody request.AddRecipeRequest

	if err := ctx.ShouldBindJSON(reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (c Controller) GetMeal(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"recipeId": ctx.Param("recipeId"), "title": "Schabowe babci", "desc": "lorem ipsum"})
}
