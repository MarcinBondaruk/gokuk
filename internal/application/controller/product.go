package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/api/response"
	"github.com/gin-gonic/gin"
)

func (c Controller) BulkAdd(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, response.ProductsBulkAddResponse{})
}

func (c Controller) GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ProductResponse{})
}
