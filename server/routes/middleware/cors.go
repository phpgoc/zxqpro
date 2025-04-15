package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/phpgoc/zxqpro/utils"
)

var allowedOrigins = map[string]bool{
	"http://localhost:5173": true,
	"http://localhost:5174": true,
	"tauri://app":           true,
}

var CorsConfig = cors.Config{
	AllowOriginFunc: func(origin string) bool {
		// 白名单：只允许指定域名

		utils.LogWarn(fmt.Sprintf("origin: %s is %t", origin, allowedOrigins[origin]))

		return allowedOrigins[origin]
	},
	AllowAllOrigins:  false,
	AllowCredentials: true,
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 允许所有域名进行跨域调用
		utils.LogError("hello")
		if allowedOrigins[r.Header.Get("Origin")] {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		}
		// 允许的请求方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		// 允许客户端携带认证信息（如 cookie）
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// 如果是 OPTIONS 请求，直接返回 204 状态码
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}
