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

type ProjectList struct {
	Page     int  `form:"page"  bindings:"min=1" default:"1"`
	PageSize int  `form:"page_size" bindings:"min=1,max=100" default:"10"`
	Status   byte `form:"status" bindings:"min=0,max=4" default:"0"`
	RoleType byte `form:"role_type" bindings:"min=0,max=5" default:"0"`
}
