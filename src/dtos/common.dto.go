package dtos

// dto: Data Transfer Object
// dung de mo ta chi tiet payload cho cac ham them, sua, xoa du lieu

// CommonListDto is struct
type CommonListDto struct {
	Limit uint `form:"limit"`
}
