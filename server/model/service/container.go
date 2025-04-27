package service

import "github.com/phpgoc/zxqpro/model/dao"

type Container struct {
	AdminService            *AdminService
	ProjectService          *ProjectService
	UserService             *UserService
	TaskService             *TaskService
	TaskTimeEstimateService *TaskTimeEstimateService
}

var ContainerInstance *Container

func InitContainer() {
	ContainerInstance = &Container{
		AdminService:   NewAdminService(dao.ContainerInstance.UserDAO),
		ProjectService: NewProjectService(dao.ContainerInstance.ProjectDAO, dao.ContainerInstance.RoleDAO),
		UserService:    NewUserService(dao.ContainerInstance.UserDAO),
	}
	ContainerInstance.TaskService = NewTaskService(dao.ContainerInstance.TaskDAO, ContainerInstance.ProjectService)
	ContainerInstance.TaskTimeEstimateService = NewTaskTimeEstimateService(dao.ContainerInstance.TaskTimeEstimateDAO, dao.ContainerInstance.TaskDAO, ContainerInstance.ProjectService)
}
