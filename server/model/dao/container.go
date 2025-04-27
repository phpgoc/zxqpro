package dao

import "github.com/phpgoc/zxqpro/my_runtime"

type Container struct {
	ProjectDAO          *ProjectDAO
	RoleDAO             *RoleDAO
	UserDAO             *UserDAO
	TaskDAO             *TaskDAO
	TaskTimeEstimateDAO *TaskTimeEstimateDAO
}

var ContainerInstance *Container

func InitContainer() {
	ContainerInstance = &Container{
		ProjectDAO:          NewProjectDAO(my_runtime.Db),
		RoleDAO:             NewRoleDAO(my_runtime.Db),
		UserDAO:             NewUserDAO(my_runtime.Db),
		TaskDAO:             NewTaskDAO(my_runtime.Db),
		TaskTimeEstimateDAO: NewTaskTimeEstimateDAO(my_runtime.Db),
	}
}
