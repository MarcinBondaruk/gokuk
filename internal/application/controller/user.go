package controller

import (
	"fmt"
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/application/api/request"
	"github.com/MarcinBondaruk/gokuk/internal/application/api/response"
	"github.com/MarcinBondaruk/gokuk/internal/application/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc UserController) CreateUser(ctx *gin.Context) {
	var reqBody request.UserRequest

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, nil)
	}

	id, err := uc.userService.CreateUser(reqBody.Username, reqBody.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	resp := response.UserResponse{
		Id:       id,
		Username: reqBody.Username,
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (uc UserController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	user, err := uc.userService.GetUser(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	resp := response.UserResponse{
		Id:       user.Id().String(),
		Username: user.Username(),
	}

	ctx.JSON(http.StatusOK, resp)
}
