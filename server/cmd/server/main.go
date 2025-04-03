package main

import (
	"githutb.com/phpgoc/zxqpro/server/routes"
	"net/http"

	"github.com/gobuffalo/packr/v2"
	"github.com/patrickmn/go-cache"
	"githutb.com/phpgoc/zxqpro/server/interfaces"
	"githutb.com/phpgoc/zxqpro/server/utils"
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
	interfaces.Cache.Set("key", "value123", cache.DefaultExpiration)
	value, _ := interfaces.Cache.Get("key")
	utils.LogInfo(value.(string))

	_ = router.Run()
}
