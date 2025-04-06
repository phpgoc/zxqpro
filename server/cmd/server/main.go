package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phpgoc/zxqpro/model/dao"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/phpgoc/zxqpro/routes/request"

	"github.com/phpgoc/zxqpro/routes"

	"github.com/gobuffalo/packr/v2"
	"github.com/phpgoc/zxqpro/utils"
)

func main() {
	router := routes.ApiRoutes()
	err := utils.InitLog()
	if err != nil {
		return
	}

	dao.InitDb()
	box := packr.New("static", "../../../static")
	router.StaticFS("/static", http.FileSystem(box))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("complexPassword", request.ComplexPasswordValidator)
		if err != nil {
			panic(err)
		}
	}
	// 如果不是release版本，开启跨域
	if gin.Mode() != gin.ReleaseMode {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
		router.Use(func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			if c.Request.Method == http.MethodOptions {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		})
	}
	_ = router.Run()
}
