package controllers

import (
	"net/http"
	"strconv"

	"restapi/src/dtos"
	"restapi/src/helpers"
	"restapi/src/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	services *services.RootService
}

// NewProductController is function
func NewProductController(services *services.RootService) *ProductController {
	return &ProductController{
		services: services,
	}
}

// Create is function
func (ctrl *ProductController) Create(c *gin.Context) {

	var productCreateInput dtos.ProductCreateInput

	// map params request sang model va kiem tra
	if err := c.ShouldBind(&productCreateInput); err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "params invalid",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusOK, response)
		return
	}

	// create product
	product, err := ctrl.services.ProductServivce.Create(c, productCreateInput)

	// check error create product
	if err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "create failure",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// data success
	response := helpers.BuildResponse(helpers.ResponseSuccess{
		Message: "success",
		Data:    product,
	})
	c.JSON(http.StatusOK, response)
}

// GetById is function
func (ctrl *ProductController) GetById(c *gin.Context) {

	idParam := c.Param("productId")
	productId, err := strconv.Atoi(idParam)

	// validate param
	if err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "params invalid",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusOK, response)
		return
	}

	// get data form db
	data, err := ctrl.services.ProductServivce.GetOne(c, productId)

	// check error get data
	if err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "product not found",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusOK, response)
		return
	}

	// data success
	response := helpers.BuildResponse(helpers.ResponseSuccess{
		Message: "success",
		Data:    data,
	})
	c.JSON(http.StatusOK, response)
}

// GetAll is function
func (ctrl *ProductController) GetAll(c *gin.Context) {

	params := dtos.ProductListDto{}

	if err := c.ShouldBind(&params); err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "param invalid",
			Errors:  err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get data from db
	data, err := ctrl.services.ProductServivce.GetAll(c, params)

	// check error get data
	if err != nil {
		response := helpers.BuildErrorResponse(helpers.ResponseError{
			Message: "can not get products",
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
func (ctrl *ProductController) Update(c *gin.Context) {

	price, err := strconv.ParseFloat(c.DefaultPostForm("price", "12.02"), 64)

	// check error
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			helpers.BuildErrorResponse(
				helpers.ResponseError{
					Message: "param invalid",
					Errors:  err.Error(),
				},
			))
		return
	}

	// create product
	product, err := ctrl.services.ProductServivce.Create(
		c.Request.Context(),
		dtos.ProductCreateInput{
			Name:  c.PostForm("name"),
			Price: price,
		},
	)

	// check error create product
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			helpers.BuildErrorResponse(
				helpers.ResponseError{
					Message: "param invalid",
					Errors:  err.Error(),
				},
			))
		return
	}

	// data success
	c.JSON(
		http.StatusOK,
		helpers.BuildResponse(
			helpers.ResponseSuccess{
				Message: "success",
				Data:    product,
			},
		),
	)
}

// Delete is function
func (ctrl *ProductController) Delete(c *gin.Context) {

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
