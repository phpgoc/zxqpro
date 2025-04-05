package routes

import (
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
	api.Use(middleware.AuthLogin())
	docs.SwaggerInfo.BasePath = "/api"
	admin := api.Group("/admin")
	admin.Use(middleware.AuthAdmin())
	admin.POST("/register", AdminRegister)
	admin.GET("/user_list", AdminUserList)
	admin.POST("update_password", AdminUpdatePassword)

	api.GET("/hello_world", HelloWorld)
	api.POST("/user/login", UserLogin)
	api.POST("/user/logout", UserLogout)
	api.GET("/user/info", UserInfo)
	api.POST("/user/update", UserUpdate)
	api.POST("/user/update_password", UserUpdatePassword)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
