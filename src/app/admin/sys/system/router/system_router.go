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
	"manager-gin/src/middleware"
)

type SystemRouter struct{}

var systemApi = api.SystemApiApp

// InitSystemRouter 初始化 系统 路由信息
func (r *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) {
	Router.POST("login", systemApi.Login) // 登录
	Router.GET("captchaImage", systemApi.CaptchaImage)
	systemRouter := Router.Group("").Use(middleware.JWTAuthFilter())
	{
		systemRouter.GET("getInfo", systemApi.GetUserInfo)
		systemRouter.GET("getRouters", systemApi.GetRouters)
	}
}
