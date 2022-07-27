package middlewares

import (
	"go_reddit/controller"
	"go_reddit/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式，放在请求头，放在请求体，放在uri
		// 假设token放在header的authorization中，并使用bearer开头
		// Authorization: Bearer xxx.xx.xx
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分隔
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
		}
		// parts[1]是token string ，解析token string
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		//当前请求的userid信息保存到请求的上下文c上
		c.Set(controller.UserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以c.get(UserIDKey)来获取当前请求的用户信息
	}

}
