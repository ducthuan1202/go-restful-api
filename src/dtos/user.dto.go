package dtos

// UserCreateInput is struct
type UserCreateInput struct {
	Name     string `form:"name" json:"name" binding:"required,alphaunicode,min=6,max=20"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

// UserLoginDto is struct
type UserLoginDto struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
