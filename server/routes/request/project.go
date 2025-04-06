package request

import "github.com/phpgoc/zxqpro/model/entity"

type ProjectUpsertRole struct {
	UserId    uint            `json:"user_id"`
	ProjectId uint            `json:"project_id"`
	RoleType  entity.RoleType `json:"role_type"`
}

type ProjectDeleteRole struct {
	UserId    uint `json:"user_id"`
	ProjectId uint `json:"project_id"`
}
