package response

type Message struct {
	Id       uint    `json:"id"`
	UserName string  `json:"user_name"`
	Message  string  `json:"message"`
	Link     *string `json:"link"`
	Time     string  `json:"time"`
	Read     bool    `json:"read"`
}

type MessageList struct {
	Total int64     `json:"total"`
	List  []Message `json:"list"`
}

type MessageRead struct {
	Id uint `json:"id" binding:"required"`
}
