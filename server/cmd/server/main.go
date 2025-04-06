package main

import (
	"net/http"

	"github.com/gin-contrib/cors"

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

	if gin.Mode() != gin.ReleaseMode {
		// 非 Release 模式下，启用 CORS 中间件，允许所有来源访问
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"*"}
		config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
		config.AllowHeaders = []string{"Content-Type", "Authorization"}
		config.AllowCredentials = true
		router.Use(cors.New(config))
	}
	_ = router.Run()
}
