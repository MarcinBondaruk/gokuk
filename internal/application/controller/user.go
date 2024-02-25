package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/api/request"
	"github.com/gin-gonic/gin"
)

func (c Controller) CreateUser(ctx *gin.Context) {
	var reqBody request.UserRequest

	if err := ctx.ShouldBindJSON(reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
}
