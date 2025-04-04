package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/docs"
	"github.com/phpgoc/zxqpro/middleware"
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
	api.GET("/hello_world", HelloWorld)
	api.POST("/user/register", UserRegister)
	api.POST("/user/login", UserLogin)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
