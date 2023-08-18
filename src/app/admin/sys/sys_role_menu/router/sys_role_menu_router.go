// Package router SysRoleMenuRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_role_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_role_menu/api"
	"manager-gin/src/middleware"
)

type SysRoleMenuRouter struct{}

var sysRoleMenuApi = api.SysRoleMenuApiApp

// InitSysRoleMenuRouter 初始化 SysRoleMenu 路由信息
func (r *SysRoleMenuRouter) InitSysRoleMenuRouter(Router *gin.RouterGroup) {
	sysRoleMenuRouter := Router.Group("sysRoleMenu")
	sysRoleMenuRouter.Use(middleware.JWTAuthFilter())
	sysRoleMenuRouterWithoutRecord := Router.Group("sysRoleMenu")
	{
		sysRoleMenuRouter.POST("create", sysRoleMenuApi.Create)             // 新建SysRoleMenu
		sysRoleMenuRouter.DELETE("delete", sysRoleMenuApi.Delete)           // 删除SysRoleMenu
		sysRoleMenuRouter.DELETE("deleteByIds", sysRoleMenuApi.DeleteByIds) // 批量删除SysRoleMenu
		sysRoleMenuRouter.POST("update", sysRoleMenuApi.Update)             // 更新SysRoleMenu
	}
	{
		sysRoleMenuRouterWithoutRecord.GET("get", sysRoleMenuApi.Get)   // 根据ID获取SysRoleMenu
		sysRoleMenuRouterWithoutRecord.GET("find", sysRoleMenuApi.Find) // 获取SysRoleMenu列表
	}
}
