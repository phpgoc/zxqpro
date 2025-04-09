package response

import "github.com/phpgoc/zxqpro/model/entity"

type Project struct {
	ID        uint                 `json:"id"`
	Name      string               `json:"name"`
	RoleType  entity.RoleType      `json:"role_type"`
	Status    entity.ProjectStatus `json:"status"`
	OwnerID   uint                 `json:"owner_id"`
	OwnerName string               `json:"owner_name"`
}
type ProjectList struct {
	Total int64     `json:"total"`
	List  []Project `json:"list"`
}
