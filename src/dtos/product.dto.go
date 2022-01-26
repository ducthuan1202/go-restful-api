package dtos

// ProductCreateInput is struct
type ProductCreateInput struct {
	Name  string  `form:"name" json:"name" binding:"required"`
	Price float64 `form:"price" json:"price" binding:"required"`
}

// ProductListDto is struct
type ProductListDto struct {
	Limit uint `form:"limit" json:"limit" binding:"gte=2,lte=5"`
}
