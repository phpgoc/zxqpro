package response

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Avatar   byte   `json:"avatar"`
}

type UserList struct {
	Total int64  `json:"total"`
	Users []User `json:"users"`
}
