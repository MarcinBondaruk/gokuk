package controller

import (
	"fmt"
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/api/request"
	"github.com/MarcinBondaruk/gokuk/internal/application/api/response"
	"github.com/gin-gonic/gin"
)

func (c Controller) CreateUser(ctx *gin.Context) {
	var reqBody request.UserRequest

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, nil)
	}

	id, err := c.userService.CreateUser(reqBody.Username, reqBody.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	resp := response.UserResponse{
		Id:       id,
		Username: reqBody.Username,
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (c Controller) GetUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	user, err := c.userService.GetUser(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	resp := response.UserResponse{
		Id:       user.Id().String(),
		Username: user.Username(),
	}

	ctx.JSON(http.StatusOK, resp)
}
