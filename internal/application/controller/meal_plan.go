package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/api/response"
	"github.com/gin-gonic/gin"
)

func (c Controller) CreateMealPlan(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, response.MealPlanResponse{})
}

func (c Controller) GetShoppingListForMealPlan(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ShoppingListResponse{})
}
