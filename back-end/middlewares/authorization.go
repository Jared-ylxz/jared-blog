package middlewares

import (
	"exchangeapp/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 Authorization Header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is missing",
			})
			ctx.Abort() // 只执行这一个中间件，不再执行后续的中间件（如有多个中间件）
			return
		}

		// 提取 Token
		inputToken := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseToken(inputToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Invalid token: %s", err.Error()),
			})
			ctx.Abort()
			return
		}
		ctx.Set("userID", claims.UserID) // 将用户ID和用户名存入上下文
		ctx.Set("username", claims.Username)
		ctx.Next() // 执行后续的中间件
	}
}
