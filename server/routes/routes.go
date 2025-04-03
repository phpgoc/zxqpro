package routes

import (
	"github.com/gin-gonic/gin"
	"githutb.com/phpgoc/zxqpro/server/docs"
)
import ginSwagger "github.com/swaggo/gin-swagger"
import swaggerfiles "github.com/swaggo/files"

func ApiRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	docs.SwaggerInfo.BasePath = "/api"
	api.GET("/hello_world", HelloWorld)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
