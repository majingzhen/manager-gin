package middleware

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/common/response"
	"manager-gin/src/framework"
	"net/http"
	"strings"
)

func JWTAuthFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "账号未登录",
			})
			// 不进行下面的请求处理了！
			c.Abort()
			return
			// c.Redirect(http.StatusFound, "/login")
		}

		// 检查Authorization头是否以Bearer开头
		tokenString := strings.TrimPrefix(strings.Replace(authHeader, "\n", "", -1), "Bearer ")
		if tokenString == authHeader {
			response.Unauthorized(c)
			return
		}

		mc, err := framework.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "账号未登录",
			})
			// 不进行下面的请求处理了！
			c.Abort()
			return
			// c.Redirect(http.StatusFound, "/login")
		}
		c.Set("user_id", mc.UserId)
		c.Next()
	}
}
