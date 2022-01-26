package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"restapi/src/dtos"
	"restapi/src/helpers"
	"restapi/src/services"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	services *services.RootService
}

// NewCategoryController is function
func NewCategoryController(services *services.RootService) *CategoryController {
	return &CategoryController{
		services: services,
	}
}

// Create is function
func (ctrl CategoryController) Create(c *gin.Context) {

	var categoryCreateInput dtos.CategoryCreateInput

	// map params request sang model va kiem tra
	if err := c.ShouldBind(&categoryCreateInput); err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "params invalid",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusOK, response)
		return
	}

	// tao category moi va kiem tra loi (neu co)
	category, err := ctrl.services.CategoryServivce.Create(c, categoryCreateInput)
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
		Data:    category,
	})
	c.JSON(http.StatusOK, response)
}

// GetById is function
func (ctrl CategoryController) GetById(c *gin.Context) {
	// data success
	c.JSON(
		http.StatusOK,
		helpers.BuildResponse(
			helpers.ResponseSuccess{
				Message: "success",
				Data:    nil,
			},
		),
	)
}

// GetAll is function
func (ctrl CategoryController) GetAll(c *gin.Context) {

	params := dtos.CategoryListDto{}
	if err := c.ShouldBind(&params); err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "param invalid",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get data from db
	data, err := ctrl.services.CategoryServivce.GetAll(c, params)
	fmt.Println(params)

	// check error get data
	if err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "can not get categories",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// data success
	response := helpers.BuildResponse(helpers.ResponseSuccess{
		Message: "success",
		Data:    data,
	})
	c.JSON(http.StatusOK, response)
}

// Update is function
func (ctrl CategoryController) Update(c *gin.Context) {
	// data success
	c.JSON(
		http.StatusOK,
		helpers.BuildResponse(
			helpers.ResponseSuccess{
				Message: "success",
				Data:    nil,
			},
		),
	)
}

// Delete is function
func (ctrl CategoryController) Delete(c *gin.Context) {

	ids := strings.TrimSpace(c.DefaultPostForm("ids", ""))
	if ids == "" {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "Bad request",
			Errors:  "list id invalid",
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get data form db
	_, err := ctrl.services.CategoryServivce.DeleteMulti(c, ids)

	// check error get data
	if err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "product not found",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// data success
	response := helpers.BuildResponse(helpers.ResponseSuccess{
		Message: "success",
		Data:    nil,
	})
	c.JSON(http.StatusOK, response)
}

// DeleteById is function
func (ctrl CategoryController) DeleteById(c *gin.Context) {

	id := c.Param("id")

	// get data form db
	_, err := ctrl.services.CategoryServivce.DeleteById(c, id)

	// check error get data
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			helpers.BuildErrorResponse(
				helpers.ResponseError{
					Message: "product not found",
					Errors:  err.Error(),
				},
			),
		)
		return
	}

	// data success
	c.JSON(
		http.StatusOK,
		helpers.BuildResponse(
			helpers.ResponseSuccess{
				Message: "success",
				Data:    nil,
			},
		),
	)
}
