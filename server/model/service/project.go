package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
	"github.com/phpgoc/zxqpro/routes/response"
)

func HasOwnPermission(c *gin.Context, projectID uint) bool {
	userID := GetUserIDFromAuthMiddleware(c)
	project := entity.Project{}
	result := my_runtime.Db.First(&project, projectID)

	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, "项目不存在"))
		return false
	}
	if userID == 1 {
		return true
	}
	if project.OwnerID != userID {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, "没有权限"))
		return false
	}
	return true
}

func GetRoleType(userID, projectID uint) (entity.RoleType, error) {
	if IsAdmin(userID) {
		return entity.RoleTypeAdmin, nil
	}
	role := entity.Role{}
	res := my_runtime.Db.Where("user_id = ? and project_id = ?", userID, projectID).First(&role)
	if res.Error != nil {
		return entity.RoleTypeNone, res.Error
	}
	return role.RoleType, nil
}

func GetProjectList(userID uint, status, roleType byte, page, pageSize int) (response.ProjectList, error) {
	var responseProjectList response.ProjectList
	var err error
	if userID == 1 {
		projects, total, err := dao.GetProjectsForAdmin(status, page, pageSize)
		if err != nil {
			return responseProjectList, err
		}
		responseProjectList.Total = total
		for _, project := range projects {
			responseProjectList.List = append(responseProjectList.List, response.Project{
				ID:        project.ID,
				Name:      project.Name,
				RoleType:  entity.RoleTypeAdmin,
				OwnerID:   project.OwnerID,
				OwnerName: project.Owner.UserName,
				Status:    project.Status,
			})
		}
	} else {
		roles, total, err := dao.GetProjectsForUser(userID, status, roleType, page, pageSize)
		if err != nil {
			return responseProjectList, err
		}
		responseProjectList.Total = total
		for _, role := range roles {
			responseProjectList.List = append(responseProjectList.List, response.Project{
				ID:        role.Project.ID,
				Name:      role.Project.Name,
				RoleType:  role.RoleType,
				Status:    role.Project.Status,
				OwnerID:   role.Project.OwnerID,
				OwnerName: role.Project.Owner.UserName,
			})
		}
	}
	return responseProjectList, err
}
