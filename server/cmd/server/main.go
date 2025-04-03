package main

import (
	"net/http"

	"github.com/phpgoc/zxqpro/middle_ware"
	"github.com/phpgoc/zxqpro/routes"

	"github.com/gobuffalo/packr/v2"
	"github.com/patrickmn/go-cache"
	"github.com/phpgoc/zxqpro/interfaces"
	"github.com/phpgoc/zxqpro/utils"
)

func main() {
	router := routes.ApiRoutes()
	router.Use(middle_ware.ValidationMiddleware())
	err := utils.InitLog()
	if err != nil {
		return
	}

	utils.InitDb()
	box := packr.New("static", "../static")
	router.StaticFS("/static", http.FileSystem(box))
	interfaces.Cache.Set("key", "value123", cache.DefaultExpiration)
	value, _ := interfaces.Cache.Get("key")
	utils.LogInfo(value.(string))

	_ = router.Run()
}
