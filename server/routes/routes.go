package routes

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/docs"
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
	// 非 Release 模式下，启用 CORS 中间件，允许所有来源访问
	config := cors.Config{
		AllowOriginFunc: func(origin string) bool {
			// 白名单：只允许指定域名
			allowedOrigins := map[string]bool{
				"http://localhost:5173": true,
				"http://localhost:5174": true,
				"tauri://app":           true,
			}
			utils.LogWarn(fmt.Sprintf("origin: %s", origin))
			return allowedOrigins[origin]
		},
		// AllowOrigins:     []string{"http://locahost:5173", "http://locahost:5174", "tauri://app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}

	utils.LogWarn("CORS is enabled in debug mode, please disable it in production.")
	router.Use(cors.New(config))
	//}
	api := router.Group("/api")

	api.Use(middleware.AuthLogin())
	api.Use(middleware.RateLimit(500))
	docs.SwaggerInfo.BasePath = "/api"
	admin := api.Group("/admin")
	admin.Use(middleware.AuthAdmin())
	admin.POST("/register", AdminRegister)
	admin.POST("update_password", AdminUpdatePassword)
	admin.POST("create_project", AdminCreateProject)
	admin.POST("/reset_rate_limit", AdminResetRateLimit)

	api.POST("/user/login", middleware.RateLimit(10), UserLogin)

	api.POST("/user/logout", UserLogout)
	api.GET("/user/info", UserInfo)
	api.POST("/user/update", UserUpdate)
	api.POST("/user/update_password", UserUpdatePassword)
	api.GET("/user/list", UserList)

	api.POST("/project/create_role", ProjectCreateRole)
	api.POST("/project/delete_role", ProjectDeleteRole)
	api.POST("/project/update_role", ProjectUpdateRole)
	api.GET("/project/list", ProjectList)

	api.POST("/message/share_link", MessageShareLink)
	api.GET("/message/receive_list", MessageReceiveList)
	api.GET("/message/send_list", MessageSendList)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
