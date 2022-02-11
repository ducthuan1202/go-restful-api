package controllers

import (
	"fmt"
	"net/http"

	"restapi/src/dtos"
	"restapi/src/helpers"
	"restapi/src/models"
	"restapi/src/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	services *services.RootService
}

// NewUserController is function
func NewUserController(services *services.RootService) *UserController {
	return &UserController{
		services: services,
	}
}

// Create is function
func (ctrl UserController) Create(c *gin.Context) {

	var createInput dtos.UserCreateInput

	// map params request sang model va kiem tra
	if err := c.ShouldBind(&createInput); err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "params invalid",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusOK, response)
		return
	}

	// tao user moi va kiem tra loi (neu co)
	user, err := ctrl.services.UserServivce.Create(c, createInput)
	if err != nil {
		res := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "create failure",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusOK, res)
		return
	}

	// data success
	response := helpers.BuildResponse(helpers.ResponseSuccess{
		Message: "success",
		Data:    user,
	})
	c.JSON(http.StatusOK, response)
}

// Login is function
func (ctrl UserController) Login(c *gin.Context) {

	var loginDto dtos.UserLoginDto

	// map params request sang model va kiem tra
	if err := c.ShouldBind(&loginDto); err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "params invalid",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusOK, response)
		return
	}

	authResult := ctrl.services.UserServivce.Repository.VerifyCredential(loginDto.Email, loginDto.Password)

	if v, ok := authResult.(models.User); ok {
		token := ctrl.services.JWTServivce.GenerateToken(fmt.Sprintf("%v", v.ID))
		v.Token = token

		ctrl.services.UserServivce.Repository.SaveAuthToken(v.CreateAuthToken(token))

		response := helpers.BuildResponse(helpers.ResponseSuccess{
			Message: "OK",
			Data:    v,
		})
		c.JSON(http.StatusOK, response)
		return
	}

	response := helpers.BuildErrorResponse(helpers.ResponseError{
		Message: "Invalid Credential",
		Errors:  "Please check again your credential",
	})
	c.JSON(http.StatusUnauthorized, response)
}
