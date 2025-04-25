package main

import (
	"fmt"
	"net/http"

	"github.com/phpgoc/zxqpro/cron"
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/my_runtime"

	"github.com/phpgoc/zxqpro/interfaces"

	"github.com/phpgoc/zxqpro/routes/middleware"

	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/phpgoc/zxqpro/routes/request"

	"github.com/phpgoc/zxqpro/routes"

	"github.com/gobuffalo/packr/v2"
)

// @title Mini Redmine API
// @version 0.1.0

func main() {
	my_runtime.InitCobra()
	if !my_runtime.GinDebugModel {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DefaultWriter = my_runtime.GinLogWriter // 需要在utils.InitCobra() 之后
	router := routes.ApiRoutes()

	interfaces.InitCache()
	dao.InitDb()
	go cron.MainTask() // 需要在 dao.InitDb() 之后

	box := packr.New("static", "../../../static")
	// router.StaticFS("/static", http.FileSystem(box))

	mux := http.NewServeMux()
	if gin.Mode() == gin.ReleaseMode {
		dist := packr.New("dist", "../../../dist")
		mux.Handle("/", spaHandler{box: dist})
	}

	mux.Handle("/static/", middleware.CORSMiddleware(http.StripPrefix("/static/", http.FileServer(box))))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("complexPassword", request.ComplexPasswordValidator)
		if err != nil {
			panic(err)
		}
	}
	mux.Handle("/api/", router)
	mux.Handle("/swagger/", router)
	//_ = router.Run()
	err := http.ListenAndServe(fmt.Sprintf(":%d", my_runtime.Port), mux)
	if err != nil {
		panic(err)
	}
}

// spaHandler 结构体定义
type spaHandler struct {
	box *packr.Box
}

// spaHandler 实现 http.Handler 接口的 ServeHTTP 方法
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 检查请求的文件是否存在于 Packr 盒子中
	if _, err := h.box.Find(r.URL.Path); err == nil {
		// 如果文件存在，使用 http.FileServer 来提供该文件
		http.FileServer(h.box).ServeHTTP(w, r)
		return
	}

	// 如果文件不存在，尝试返回 index.html
	index, err := h.box.Find("index.html")
	if err != nil {
		// 如果 index.html 也不存在，返回 404 错误
		http.NotFound(w, r)
		return
	}

	// 设置响应头并返回 index.html 的内容
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(index)
}
