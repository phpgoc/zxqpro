package main

import (
	"net/http"

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

	_ = router.Run()
}
