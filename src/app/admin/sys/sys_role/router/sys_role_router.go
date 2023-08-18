// Package router SysRoleRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-18 14:00:54
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_role/api"
	"manager-gin/src/middleware"
)

type SysRoleRouter struct{}

var sysRoleApi = api.SysRoleApiApp

// InitSysRoleRouter 初始化 SysRole 路由信息
func (r *SysRoleRouter) InitSysRoleRouter(Router *gin.RouterGroup) {
	sysRoleRouter := Router.Group("sysRole")
	sysRoleRouter.Use(middleware.JWTAuthFilter())
	sysRoleRouterWithoutRecord := Router.Group("sysRole")
	{
		sysRoleRouter.POST("create", sysRoleApi.Create)             // 新建SysRole
		sysRoleRouter.DELETE("delete", sysRoleApi.Delete)           // 删除SysRole
		sysRoleRouter.DELETE("deleteByIds", sysRoleApi.DeleteByIds) // 批量删除SysRole
		sysRoleRouter.POST("update", sysRoleApi.Update)             // 更新SysRole
	}
	{
		sysRoleRouterWithoutRecord.GET("get", sysRoleApi.Get)   // 根据ID获取SysRole
		sysRoleRouterWithoutRecord.GET("find", sysRoleApi.Find) // 获取SysRole列表
	}
}
