package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/api/response"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
}

func NewMealPlanController() *MenuController {
	return &MenuController{}
}

func (mc MenuController) CreateMealPlan(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, response.MealPlanResponse{})
}

func (mc MenuController) GetShoppingListForMealPlan(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ShoppingListResponse{})
}
