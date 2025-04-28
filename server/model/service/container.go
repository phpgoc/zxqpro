package service

import "github.com/phpgoc/zxqpro/model/dao"

type Container struct {
	AdminService            *AdminService
	ProjectService          *ProjectService
	UserService             *UserService
	TaskService             *TaskService
	TaskTimeEstimateService *TaskTimeEstimateService
	MessageService          *MessageService
}

var ContainerInstance *Container

func InitContainer() {
	ContainerInstance = &Container{
		AdminService:   NewAdminService(dao.ContainerInstance.UserDAO),
		ProjectService: NewProjectService(dao.ContainerInstance.ProjectDAO, dao.ContainerInstance.RoleDAO),
		UserService:    NewUserService(dao.ContainerInstance.UserDAO),
		MessageService: NewMessageService(dao.ContainerInstance.MessageDAO, dao.ContainerInstance.UserDAO),
	}
	// 下边是依赖其他service的service初始化，需要先初始化依赖的service
	ContainerInstance.TaskService = NewTaskService(dao.ContainerInstance.TaskDAO, dao.ContainerInstance.UserDAO, ContainerInstance.ProjectService)
	ContainerInstance.TaskTimeEstimateService = NewTaskTimeEstimateService(dao.ContainerInstance.TaskTimeEstimateDAO, dao.ContainerInstance.TaskDAO, ContainerInstance.ProjectService)
}
