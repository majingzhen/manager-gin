// Package router SysMenuRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_menu/api"
	"manager-gin/src/middleware"
)

type SysMenuRouter struct{}

var sysMenuApi = api.SysMenuApiApp

// InitSysMenuRouter 初始化 SysMenu 路由信息
func (r *SysMenuRouter) InitSysMenuRouter(Router *gin.RouterGroup) {
	sysMenuRouter := Router.Group("sysMenu")
	sysMenuRouter.Use(middleware.JWTAuthFilter())
	sysMenuRouterWithoutRecord := Router.Group("sysMenu")
	{
		sysMenuRouter.POST("create", sysMenuApi.Create)             // 新建SysMenu
		sysMenuRouter.DELETE("delete", sysMenuApi.Delete)           // 删除SysMenu
		sysMenuRouter.DELETE("deleteByIds", sysMenuApi.DeleteByIds) // 批量删除SysMenu
		sysMenuRouter.POST("update", sysMenuApi.Update)             // 更新SysMenu
	}
	{
		sysMenuRouterWithoutRecord.GET("get", sysMenuApi.Get)   // 根据ID获取SysMenu
		sysMenuRouterWithoutRecord.GET("find", sysMenuApi.Find) // 获取SysMenu列表
	}
}
