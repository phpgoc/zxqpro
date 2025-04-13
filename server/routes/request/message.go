package request

type MessageShareLink struct {
	ToUserId uint   `json:"to_user_id" binding:"required"`
	Link     string `json:"link" binding:"required"`
}
