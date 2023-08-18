// Package router SysPostRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/system/api"
	"manager-gin/src/utils"
	"net/http"
)

type SystemRouter struct{}

var systemApi = api.SystemApiApp

// InitSystemRouter 初始化 系统 路由信息
func (r *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) {
	systemRouter := Router.Group("")
	{
		systemRouter.POST("login", systemApi.Login) // 登录
		systemRouter.GET("captcha", systemApi.CaptchaImage)
		systemRouter.GET("/", func(c *gin.Context) {
			password := utils.EncryptionPassword("123456", "c989acb7109d46c0bae3ce1b3c962f36")
			c.JSON(http.StatusOK, password)
		})
	}
}
