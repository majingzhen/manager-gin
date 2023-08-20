// Package router SysUserRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_user/api"
	"manager-gin/src/middleware"
)

type SysUserRouter struct{}

var sysUserApi = api.SysUserApiApp

// InitSysUserRouter 初始化 SysUser 路由信息
func (r *SysUserRouter) InitSysUserRouter(Router *gin.RouterGroup) {
	sysUserRouter := Router.Group("sysUser")
	sysUserRouter.Use(middleware.JWTAuthFilter())
	sysUserRouterWithoutRecord := Router.Group("sysUser")
	{
		sysUserRouter.POST("create", sysUserApi.Create)             // 新建SysUser
		sysUserRouter.DELETE("delete", sysUserApi.Delete)           // 删除SysUser
		sysUserRouter.DELETE("deleteByIds", sysUserApi.DeleteByIds) // 批量删除SysUser
		sysUserRouter.POST("update", sysUserApi.Update)             // 更新SysUser
	}
	{
		sysUserRouterWithoutRecord.GET("get", sysUserApi.Get)   // 根据ID获取SysUser
		sysUserRouterWithoutRecord.GET("list", sysUserApi.List) // 获取SysUser列表
	}
}
