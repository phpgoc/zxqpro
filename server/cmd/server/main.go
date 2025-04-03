package main

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/phpgoc/zxqpro/request"

	"github.com/phpgoc/zxqpro/routes"

	"github.com/gobuffalo/packr/v2"
	"github.com/patrickmn/go-cache"
	"github.com/phpgoc/zxqpro/interfaces"
	"github.com/phpgoc/zxqpro/utils"
)

func main() {
	router := routes.ApiRoutes()
	err := utils.InitLog()
	if err != nil {
		return
	}

	utils.InitDb()
	box := packr.New("static", "../static")
	router.StaticFS("/static", http.FileSystem(box))
	router.With()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("complexPassword", request.ComplexPasswordValidator)
		if err != nil {
			panic(err)
		}
	}
	interfaces.Cache.Set("key", "value123", cache.DefaultExpiration)
	value, _ := interfaces.Cache.Get("key")
	utils.LogInfo(value.(string))

	_ = router.Run()
}
