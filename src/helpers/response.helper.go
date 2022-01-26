package helpers

// Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// ResponseSuccess is struct
type ResponseSuccess struct {
	Message string
	Data    interface{}
}

// ResponseError is struct
type ResponseError struct {
	Message string
	Errors  string
}

// BuildResponse method is to inject data value to
// dynamic success response
func BuildResponse(param ResponseSuccess) Response {
	res := Response{
		Status:  true,
		Errors:  nil,
		Message: param.Message,
		Data:    param.Data,
	}
	return res
}

// BuildErrorResponse method is to inject data value to
// dynamic failed response
func BuildErrorResponse(param ResponseError) Response {
	res := Response{
		Status:  false,
		Data:    nil,
		Message: param.Message,
		Errors:  param.Errors,
	}
	return res
}
