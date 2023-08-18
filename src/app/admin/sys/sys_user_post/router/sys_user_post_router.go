// Package router SysUserPostRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user_post
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_user_post/api"
	"manager-gin/src/middleware"
)

type SysUserPostRouter struct{}

var sysUserPostApi = api.SysUserPostApiApp

// InitSysUserPostRouter 初始化 SysUserPost 路由信息
func (r *SysUserPostRouter) InitSysUserPostRouter(Router *gin.RouterGroup) {
	sysUserPostRouter := Router.Group("sysUserPost")
	sysUserPostRouter.Use(middleware.JWTAuthFilter())
	sysUserPostRouterWithoutRecord := Router.Group("sysUserPost")
	{
		sysUserPostRouter.POST("create", sysUserPostApi.Create)             // 新建SysUserPost
		sysUserPostRouter.DELETE("delete", sysUserPostApi.Delete)           // 删除SysUserPost
		sysUserPostRouter.DELETE("deleteByIds", sysUserPostApi.DeleteByIds) // 批量删除SysUserPost
		sysUserPostRouter.POST("update", sysUserPostApi.Update)             // 更新SysUserPost
	}
	{
		sysUserPostRouterWithoutRecord.GET("get", sysUserPostApi.Get)   // 根据ID获取SysUserPost
		sysUserPostRouterWithoutRecord.GET("find", sysUserPostApi.Find) // 获取SysUserPost列表
	}
}
