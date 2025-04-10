package response

import "github.com/phpgoc/zxqpro/model/entity"

type User struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	UserName        string  `json:"user_name"`
	Email           *string `json:"email"`
	Avatar          byte    `json:"avatar"`
	entity.RoleType `json:"role_type"`
}

type UserList struct {
	List []User `json:"list"`
}
