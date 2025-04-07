package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/docs"
	"github.com/phpgoc/zxqpro/routes/middleware"
)

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApiRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	if gin.Mode() != gin.ReleaseMode {
		// 非 Release 模式下，启用 CORS 中间件，允许所有来源访问
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"*"}
		config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
		config.AllowHeaders = []string{"Content-Type", "Authorization"}
		config.AllowCredentials = true
		api.Use(cors.New(config))
	}
	api.Use(middleware.AuthLogin())
	api.Use(middleware.RateLimit(1000))
	docs.SwaggerInfo.BasePath = "/api"
	admin := api.Group("/admin")
	admin.Use(middleware.AuthAdmin())
	admin.POST("/register", AdminRegister)
	admin.POST("update_password", AdminUpdatePassword)
	admin.POST("create_project", AdminCreateProject)
	admin.POST("/reset_rate_limit", AdminResetRateLimit)

	api.POST("/user/login", middleware.RateLimit(20), UserLogin)

	api.POST("/user/logout", UserLogout)
	api.GET("/user/info", UserInfo)
	api.POST("/user/update", UserUpdate)
	api.POST("/user/update_password", UserUpdatePassword)
	api.GET("/user/list", UserList)

	api.POST("/project/create_role", ProjectCreateRole)
	api.POST("/project/delete_role", ProjectDeleteRole)
	api.POST("/project/update_role", ProjectUpdateRole)
	api.GET("/project/list", ProjectList)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
