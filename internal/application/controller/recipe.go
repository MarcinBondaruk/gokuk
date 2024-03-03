package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/api/request"
	"github.com/MarcinBondaruk/gokuk/internal/application/api/response"
	"github.com/gin-gonic/gin"
)

func (c Controller) CreateRecipe(ctx *gin.Context) {
	var reqBody request.RecipeRequest

	if err := ctx.ShouldBindJSON(reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{})
		return
	}

	ctx.JSON(http.StatusCreated, response.RecipeResponse{})
}

func (c Controller) GetRecipes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []response.RecipeResponse{})
}

func (c Controller) GetRecipeById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.RecipeResponse{})
}

func (c Controller) CreateShoppingListFromRecipes(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, response.ShoppingListResponse{})
}
