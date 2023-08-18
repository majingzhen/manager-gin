// Package router SysUserRoleRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user_role
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_user_role/api"
	"manager-gin/src/middleware"
)

type SysUserRoleRouter struct{}

var sysUserRoleApi = api.SysUserRoleApiApp

// InitSysUserRoleRouter 初始化 SysUserRole 路由信息
func (r *SysUserRoleRouter) InitSysUserRoleRouter(Router *gin.RouterGroup) {
	sysUserRoleRouter := Router.Group("sysUserRole")
	sysUserRoleRouter.Use(middleware.JWTAuthFilter())
	sysUserRoleRouterWithoutRecord := Router.Group("sysUserRole")
	{
		sysUserRoleRouter.POST("create", sysUserRoleApi.Create)             // 新建SysUserRole
		sysUserRoleRouter.DELETE("delete", sysUserRoleApi.Delete)           // 删除SysUserRole
		sysUserRoleRouter.DELETE("deleteByIds", sysUserRoleApi.DeleteByIds) // 批量删除SysUserRole
		sysUserRoleRouter.POST("update", sysUserRoleApi.Update)             // 更新SysUserRole
	}
	{
		sysUserRoleRouterWithoutRecord.GET("get", sysUserRoleApi.Get)   // 根据ID获取SysUserRole
		sysUserRoleRouterWithoutRecord.GET("find", sysUserRoleApi.Find) // 获取SysUserRole列表
	}
}
