package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/docs"
	"github.com/phpgoc/zxqpro/model/service"
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
	adminHandler := NewAdminHandler(service.ContainerInstance.AdminService)
	admin.POST("/register", adminHandler.Register)
	admin.POST("update_password", adminHandler.UpdatePassword)
	admin.POST("create_project", adminHandler.CreateProject)
	admin.POST("/reset_rate_limit", AdminResetRateLimit)

	userHandler := NewUserHandler(service.ContainerInstance.UserService)
	api.POST("/user/login", middleware.RateLimit(10), userHandler.UserLogin)

	// 这个接口不是不需要验证登录，而是单独验证，在验证错误时不能返回json，这是不符合sse规范的
	router.GET("/api/sse", ServerSideEvent)

	api.GET("/sse/test", TestSendSelf)

	api.POST("/user/logout", userHandler.Logout)
	api.GET("/user/info", userHandler.Info)
	api.POST("/user/update", userHandler.Update)
	api.POST("/user/update_password", userHandler.UpdatePassword)
	api.GET("/user/list", userHandler.List)

	projectHandler := NewProjectHandler(service.ContainerInstance.ProjectService)

	api.POST("/project/create_role", projectHandler.CreateRole)
	api.POST("/project/delete_role", projectHandler.DeleteRole)
	api.POST("/project/update_role", projectHandler.UpdateRole)
	api.GET("/project/list", projectHandler.List)
	api.POST("/project/update", projectHandler.Update)
	api.POST("/project/update_status", projectHandler.UpdateStatus)
	api.GET("/project/info", projectHandler.Info)
	api.GET("/project/role_in", projectHandler.RoleIn)

	api.GET("/project/user_list", projectHandler.UserList)

	messageHandler := NewMessageHandler(service.ContainerInstance.MessageService)
	api.POST("/message/share_link", messageHandler.ShareLink)
	api.GET("/message/receive_list", messageHandler.ReceiveList)
	api.GET("/message/send_list", messageHandler.SendList)
	api.POST("/message/read", messageHandler.Read)
	api.POST("/message/manual", messageHandler.Manual)

	taskHandler := NewTaskHandler(service.ContainerInstance.TaskService)
	api.POST("/task/create_top", taskHandler.CreateTop)
	api.POST("/task/update_top", taskHandler.UpdateTop)
	api.POST("/task/assign_top", taskHandler.AssignTop)
	api.GET("/task/public_info", taskHandler.PublicInfo)

	// taskHandle，路由却在project里，这样做是为了这样的人物是共有的，没有所有人的权限，不在项目中的人也能通过分享的方式看到
	api.POST("/project/task_list", taskHandler.ProjectTaskList)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
