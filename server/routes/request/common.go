package request

type Page struct {
	Page     int `json:"page" form:"page" binding:"required,min=1" default:"1"`
	PageSize int `json:"page_size" form:"page_size" binding:"required,min=5" default:"10"`
}

type CommonID struct {
	ID uint `json:"id" form:"id" binding:"required,min=1"`
}
