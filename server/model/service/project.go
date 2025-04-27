package service

import (
	"errors"

	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/utils"
	"gorm.io/gorm/clause"

	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
	"github.com/phpgoc/zxqpro/routes/response"
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
	project, err := s.projectDAO.GetProjectByID(projectID)
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

func (s *ProjectService) GetTaskList(req request.ProjectTaskList) (res response.TaskList, err error) {
	orderValidList := []string{"id", "expect_complete_time", "test_user_id", "completed_at"}
	if err = request.IsAllValidOrder(req.OrderList, orderValidList); err != nil {
		return res, err
	}

	var taskList []entity.Task
	model := my_runtime.Db.Model(&entity.Task{}).Preload(clause.Associations).Where("project_id = ?", req.ID)
	if req.CreateUserID != 0 {
		model = model.Where("create_user_id = ?", req.CreateUserID)
	}
	if req.Status != 0 {
		model = model.Where("status = ?", req.Status)
	}
	if req.TopStatus != 0 {
		if req.TopStatus == 1 {
			model = model.Where("parent_id = ?", 0)
		} else {
			model = model.Where("parent_id != ?", 0)
		}
	}
	for _, order := range req.OrderList {
		if order.Desc {
			model = model.Order(order.Field + " desc")
		} else {
			model = model.Order(order.Field + " asc")
		}
	}
	err = model.Count(&res.Total).Error
	if err != nil {
		return res, err
	}
	err = model.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&taskList).Error
	var testUser *response.CommonIDAndName = nil
	var subAssignUser *response.CommonIDAndName = nil
	var topTaskAssignUsers []response.CommonIDAndName
	for _, task := range taskList {
		if task.TesterID != 0 {
			testUser = &response.CommonIDAndName{
				ID:   task.TesterID,
				Name: task.Tester.UserName,
			}
		} else {
			testUser = nil
		}
		if task.AssignUserID != 0 {
			subAssignUser = &response.CommonIDAndName{
				ID:   task.AssignUserID,
				Name: task.AssignUser.UserName,
			}
		} else {
			subAssignUser = nil
		}
		if task.TopTaskAssignUsers != nil {
			topTaskAssignUsers = make([]response.CommonIDAndName, 0)
			for _, user := range task.TopTaskAssignUsers {
				topTaskAssignUsers = append(topTaskAssignUsers, response.CommonIDAndName{
					ID:   user.ID,
					Name: user.UserName,
				})
			}
		}
		res.List = append(res.List, response.TaskOneForList{
			ID:   task.ID,
			Name: task.Name,
			CreateUser: response.CommonIDAndName{
				ID:   task.CreateUserID,
				Name: task.CreateUser.UserName,
			},
			ExpectCompleteTime: task.ExpectCompleteTime,
			Status:             task.Status,
			TestUser:           testUser,
			SubAssignUser:      subAssignUser,
			TopTaskAssignUsers: topTaskAssignUsers,
			CompletedAt:        task.CompletedAt,
		})
	}

	return res, err
}

func (s *ProjectService) UpdateProject(projectID uint, project entity.Project) error {
	originalProject := entity.Project{}
	res := my_runtime.Db.Model(&entity.Project{}).Where("id = ?", projectID).First(&originalProject)
	if res.Error != nil {
		return res.Error
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

	res = my_runtime.Db.Model(&entity.Project{}).Where("id = ?", projectID).Updates(project)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		return errors.New("not found")
	}
	return nil
}

func (s *ProjectService) ProjectInfo(projectID uint) (*response.ProjectInfo, error) {
	project, err := s.projectDAO.GetProjectByID(projectID)
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
