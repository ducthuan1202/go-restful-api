package dtos

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// TODO: CAN CUSTOM LAI MESSAGE ERROR VALIDATE TRA VE

// link validate guide:
// https://pkg.go.dev/github.com/go-playground/validator#hdr-Alpha_Unicode

// CategoryListDto is struct
type CategoryListDto struct {
	Limit uint `form:"limit"`
}

// CategoryListDto is struct
type CategoryCreateInput struct {
	Name string `form:"name" json:"name" binding:"required,alphaunicode,min=6,max=20"`

	// Name string `form:"name" json:"name" binding:"required,categoryName"`

	// Name string `form:"name" json:"name" binding:"required,oneof=male female"`
}

// function validate
var categoryNameValidate validator.Func = func(fl validator.FieldLevel) bool {

	match, err := regexp.MatchString("^([a-zA-Z0-9]{4,10})$", fl.Field().String())
	if err != nil {
		return false
	}

	return match
}

// su dung ham init de dang ky func validate cho binding
func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("categoryName", categoryNameValidate)
	}
}
