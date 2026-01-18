package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/internal/handler"
	"github.com/zyy125/my-blog/backend/internal/pkg/response"
)

// AdminAuth 管理员认证中间件（验证临时token）
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Token
		token := c.GetHeader("X-Admin-Token")
		if token == "" {
			response.Error(c, "缺少认证Token")
			c.Abort()
			return
		}

		// 2. 验证 Token 是否有效
		if !handler.CheckToken(token) {
			response.Error(c, "Token无效或已过期")
			c.Abort()
			return
		}

		// 3. 验证通过，继续处理请求
		c.Next()
	}
}
