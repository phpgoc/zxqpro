package service

import (
	"errors"

	"github.com/phpgoc/zxqpro/routes/request"

	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
	"github.com/phpgoc/zxqpro/routes/response"
	"github.com/phpgoc/zxqpro/utils"
)

type ProjectService struct {
	projectDAO *dao.ProjectDAO
	roleDAO    *dao.RoleDAO
}

func NewProjectService(projectDAO *dao.ProjectDAO, roleDAO *dao.RoleDAO) *ProjectService {
	return &ProjectService{
		projectDAO: projectDAO,
		roleDAO:    roleDAO,
	}
}

func (s *ProjectService) HasOwnPermission(userID, projectID uint) error {
	project, err := s.projectDAO.GetByID(projectID)
	if err != nil {
		return err
	}
	if IsAdmin(userID) {
		return nil
	}
	if project.OwnerID != userID {
		return errors.New("没有权限")
	}
	return nil
}

func (s *ProjectService) GetRoleType(userID, projectID uint) entity.RoleType {
	if IsAdmin(userID) {
		return entity.RoleTypeAdmin
	}
	role, res := s.roleDAO.GetRole(userID, projectID)
	if res != nil {
		return entity.RoleTypeNone
	}
	return role.RoleType
}

func (s *ProjectService) GetProjectList(userID uint, status, roleType byte, page, pageSize int) (response.ProjectList, error) {
	var responseProjectList response.ProjectList
	var err error
	if userID == 1 {
		projects, total, err := s.projectDAO.GetProjectsForAdmin(status, page, pageSize)
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
		roles, total, err := s.projectDAO.GetProjectsForUser(userID, status, roleType, page, pageSize)
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

//type TaskOneForList struct {
//	ID                 uint              `json:"id"`
//	Name               string            `json:"name"`
//	Create         CommonIDAndName   `json:"create_user"`
//	ExpectCompleteTime *time.Time        `json:"expect_complete_time"`
//	Status             entity.TaskStatus `json:"status"`
//	TestUser           *CommonIDAndName  `json:"test_user"`
//	SubAssignUser      *CommonIDAndName  `json:"sub_assign_user"`           // 如果是顶级任务，不使用这个字典，非顶级任务，指定的用户是一对一的。顶级任务是一对多的
//	TopTaskAssignUsers []CommonIDAndName `json:"top_task_assign_user_list"` // 顶级任务使用这个，即使只有一个人，也要使用这个
//	CompletedAt        *time.Time        `json:"completed_at"`
//}

func (s *ProjectService) Update(req request.ProjectUpdate) error {
	projectID := req.ID
	project := entity.Project{
		Name:        req.Name,
		Description: req.Description,
		GitAddress:  req.GitAddress,
		Config:      req.Config,
	}
	originalProject, err := s.projectDAO.GetByID(projectID)
	if err != nil {
		return err
	}

	if originalProject.GitAddress == "" && project.GitAddress != "" {
		if !utils.IsGitRepository(project.GitAddress) {
			return errors.New("not a git repository")
		}
		// 如果新的GitAddress不为空，而原来的为空，做一些操作
		project.Status = entity.ProjectStatusActive
		my_runtime.GitPathList.Add(project.GitAddress)
	} else if originalProject.GitAddress != project.GitAddress {
		if project.GitAddress == "" {
			my_runtime.GitPathList.Remove(originalProject.GitAddress)
		} else if !utils.IsGitRepository(project.GitAddress) {
			return errors.New("not a git repository")
		} else {
			my_runtime.GitPathList.Remove(originalProject.GitAddress)
			my_runtime.GitPathList.Add(project.GitAddress)
		}
	}

	return s.projectDAO.Update(&project)
}

func (s *ProjectService) ProjectInfo(projectID uint) (*response.ProjectInfo, error) {
	project, err := s.projectDAO.GetByID(projectID)
	if err != nil {
		return nil, err
	}
	projectInfo := response.ProjectInfo{
		ID:          project.ID,
		Name:        project.Name,
		OwnerID:     project.OwnerID,
		OwnerName:   project.Owner.UserName,
		Description: project.Description,
		GitAddress:  project.GitAddress,
		Config:      project.Config,
		Status:      project.Status,
	}
	return &projectInfo, nil
}

func (s *ProjectService) UpdateProjectStatus(projectID uint, status entity.ProjectStatus) error {
	return s.projectDAO.UpdateProjectStatus(projectID, status)
}

func (s *ProjectService) UserList(projectID uint) (*response.UserList, error) {
	roles, err := s.roleDAO.GetAllUserByProjectID(projectID)
	if err != nil {
		return nil, err
	}
	res := &response.UserList{}
	for _, role := range roles {
		res.List = append(res.List, response.User{
			ID:       role.UserID,
			Name:     role.User.Name,
			UserName: role.User.UserName,
			Avatar:   role.User.Avatar,
			Email:    role.User.Email,
			RoleType: role.RoleType,
		})
	}
	return res, nil
}
