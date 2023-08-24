// Package router SysRoleRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-21 17:37:56
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
	sysRoleRouter := Router.Group("sysRole").Use(middleware.JWTAuthFilter())
	sysRoleRouterWithoutRecord := Router.Group("sysRole").Use(middleware.JWTAuthFilter())
	{
		sysRoleRouter.POST("create", sysRoleApi.Create)            // 新建SysRole
		sysRoleRouter.DELETE("delete/:ids", sysRoleApi.Delete)     // 删除SysRole
		sysRoleRouter.POST("update", sysRoleApi.Update)            // 更新SysRole
		sysRoleRouter.PUT("changeStatus", sysRoleApi.ChangeStatus) // 更新SysRole状态
	}
	{
		sysRoleRouterWithoutRecord.GET("get/:id", sysRoleApi.Get) // 根据ID获取SysRole
		sysRoleRouterWithoutRecord.GET("page", sysRoleApi.Page)   // 分页获取SysRole列表
		sysRoleRouterWithoutRecord.GET("list", sysRoleApi.List)   // 分页获取SysRole列表
	}
}
