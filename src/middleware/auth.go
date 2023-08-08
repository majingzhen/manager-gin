package middleware

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/common"
	"net/http"
)

func JWTAuthFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "无效的Token",
			})
			// 不进行下面的请求处理了！
			c.Abort()
			return
			// c.Redirect(http.StatusFound, "/login")
		}

		mc, err := common.ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "无效的Token",
			})
			// 不进行下面的请求处理了！
			c.Abort()
			return
			// c.Redirect(http.StatusFound, "/login")
		}
		c.Set("loginUser", mc.UserId)
		c.Next()
	}
}
