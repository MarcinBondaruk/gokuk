package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c Controller) CreateMeal(ctx *gin.Context) {
	log.Println("new meal created")
	ctx.JSON(http.StatusOK, nil)
}

func (c Controller) GetMeal(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"recipeId": ctx.Param("recipeId"), "title": "Schabowe babci", "desc": "lorem ipsum"})
}
