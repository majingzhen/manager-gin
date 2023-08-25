// Package router SysMenuRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/api"
	"manager-gin/src/middleware"
)

type SysMenuRouter struct {
	sysMenuApi api.SysMenuApi
}

// InitSysMenuRouter 初始化 SysMenu 路由信息
func (r *SysMenuRouter) InitSysMenuRouter(Router *gin.RouterGroup) {
	sysMenuRouter := Router.Group("sysMenu").Use(middleware.JWTAuthFilter())
	sysMenuRouterWithoutRecord := Router.Group("sysMenu").Use(middleware.JWTAuthFilter())
	{
		sysMenuRouter.POST("create", r.sysMenuApi.Create)       // 新建SysMenu
		sysMenuRouter.DELETE("delete/:id", r.sysMenuApi.Delete) // 删除SysMenu
		sysMenuRouter.POST("update", r.sysMenuApi.Update)       // 更新SysMenu
	}
	{
		sysMenuRouterWithoutRecord.GET("get/:id", r.sysMenuApi.Get)                                       // 根据ID获取SysMenu
		sysMenuRouterWithoutRecord.GET("treeSelect", r.sysMenuApi.SelectMenuTree)                         // 获取SysMenu树形列表
		sysMenuRouterWithoutRecord.GET("roleMenuTreeSelect/:roleId", r.sysMenuApi.SelectMenuTreeByRoleId) // 根据ID获取SysMenu
		sysMenuRouterWithoutRecord.GET("list", r.sysMenuApi.List)                                         // 获取SysMenu列表
	}
}
