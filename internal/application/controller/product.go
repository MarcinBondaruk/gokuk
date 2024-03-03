package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/api/response"
	"github.com/MarcinBondaruk/gokuk/internal/application/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *service.ProductService
}

func (pc ProductController) BulkAdd(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, response.ProductsBulkAddResponse{})
}

func (pc ProductController) GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ProductResponse{})
}
