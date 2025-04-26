package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/docs"
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/service"
	"github.com/phpgoc/zxqpro/my_runtime"
	"github.com/phpgoc/zxqpro/routes/middleware"
	"github.com/phpgoc/zxqpro/utils"
)

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApiRoutes() *gin.Engine {
	router := gin.Default()
	// if gin.Mode() != gin.ReleaseMode {

	router.Use(cors.New(middleware.CorsConfig))

	sseManager := utils.NewSSEManager()
	router.Use(func(c *gin.Context) {
		c.Set("sseManager", sseManager)
		c.Next()
	})

	api := router.Group("/api")

	api.Use(middleware.AuthLogin())
	api.Use(middleware.RateLimit(500))
	docs.SwaggerInfo.BasePath = "/api"
	admin := api.Group("/admin")
	admin.Use(middleware.AuthAdmin())
	adminHandler := NewAdminHandler(service.NewAdminService(dao.NewUserDAO(my_runtime.Db)))
	admin.POST("/register", adminHandler.AdminRegister)
	admin.POST("update_password", adminHandler.AdminUpdatePassword)
	admin.POST("create_project", adminHandler.AdminCreateProject)
	admin.POST("/reset_rate_limit", AdminResetRateLimit)

	userHandler := NewUserHandler(service.NewUserService(dao.NewUserDAO(my_runtime.Db)))
	api.POST("/user/login", middleware.RateLimit(10), userHandler.UserLogin)

	// 这个接口不是不需要验证登录，而是单独验证，在验证错误时不能返回json，这是不符合sse规范的
	router.GET("/api/sse", ServerSideEvent)

	api.GET("/sse/test", TestSendSelf)

	api.POST("/user/logout", userHandler.Logout)
	api.GET("/user/info", userHandler.Info)
	api.POST("/user/update", userHandler.Update)
	api.POST("/user/update_password", userHandler.UpdatePassword)
	api.GET("/user/list", userHandler.List)

	projectHandler := NewProjectHandler(service.NewProjectService(dao.NewProjectDAO(my_runtime.Db), dao.NewRoleDAO(my_runtime.Db)))

	api.POST("/project/create_role", projectHandler.ProjectCreateRole)
	api.POST("/project/delete_role", projectHandler.ProjectDeleteRole)
	api.POST("/project/update_role", projectHandler.ProjectUpdateRole)
	api.GET("/project/list", projectHandler.ProjectList)
	api.POST("/project/update", projectHandler.ProjectUpdate)
	api.POST("/project/update_status", projectHandler.ProjectUpdateStatus)
	api.GET("/project/info", projectHandler.ProjectInfo)
	api.GET("/project/role_in", projectHandler.ProjectRoleIn)
	api.POST("/project/task_list", projectHandler.ProjectTaskList)
	api.GET("/project/user_list", projectHandler.UserList)

	api.POST("/message/share_link", MessageShareLink)
	api.GET("/message/receive_list", MessageReceiveList)
	api.GET("/message/send_list", MessageSendList)
	api.POST("/message/read", MessageRead)
	api.POST("/message/manual", MessageManual)

	taskHandler := NewTaskHandler(service.NewTaskService(dao.NewTaskDAO(my_runtime.Db), service.NewProjectService(dao.NewProjectDAO(my_runtime.Db), dao.NewRoleDAO(my_runtime.Db))))
	api.POST("/task/create_top", taskHandler.TaskCreateTop)
	api.POST("/task/update_top", taskHandler.TaskUpdateTop)
	api.POST("/task/assign_top", taskHandler.TaskAssignTop)
	api.GET("/task/public_info", taskHandler.TaskPublicInfo)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
