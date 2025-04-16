package request

type MessageShareLink struct {
	ToUserId uint   `json:"to_user_id" binding:"required"`
	Link     string `json:"link" binding:"required"`
}

type MessageList struct {
	Read     bool `form:"read"`
	Page     int  `form:"page"  bindings:"min=1" default:"1"`
	PageSize int  `form:"page_size" bindings:"min=1,max=100" default:"10"`
}

type MessageRead struct {
	Id uint `json:"id" binding:"required"`
}

type ManualMessage struct {
	UserIds []uint  `json:"user_ids" binding:"required"`
	Content string  `json:"content" binding:"required"`
	Link    *string `json:"link"`
}
