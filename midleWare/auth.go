package midleWare

import (
	"demo02/initjwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWTAuthMiddleware JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "请求头中缺少Authorization字段",
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		tokenParts := strings.SplitN(authHeader, " ", 2)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "Authorization格式错误，应为：Bearer {token}",
			})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		// 解析并验证token
		claims, err := initjwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "无效的token",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		// 检查token类型，只允许access token访问API
		if claims.TokenType != "access" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "token类型错误，请使用access token",
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("userID", int(claims.UserID))
		c.Set("username", claims.Username)

		// 继续处理请求
		c.Next()
	}
}
