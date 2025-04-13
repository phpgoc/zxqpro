package response

type Message struct {
	Id      uint    `json:"id"`
	Message string  `json:"message"`
	Link    *string `json:"link"`
	Read    bool    `json:"read"`
}

type MessageList struct {
	Total int64     `json:"total"`
	List  []Message `json:"messages"`
}

type MessageRead struct {
	Id uint `json:"id" binding:"required"`
}
