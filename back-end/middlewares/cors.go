package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		// AllowOrigins: []string{"*"}, // 这可能会带来安全风险
		AllowOrigins:     []string{"http://localhost:5173", "https://my-frontend-domain.com"}, // 指定允许访问的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                 // 指定允许的 HTTP 方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},                 // 指定允许的 HTTP 请求头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 如果跨域请求需要携带 Cookie 或其他凭据，需要设置为 true
		MaxAge:           24 * time.Hour,
	})
}
