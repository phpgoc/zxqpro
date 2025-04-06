package response

import "github.com/phpgoc/zxqpro/model/entity"

type Project struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	RoleType entity.RoleType `json:"role_type"`
}
type ProjectList struct {
	Projects []Project `json:"projects"`
}
