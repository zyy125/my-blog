package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/config"
	"github.com/zyy125/my-blog/backend/internal/pkg/response"
)

// AdminAuth 管理员认证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 验证Token
		token := c.GetHeader("X-Admin-Token")
		if token == "" {
			response.Error(c, "缺少认证Token")
			c.Abort()
			return
		}
		
		if token != config.App.Admin.Token {
			response.Error(c, "Token无效")
			c.Abort()
			return
		}
		
		// 2. 验证IP白名单（如果配置了）
		if len(config.App.Admin.AllowedIPs) > 0 {
			clientIP := c.ClientIP()
			allowed := false
			
			for _, ip := range config.App.Admin.AllowedIPs {
				if ip == clientIP {
					allowed = true
					break
				}
			}
			
			if ! allowed {
				response.Error(c, "IP未授权")
				c.Abort()
				return
			}
		}
		
		// 3. 验证通过，继续处理请求
		c.Next()
	}
}