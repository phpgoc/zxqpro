package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/docs"
)

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
)

func ApiRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	docs.SwaggerInfo.BasePath = "/api"
	api.GET("/hello_world", HelloWorld)
	api.POST("/user/register", UserRegister)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
