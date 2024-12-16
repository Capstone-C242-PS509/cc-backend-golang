package handler

import (
	"backend-bangkit/dto"
	"backend-bangkit/pkg/common"
	"backend-bangkit/pkg/errs"
	service "backend-bangkit/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.AuthService
}

func NewUserHandler(userService service.AuthService) *UserHandler {
	return &UserHandler{userService}
}

// ==================================================

func (u *UserHandler) Register(ctx *gin.Context) {
	var requestBody dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	registeredUser, err := u.userService.Register(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusOK, registeredUser))
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var requestBody dto.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	token, err := u.userService.Login(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusOK, token))
}
