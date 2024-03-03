package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/api/request"
	"github.com/MarcinBondaruk/gokuk/internal/application/api/response"
	"github.com/MarcinBondaruk/gokuk/internal/application/service"
	"github.com/gin-gonic/gin"
)

type RecipeController struct {
	recipeService *service.RecipeService
}

func NewRecipeController(recipeService *service.RecipeService) *RecipeController {
	return &RecipeController{
		recipeService: recipeService,
	}
}

func (rc RecipeController) CreateRecipe(ctx *gin.Context) {
	var reqBody request.RecipeRequest

	if err := ctx.ShouldBindJSON(reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{})
		return
	}

	ctx.JSON(http.StatusCreated, response.RecipeResponse{})
}

func (rc RecipeController) GetRecipes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []response.RecipeResponse{})
}

func (rc RecipeController) GetRecipeById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.RecipeResponse{})
}

func (rc RecipeController) CreateShoppingListFromRecipes(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, response.ShoppingListResponse{})
}
